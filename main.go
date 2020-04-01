package main

import (
	"aoc/day01"
	"aoc/day02"
	"aoc/day05"
	"aoc/day07"
	"fmt"
	"os"
	"strconv"
)

func main() {
	solutions := map[int64][]int64{
		1: []int64{day01.Part1(), day01.Part2()},
		2: []int64{day02.Part1(), day02.Part2()},
		5: []int64{day05.Part1(), day05.Part2()},
		7: []int64{day07.Part1(), day07.Part2()},
	}

	if len(os.Args) > 1 {
		var day, part int64
		var err error
		day, err = strconv.ParseInt(os.Args[1], 10, 0)
		if err != nil {
			fmt.Printf("Expected day number as first arg, received %v\n", os.Args[1])
			return
		}
		if len(os.Args) > 2 {
			part, err = strconv.ParseInt(os.Args[2], 10, 0)
			if err != nil {
				fmt.Printf("Expected part number as second arg, received %v\n", os.Args[2])
				return
			}
			part = part - 1
			if part < 0 || part > int64(len(solutions[day])) {
				fmt.Printf("Part number out of range, received %v\n", os.Args[2])
				return
			}
			fmt.Printf("day %v part %v: %v\n", day, part+1, solutions[day][part])
		} else {
			for i, sol := range solutions[day] {
				fmt.Printf("day %v part %v: %v\n", day, i+1, sol)
			}
		}
	} else {
		for day, sols := range solutions {
			for i, sol := range sols {
				fmt.Printf("day %02d part %v: %v\n", day, i+1, sol)
			}
		}
	}
	return
}
