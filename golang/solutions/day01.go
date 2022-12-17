package solutions

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

type Day1 struct {
	maxCalories []int
}

func newDay1() Problem {
	return &Day1{}
}

func checkMax(maxCalories []int, currentCalories int) {
	minIndex := findMin(maxCalories)
	if maxCalories[minIndex] < currentCalories {
		maxCalories[minIndex] = currentCalories
	}
}

func findMin(maxCalories []int) int {
	min := math.MaxInt
	var minIndex int
	for i := 0; i < len(maxCalories); i++ {
		if maxCalories[i] < min {
			min = maxCalories[i]
			minIndex = i
		}
	}
	return minIndex
}

func (d *Day1) Parse(input string) {
	maxCalories := []int{0, 0, 0}
	var currentCalories int

	for _, s := range strings.Split(input, "\n") {
		if len(s) == 0 {
			checkMax(maxCalories, currentCalories)
			currentCalories = 0
		} else {
			val, _ := strconv.Atoi(s)
			currentCalories += val
		}
	}

	checkMax(maxCalories, currentCalories)
	d.maxCalories = maxCalories
}

func (d *Day1) Part1() {
	max := d.maxCalories[0]
	for i := 1; i < len(d.maxCalories); i++ {
		if d.maxCalories[i] > max {
			max = d.maxCalories[i]
		}
	}
	fmt.Println("part 1)", max)
}

func (d *Day1) Part2() {
	fmt.Println("part 2)", d.maxCalories[0]+d.maxCalories[1]+d.maxCalories[2])
}
