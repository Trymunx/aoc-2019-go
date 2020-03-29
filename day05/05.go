package day05

import "aoc/intcode"

// Part1 calculates the answer to day 5 part 1
func Part1() int64 {
	input := intcode.LoadFromFile("day05/input.txt")

	// Initialise an incode computer with pointer at position 0
	computer := intcode.NewComputer(input, 0, false)

	for !computer.Halted {
		computer.Step()
	}

	return computer.Memory[0]
}
