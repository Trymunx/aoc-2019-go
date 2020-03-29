package intcode

import (
	"fmt"
)

// Op1 is the add operation
type Op1 struct {
	ArgLen   int64
	ArgModes []int64
}

// NewOp1 returns an instance of the add operation
func NewOp1(argLen int64, argModes []int64) *Op1 {
	return &Op1{
		ArgLen:   argLen,
		ArgModes: padMissingArgModes(argLen-1, argModes),
	}
}

// NumberToNext gives the number to jump to next instruction
func (op *Op1) NumberToNext() int64 {
	return op.ArgLen + 1
}

// Compute adds two numbers together and stores them
func (op *Op1) Compute(ptr int64, memory []int64, input int64) (int64, error) {
	if int64(len(memory)) < ptr+op.ArgLen {
		return 0, fmt.Errorf("Expected %v chars after ptr position: %v but memory length is %v", op.ArgLen, ptr, len(memory))
	}
	args := calculateArgModes(op.ArgLen-1, op.ArgModes, ptr, memory)
	memory[memory[ptr+3]] = args[0] + args[1]
	return 0, nil
}

// Op2 is the multiply operation
type Op2 struct {
	ArgLen   int64
	ArgModes []int64
}

// NewOp2 returns an instance of the multiply operation
func NewOp2(argLen int64, argModes []int64) *Op2 {
	return &Op2{
		ArgLen:   argLen,
		ArgModes: padMissingArgModes(argLen-1, argModes),
	}
}

// NumberToNext gives the number to jump to next instruction
func (op *Op2) NumberToNext() int64 {
	return op.ArgLen + 1
}

// Compute adds two numbers together and stores them
func (op *Op2) Compute(ptr int64, memory []int64, input int64) (int64, error) {
	if int64(len(memory)) < ptr+op.ArgLen {
		return 0, fmt.Errorf("Expected %v chars after ptr position: %v but memory length is %v", op.ArgLen, ptr, len(memory))
	}
	args := calculateArgModes(op.ArgLen-1, op.ArgModes, ptr, memory)
	memory[memory[ptr+3]] = args[0] * args[1]
	return 0, nil
}

// Op3 saves an input into a given position in memory
type Op3 struct {
	ArgLen int64
}

// NewOp3 returns an instance of the input operation
func NewOp3(argLen int64) *Op3 {
	return &Op3{
		ArgLen: argLen,
	}
}

// NumberToNext gives the number to jump to next instruction
func (op *Op3) NumberToNext() int64 {
	return op.ArgLen + 1
}

// Compute inputs a number in a given position
func (op *Op3) Compute(ptr int64, memory []int64, input int64) (int64, error) {
	if int64(len(memory)) < ptr+op.ArgLen {
		return 0, fmt.Errorf("Expected %v chars after ptr position: %v but memory length is %v", op.ArgLen, ptr, len(memory))
	}
	memory[memory[ptr+1]] = input
	return 0, nil
}

// Op4 outputs the value of its argument
type Op4 struct {
	ArgLen   int64
	ArgModes []int64
}

// NewOp4 returns an instance of the output operation
func NewOp4(argLen int64, argModes []int64) *Op4 {
	return &Op4{
		ArgLen:   argLen,
		ArgModes: padMissingArgModes(argLen, argModes),
	}
}

// NumberToNext gives the number to jump to next instruction
func (op *Op4) NumberToNext() int64 {
	return op.ArgLen + 1
}

// Compute outputs a number from a position or value
func (op *Op4) Compute(ptr int64, memory []int64, input int64) (int64, error) {
	if int64(len(memory)) < ptr+op.ArgLen {
		return 0, fmt.Errorf("Expected %v chars after ptr position: %v but memory length is %v", op.ArgLen, ptr, len(memory))
	}
	args := calculateArgModes(op.ArgLen, op.ArgModes, ptr, memory)
	return args[0], nil
}

// Op99 is the halt operation
type Op99 struct {
	ArgLen int64
}

// NewOp99 returns an instance of the halt instruction
func NewOp99() *Op99 {
	return &Op99{0}
}

// NumberToNext gives the number to jump to next instruction
func (op *Op99) NumberToNext() int64 {
	return op.ArgLen
}

// Compute terminates the program
func (op *Op99) Compute(ptr int64, memory []int64, input int64) (int64, error) {
	return 0, nil
}

func padMissingArgModes(argLen int64, argModes []int64) []int64 {
	missing := argLen - int64(len(argModes))
	if missing > 0 {
		return append(argModes, make([]int64, missing)...)
	}
	return argModes
}

func calculateArgModes(argLen int64, argModes []int64, ptr int64, memory []int64) []int64 {
	args := make([]int64, argLen)
	for i := int64(0); i < argLen; i++ {
		if argModes[i] == 1 {
			args[i] = memory[ptr+i+1]
		} else {
			args[i] = memory[memory[ptr+i+1]]
		}
	}
	return args
}
