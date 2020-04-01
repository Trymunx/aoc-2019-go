package day05

import (
	"aoc/intcode"
	"fmt"
)

// Part1 calculates the answer to day 5 part 1
func Part1() int64 {
	input := intcode.LoadFromFile("day05/input.txt")

	// Initialise an incode computer with pointer at position 0
	computer := intcode.NewComputer(input)

	var diagnosticCode int64

	for !computer.Halted {
		output, _ := computer.Step(1)
		if output == -1 {
			fmt.Println("Received error code from computer.Step(), printing state:")
			fmt.Println(intcode.PrintStatus(computer))
			return -1
		}
		if output != 0 && output != -1 {
			diagnosticCode = output
		}
	}

	return diagnosticCode
}

// Part2 calculates the answer to day 5 part 2
func Part2() int64 {
	input := intcode.LoadFromFile("day05/input.txt")

	// Initialise an incode computer with pointer at position 0
	computer := intcode.NewComputer(input)

	var diagnosticCode int64

	for !computer.Halted {
		output, _ := computer.Step(5)
		if output == -1 {
			fmt.Println("Received error code from computer.Step(), printing state:")
			fmt.Println(intcode.PrintStatus(computer))
			return -1
		}
		if output != 0 && output != -1 {
			diagnosticCode = output
		}
	}

	return diagnosticCode
}
