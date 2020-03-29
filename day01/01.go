package day01

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func loadInput() []int64 {
	inputFile, err := os.Open("day01/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer inputFile.Close()

	scanner := bufio.NewScanner(inputFile)
	var input []int64

	for scanner.Scan() {
		number, err := strconv.ParseInt(scanner.Text(), 10, 0)
		if err != nil {
			log.Printf("Could not parse %v as int\n", scanner.Text())
		} else {
			input = append(input, number)
		}
	}

	return input
}

func Part1() int64 {
	input := loadInput()
	var output int64
	for _, module := range input {
		output = output + module/3 - 2
	}

	return output
}

func Part2() int64 {
	input := loadInput()
	var output int64

	for _, module := range input {
		for fuel := module/3 - 2; fuel > 0; fuel = fuel/3 - 2 {
			output = output + fuel
		}
	}

	return output
}
