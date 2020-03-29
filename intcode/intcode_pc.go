package intcode

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"

	"github.com/pkg/errors"
)

// LoadFromFile loads intcode text files and formats them as strings
func LoadFromFile(fileName string) []int64 {
	inputFile, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer inputFile.Close()

	scanner := bufio.NewScanner(inputFile)
	var input []int64

	scanner.Split(func(data []byte, atEOF bool) (advance int, token []byte, err error) {
		if atEOF && len(data) == 0 {
			return 0, nil, nil
		}
		if i := bytes.IndexByte(data, ','); i >= 0 {
			return i + 1, data[0:i], nil
		}
		if atEOF {
			// Strip new lines or terminal \r
			if data[len(data)-1] == '\r' || data[len(data)-1] == '\n' {
				return len(data), data[0 : len(data)-1], nil
			}
			return len(data), data, nil
		}
		return 0, nil, nil
	})

	for scanner.Scan() {
		number, err := strconv.ParseInt(scanner.Text(), 10, 0)
		if err != nil {
			fmt.Print(fmt.Errorf("could not parse %v as int: %w", scanner.Text(), err))
		} else {
			input = append(input, number)
		}
	}

	return input
}

// Computer is an instance of a computer
type Computer struct {
	Memory []int64
	ptr    int64
	Halted bool
}

// NewComputer returns an instance of a new intcode computer
func NewComputer(memory []int64, ptr int64, halted bool) *Computer {
	return &Computer{memory, ptr, halted}
}

// Step does the next computation and moves the pointer forward.
// It returns whether the program should halt
func (pc *Computer) Step() {
	if pc.Halted {
		return
	}
	if pc.ptr >= int64(len(pc.Memory)) || pc.ptr < 0 {
		fmt.Println(fmt.Errorf("Pointer out of range: %v", pc.ptr))
		pc.Halted = true
		return
	}
	opcode, err := ToOpcode(pc.Memory[pc.ptr])
	if err != nil {
		fmt.Println(errors.Wrapf(err, "pointer: %v", pc.ptr))
		pc.Halted = true
		return
	}
	err = opcode.Compute(pc.ptr, pc.Memory)
	if err != nil {
		fmt.Println(errors.Wrap(err, "failed to compute opcode"))
		pc.Halted = true
		return
	}
	if opcode.NumberToNext() == 0 {
		pc.Halted = true
		return
	}
	pc.ptr += opcode.NumberToNext() + 1
	return
}

// Opcode is an instruction for an intcode computer
type Opcode interface {
	NumberToNext() int64
	Compute(ptr int64, memory []int64) error
}

// ToOpcode turns a string into an opcode
func ToOpcode(val int64) (Opcode, error) {
	code := val % 100

	var argModes []int64
	digits := math.Trunc(math.Log10(float64(val))) - 2
	for i := digits; i >= 0; i-- {
		divisor := int64(math.Pow(10, digits-i+2))
		argModes = append(argModes, (val/divisor)%10)
	}

	switch code {
	case 1:
		return NewOp1(3, argModes), nil
	case 2:
		return NewOp2(3, argModes), nil
	case 99:
		return NewOp99(), nil
	default:
		return nil, fmt.Errorf("Error: expected opcode got: %v: %T", val, val)
	}
}

// PrintStatus logs the current running status of a computer
func PrintStatus(comp *Computer) string {
	out := ""
	for _, val := range comp.Memory[:comp.ptr] {
		out += fmt.Sprintf("%v ", strconv.FormatInt(val, 10))
	}
	out += fmt.Sprintf("[%v]", comp.Memory[comp.ptr])
	for _, val := range comp.Memory[comp.ptr+1:] {
		out += fmt.Sprintf(" %v", strconv.FormatInt(val, 10))
	}
	return out
}
