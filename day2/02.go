package day2

import (
	"aoc/intcode"
)

// Part1 calculates the answer to day 2 part 1
func Part1() int64 {
	input := intcode.LoadFromFile("day2/input.txt")

	// Initialise an incode computer with pointer at position 0
	computer := intcode.NewComputer(input, 0, false)

	// day 2 part 1 says to replace value at position 1 with 12 and position 2 with 2
	computer.Memory[1] = 12
	computer.Memory[2] = 2

	for !computer.Halted {
		computer.Step()
	}

	return computer.Memory[0]
}

// Part2 calculates the answer to day 2 part 2
func Part2() int {
	input := intcode.LoadFromFile("day2/input.txt")

	for noun := 0; noun < 99; noun++ {
		for verb := 0; verb < 99; verb++ {
			computer := intcode.NewComputer(append([]int64{}, input...), 0, false)
			computer.Memory[1] = int64(noun)
			computer.Memory[2] = int64(verb)
			for !computer.Halted {
				computer.Step()
			}
			if computer.Memory[0] == 19690720 {
				return 100*noun + verb
			}
		}
	}
	return 0
}
