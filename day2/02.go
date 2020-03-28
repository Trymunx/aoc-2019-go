package day2

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/pkg/errors"
)

func loadInput() []int64 {
	inputFile, err := os.Open("day2/input.txt")
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

type intcodeComputer struct {
	memory []int64
	ptr    int
}

type opcode struct {
	code    int64
	argLen  int
	instant bool
}

type instruction interface {
	Next() instruction
	Compute()
}

func (pc *intcodeComputer) doOp() error {
	if pc.ptr >= len(pc.memory) || pc.ptr < 0 {
		return fmt.Errorf("Pointer out of range: %v", pc.ptr)
	}
	opCode := pc.memory[pc.ptr]
	instr, err := opCodeToInstr(opCode)
	if err != nil {
		fmt.Println(errors.Wrapf(err, "pointer: %v", pc.ptr))
	}

	return nil
}

func opCodeToInstr(code int64) (*instruction, error) {
	switch code {
	case 1:
		return &opcode{1, 3, false}, nil
	case 2:
		return &opcode{2, 3, false}, nil
	case 99:
		return &opcode{99, 0, false}, nil
	default:
		return nil, fmt.Errorf("Error: expected opcode got: %v: %T", code, code)
	}
}

// Part1 calculates the answer to day2 part1
func Part1() int64 {
	input := loadInput()

	for _, op := range input {
		switch op {
		case 1:
		case 2:
		case 99:
			return 0
		default:
			log.Printf("Error: expected opcode got: %v: %T", op, op)
		}
	}
	return int64(len(input))
}
