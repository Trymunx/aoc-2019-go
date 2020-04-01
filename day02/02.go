package day02

import (
	"aoc/intcode"
)

// Part1 calculates the answer to day 2 part 1
func Part1() int64 {
	input := intcode.LoadFromFile("day02/input.txt")

	// Initialise an incode computer with pointer at position 0
	computer := intcode.NewComputer(input)

	// day 2 part 1 says to replace value at position 1 with 12 and position 2 with 2
	computer.Memory[1] = 12
	computer.Memory[2] = 2

	for !computer.Halted {
		computer.Step(0)
	}

	return computer.Memory[0]
}

// Part2 calculates the answer to day 2 part 2
func Part2() int64 {
	input := intcode.LoadFromFile("day02/input.txt")

	for noun := int64(0); noun < 99; noun++ {
		for verb := int64(0); verb < 99; verb++ {
			computer := intcode.NewComputer(append([]int64{}, input...))
			computer.Memory[1] = noun
			computer.Memory[2] = verb
			for !computer.Halted {
				computer.Step(0)
			}
			if computer.Memory[0] == 19690720 {
				return 100*noun + verb
			}
		}
	}
	return 0
}
