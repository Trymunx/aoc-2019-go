package opcodes

// Opcode is a specific operation
type Opcode struct {
	code    int64
	instant bool
}

// Op1 is the add operation
type Op1 struct {
	code    int64
	instant bool
}

func (op *Op1) Next() Opcode {

}
