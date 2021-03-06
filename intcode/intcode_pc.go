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
func NewComputer(memory []int64) *Computer {
	return &Computer{memory, 0, false}
}

// Step does the next computation and moves the pointer forward.  It returns
// the output of the step computation or -1 on error, as well as whether the
// output is valid output or just the zero value.
func (pc *Computer) Step(input int64) (int64, bool) {
	if pc.Halted {
		return 0, false
	}
	if pc.ptr >= int64(len(pc.Memory)) || pc.ptr < 0 {
		fmt.Println(fmt.Errorf("Pointer out of range: %v", pc.ptr))
		pc.Halted = true
		return -1, false
	}
	opcode, err := ToOpcode(pc.Memory[pc.ptr])
	if err != nil {
		fmt.Println(errors.Wrapf(err, "pointer: %v", pc.ptr))
		pc.Halted = true
		return -1, false
	}
	output, ok, err := opcode.Compute(pc, input)
	if err != nil {
		fmt.Println(errors.Wrap(err, "failed to compute opcode"))
		pc.Halted = true
		return -1, false
	}
	return output, ok
}

// Opcode is an instruction for an intcode computer
type Opcode interface {
	// Compute does the operation indicated by the opcode and advances the
	// pointer.  It returns the output or 0, a boolean to say whether or not that
	// output is valid or nil, and a possible error.
	Compute(pc *Computer, input int64) (int64, bool, error)
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
	case 3:
		return NewOp3(1), nil
	case 4:
		return NewOp4(1, argModes), nil
	case 5:
		return NewOp5(2, argModes), nil
	case 6:
		return NewOp6(2, argModes), nil
	case 7:
		return NewOp7(3, argModes), nil
	case 8:
		return NewOp8(3, argModes), nil
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
