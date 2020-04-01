package day07

import (
	"aoc/intcode"
	"fmt"
)

// Part1 calculates the answer to day 7 part 1
func Part1() int64 {
	input := intcode.LoadFromFile("day07/input.txt")

	permutations := calculatePermutations()

	var finalOutput int64
	// Input for the first instruction, the phase setting from 0 to 4
	for _, permutation := range permutations {

		pcs := make([]*intcode.Computer, 5)
		for i := range pcs {
			pcs[i] = intcode.NewComputer(append([]int64{}, input...))
		}

		for i, pc := range pcs {
			pc.Step(permutation[i])
		}

		var loopOutputSignal int64
		for _, pc := range pcs {
			for !pc.Halted {
				output, _ := pc.Step(loopOutputSignal)
				if output == -1 {
					fmt.Println("Received error code from computer.Step(), printing state:")
					fmt.Println(intcode.PrintStatus(pc))
					return -1
				} else if output != 0 {
					loopOutputSignal = output
				}
			}
		}
		if loopOutputSignal > finalOutput {
			finalOutput = loopOutputSignal
		}
	}

	return finalOutput
}

// Part2 calculates the answer to day 7 part 2
func Part2() int64 {
	input := intcode.LoadFromFile("day07/input.txt")
	permutations := calculatePermutations()

	var finalOutput int64
	// Input for the first instruction, the phase setting from 0 to 4
	for _, permutation := range permutations {
		pcs := make([]*intcode.Computer, 5)
		for i := range pcs {
			pcs[i] = intcode.NewComputer(append([]int64{}, input...))
		}
		for i, pc := range pcs {
			pc.Step(permutation[i] + 5)
		}

		var loopInputSignal int64
		var loopOutputSignal int64
		finished := false
		for !finished {
			for _, pc := range pcs {
				loopOutputSignal = 0
				hasOutput := false
				// continue running until the pc outputs something, then pass it on
				for !hasOutput && !pc.Halted {
					output, ok := pc.Step(loopInputSignal)
					if output == -1 {
						fmt.Println("Received error code from computer.Step(), printing state:")
						fmt.Println(intcode.PrintStatus(pc))
						return -1
					}
					if ok {
						hasOutput = ok
						loopOutputSignal = output
					}
				}
				loopInputSignal = loopOutputSignal
				finished = pc.Halted
			}
			if loopOutputSignal > finalOutput {
				finalOutput = loopOutputSignal
			}
		}
	}

	return finalOutput
}

func calculatePermutations() [][]int64 {
	permutations := [][]int64{}
	for a := int64(0); a < 5; a++ {
		for b := int64(0); b < 5; b++ {
			if b == a {
				continue
			}
			for c := int64(0); c < 5; c++ {
				if c == a || c == b {
					continue
				}
				for d := int64(0); d < 5; d++ {
					if d == a || d == b || d == c {
						continue
					}
					for e := int64(0); e < 5; e++ {
						if e == a || e == b || e == c || e == d {
							continue
						}
						permutations = append(permutations, []int64{a, b, c, d, e})
					}
				}
			}
		}
	}
	return permutations
}
