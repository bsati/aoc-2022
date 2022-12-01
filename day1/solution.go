package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	file, _ := os.Open("./input1.txt")
	scanner := bufio.NewScanner(file)

	maxCalories := []int{0, 0, 0}
	var currentCalories int

	for scanner.Scan() {
		if scanner.Text() == "" {
			checkMax(maxCalories, currentCalories)
			currentCalories = 0
		} else {
			val, _ := strconv.Atoi(scanner.Text())
			currentCalories += val
		}
	}

	checkMax(maxCalories, currentCalories)

	fmt.Println(maxCalories)
	fmt.Println(maxCalories[0] + maxCalories[1] + maxCalories[2])

	file.Close()
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
