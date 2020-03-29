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
	missingArgModes := argLen - int64(len(argModes))
	if missingArgModes > 0 {
		argModes = append(make([]int64, missingArgModes), argModes...)
	}
	return &Op1{
		ArgLen:   argLen,
		ArgModes: argModes,
	}
}

// NumberToNext gives the number to jump to next instruction
func (op *Op1) NumberToNext() int64 {
	return op.ArgLen
}

// Compute adds two numbers together and stores them
func (op *Op1) Compute(ptr int64, memory []int64) error {
	if int64(len(memory)) < ptr+op.ArgLen {
		return fmt.Errorf("Expected %v chars after ptr position: %v but memory length is %v", op.ArgLen, ptr, len(memory))
	}
	args := make([]int64, op.ArgLen)
	for i := int64(0); i < op.ArgLen; i++ {
		if op.ArgModes[i] == 1 {
			args[i] = memory[ptr+i+1]
		} else {
			args[i] = memory[memory[ptr+i+1]]
		}
	}
	memory[memory[ptr+3]] = args[0] + args[1]
	return nil
}

// Op2 is the multiply operation
type Op2 struct {
	ArgLen   int64
	ArgModes []int64
}

// NewOp2 returns an instance of the multiply operation
func NewOp2(argLen int64, argModes []int64) *Op2 {
	missingArgModes := argLen - int64(len(argModes))
	if missingArgModes > 0 {
		argModes = append(make([]int64, missingArgModes), argModes...)
	}
	return &Op2{
		ArgLen:   argLen,
		ArgModes: argModes,
	}
}

// NumberToNext gives the number to jump to next instruction
func (op *Op2) NumberToNext() int64 {
	return op.ArgLen
}

// Compute adds two numbers together and stores them
func (op *Op2) Compute(ptr int64, memory []int64) error {
	if int64(len(memory)) < ptr+op.ArgLen {
		return fmt.Errorf("Expected %v chars after ptr position: %v but memory length is %v", op.ArgLen, ptr, len(memory))
	}
	args := make([]int64, op.ArgLen)
	for i := int64(0); i < op.ArgLen; i++ {
		if op.ArgModes[i] == 1 {
			args[i] = memory[ptr+i+1]
		} else {
			args[i] = memory[memory[ptr+i+1]]
		}
	}
	memory[memory[ptr+3]] = args[0] * args[1]
	return nil
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
func (op *Op99) Compute(ptr int64, memory []int64) error {
	return nil
}
