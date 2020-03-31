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

// Compute adds two numbers together and stores them
func (op *Op1) Compute(pc *Computer, input int64) (int64, error) {
	if err := checkMemoryRange(pc, op.ArgLen); err != nil {
		return 0, err
	}
	args := calculateArgModes(op.ArgLen-1, op.ArgModes, pc.ptr, pc.Memory)
	pc.Memory[pc.Memory[pc.ptr+3]] = args[0] + args[1]
	pc.ptr += op.ArgLen + 1
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

// Compute adds two numbers together and stores them
func (op *Op2) Compute(pc *Computer, input int64) (int64, error) {
	if err := checkMemoryRange(pc, op.ArgLen); err != nil {
		return 0, err
	}
	args := calculateArgModes(op.ArgLen-1, op.ArgModes, pc.ptr, pc.Memory)
	pc.Memory[pc.Memory[pc.ptr+3]] = args[0] * args[1]
	pc.ptr += op.ArgLen + 1
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

// Compute inputs a number in a given position
func (op *Op3) Compute(pc *Computer, input int64) (int64, error) {
	if err := checkMemoryRange(pc, op.ArgLen); err != nil {
		return 0, err
	}
	pc.Memory[pc.Memory[pc.ptr+1]] = input
	pc.ptr += op.ArgLen + 1
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

// Compute outputs a number from a position or value
func (op *Op4) Compute(pc *Computer, input int64) (int64, error) {
	if err := checkMemoryRange(pc, op.ArgLen); err != nil {
		return 0, err
	}
	args := calculateArgModes(op.ArgLen, op.ArgModes, pc.ptr, pc.Memory)
	pc.ptr += op.ArgLen + 1
	return args[0], nil
}

// Op5 jumps the pointer if first arg is non-zero
type Op5 struct {
	ArgLen   int64
	ArgModes []int64
}

// NewOp5 returns an instance of the jump if true operation
func NewOp5(argLen int64, argModes []int64) *Op5 {
	return &Op5{
		ArgLen:   argLen,
		ArgModes: padMissingArgModes(argLen, argModes),
	}
}

// Compute jumps the pointer to the new position
func (op *Op5) Compute(pc *Computer, input int64) (int64, error) {
	if err := checkMemoryRange(pc, op.ArgLen); err != nil {
		return 0, err
	}
	args := calculateArgModes(op.ArgLen, op.ArgModes, pc.ptr, pc.Memory)
	if args[0] != 0 {
		pc.ptr = args[1]
	} else {
		pc.ptr += op.ArgLen + 1
	}
	return 0, nil
}

// Op6 jumps the pointer if first arg is zero
type Op6 struct {
	ArgLen   int64
	ArgModes []int64
}

// NewOp6 returns an instance of the jump if true operation
func NewOp6(argLen int64, argModes []int64) *Op6 {
	return &Op6{
		ArgLen:   argLen,
		ArgModes: padMissingArgModes(argLen, argModes),
	}
}

// Compute jumps the pointer to the new position
func (op *Op6) Compute(pc *Computer, input int64) (int64, error) {
	if err := checkMemoryRange(pc, op.ArgLen); err != nil {
		return 0, err
	}
	args := calculateArgModes(op.ArgLen, op.ArgModes, pc.ptr, pc.Memory)
	if args[0] == 0 {
		pc.ptr = args[1]
	} else {
		pc.ptr += op.ArgLen + 1
	}
	return 0, nil
}

// Op7 is the less than operator
type Op7 struct {
	ArgLen   int64
	ArgModes []int64
}

// NewOp7 returns an instance of the less than operation
func NewOp7(argLen int64, argModes []int64) *Op7 {
	return &Op7{
		ArgLen:   argLen,
		ArgModes: padMissingArgModes(argLen, argModes),
	}
}

// Compute sets its write pointer to 1 if arg 0 is less than arg 1
func (op *Op7) Compute(pc *Computer, input int64) (int64, error) {
	if err := checkMemoryRange(pc, op.ArgLen); err != nil {
		return 0, err
	}
	args := calculateArgModes(op.ArgLen, op.ArgModes, pc.ptr, pc.Memory)
	if args[0] < args[1] {
		pc.Memory[pc.Memory[pc.ptr+3]] = 1
	} else {
		pc.Memory[pc.Memory[pc.ptr+3]] = 0
	}
	pc.ptr += op.ArgLen + 1
	return 0, nil
}

// Op8 is the equals operator
type Op8 struct {
	ArgLen   int64
	ArgModes []int64
}

// NewOp8 returns an instance of the equals operation
func NewOp8(argLen int64, argModes []int64) *Op8 {
	return &Op8{
		ArgLen:   argLen,
		ArgModes: padMissingArgModes(argLen, argModes),
	}
}

// Compute sets its write pointer to 1 if arg 0 equals arg 1
func (op *Op8) Compute(pc *Computer, input int64) (int64, error) {
	if err := checkMemoryRange(pc, op.ArgLen); err != nil {
		return 0, err
	}
	args := calculateArgModes(op.ArgLen, op.ArgModes, pc.ptr, pc.Memory)
	if args[0] == args[1] {
		pc.Memory[pc.Memory[pc.ptr+3]] = 1
	} else {
		pc.Memory[pc.Memory[pc.ptr+3]] = 0
	}
	pc.ptr += op.ArgLen + 1
	return 0, nil
}

// Op99 is the halt operation
type Op99 struct {
	ArgLen int64
}

// NewOp99 returns an instance of the halt instruction
func NewOp99() *Op99 {
	return &Op99{0}
}

// Compute terminates the program
func (op *Op99) Compute(pc *Computer, input int64) (int64, error) {
	pc.Halted = true
	return 0, nil
}

func checkMemoryRange(pc *Computer, argLen int64) error {
	if int64(len(pc.Memory)) < pc.ptr+argLen {
		return fmt.Errorf("Expected %v chars after ptr position: %v but memory length is %v", argLen, pc.ptr, len(pc.Memory))
	}
	return nil
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

func padMissingArgModes(argLen int64, argModes []int64) []int64 {
	missing := argLen - int64(len(argModes))
	if missing > 0 {
		return append(argModes, make([]int64, missing)...)
	}
	return argModes
}
