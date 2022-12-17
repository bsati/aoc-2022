package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/bsati/aoc-2022/golang/solutions"
)

func main() {
	args := os.Args

	if len(args) < 2 {
		panic("Expected at least one argument (day)")
	}

	day, err := strconv.Atoi(args[1])

	if err != nil {
		panic("Supplied day is not an integer")
	}

	solver := solutions.GetSolver(day)

	if solver == nil {
		panic("Solver for specified day has not been implemented")
	}

	filepath := fmt.Sprintf("../inputs/day%02d.txt", day)
	input, err := readInput(filepath)

	if err != nil {
		log.Panicf("Error occured reading input: %v", err)
	}

	solver.Parse(input)
	solver.Part1()
	solver.Part2()
}
