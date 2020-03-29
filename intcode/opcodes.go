package intcode

import (
	"fmt"
)

// Op1 is the add operation
type Op1 struct {
	ArgLen   int64
	ArgModes []int64
	Relative bool
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
	memory[memory[ptr+3]] = memory[memory[ptr+1]] + memory[memory[ptr+2]]
	return nil
}

// Op2 is the add operation
type Op2 struct {
	ArgLen   int64
	ArgModes []int64
	Relative bool
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
	memory[memory[ptr+3]] = memory[memory[ptr+1]] * memory[memory[ptr+2]]
	return nil
}

// Op99 is the add operation
type Op99 struct {
	ArgLen   int64
	ArgModes []int64
	Relative bool
}

// NumberToNext gives the number to jump to next instruction
func (op *Op99) NumberToNext() int64 {
	return op.ArgLen
}

// Compute terminates the program
func (op *Op99) Compute(ptr int64, memory []int64) error {
	return nil
}
