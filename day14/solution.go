package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type Path struct {
	coordinates []Coordinate
}

func (p *Path) getCoordinatePairs() []Coordinate {
	result := []Coordinate{}
	result = append(result, p.coordinates...)
	if len(result) <= 1 {
		return result
	}
	for i := len(p.coordinates) - 1; i >= 1; i-- {
		target := p.coordinates[i-1]
		current := p.coordinates[i]
		if target.x != current.x {
			dir := -1
			if current.x < target.x {
				dir = 1
			}
			current.x += dir
			for ; current.x != target.x; current.x += dir {
				result = append(result, current)
			}
		} else {
			dir := -1
			if current.y < target.y {
				dir = 1
			}
			current.y += dir
			for ; current.y != target.y; current.y += dir {
				result = append(result, current)
			}
		}
	}
	return result
}

type Coordinate struct {
	x int
	y int
}

func printBoard(positions map[Coordinate]bool, width, height int) {
	min_x := math.MaxInt
	max_x := math.MinInt

	for k := range positions {
		if k.x < min_x {
			min_x = k.x
		}
		if k.x > max_x {
			max_x = k.x
		}
	}

	for j := 0; j < height; j++ {
		for i := min_x; i <= max_x; i++ {
			if s, ok := positions[Coordinate{x: i, y: j}]; ok {
				if s {
					fmt.Print("o")
				} else {
					fmt.Print("#")
				}
			} else {
				fmt.Print(".")
			}
		}
		fmt.Print("\n")
	}
}

func main() {
	file, _ := os.Open("./input.txt")
	scanner := bufio.NewScanner(file)

	lines := []string{}

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	min_x := 1000
	max_x := 0
	max_y := 0
	paths := []Path{}
	for _, line := range lines {
		path := Path{coordinates: []Coordinate{}}
		split := strings.Split(line, " -> ")
		for _, coordinate := range split {
			xySplit := strings.Split(coordinate, ",")
			x, _ := strconv.Atoi(xySplit[0])
			y, _ := strconv.Atoi(xySplit[1])
			if x < min_x {
				min_x = x
			}
			if x > max_x {
				max_x = x
			}
			if y > max_y {
				max_y = y
			}
			path.coordinates = append(path.coordinates, Coordinate{
				x: x,
				y: y,
			})
		}
		paths = append(paths, path)
	}

	height := max_y + 1
	width := max_x - min_x + 1

	positions := map[Coordinate]bool{}

	for i := 0; i < len(paths); i++ {
		for j := 0; j < len(paths[i].coordinates); j++ {
			paths[i].coordinates[j].x -= min_x
		}
	}

	for _, p := range paths {
		for _, c := range p.getCoordinatePairs() {
			positions[c] = false
		}
	}

	printBoard(positions, width, height)

	part1 := runner(positions, min_x, width, height, true)
	printBoard(positions, width, height)
	part2 := part1 + runner(positions, min_x, width, height, false) + 1
	printBoard(positions, width, height+1)
	fmt.Println("part 1) ", part1)
	fmt.Println("part 2) ", part2)
}
func runner(positions map[Coordinate]bool, min_x, width, height int, part1 bool) int {
	placed := 0

	sand := Coordinate{x: 500 - min_x, y: 0}

	h := height
	if !part1 {
		h += 1
	}

	for {
		newPos, invalid, rest := tryPositions(positions, width, h, &sand, part1)
		if invalid || !part1 && newPos.x == 500-min_x && newPos.y == 0 {
			break
		}
		if rest {
			positions[newPos] = true
			sand = Coordinate{x: 500 - min_x, y: 0}
			placed++
		} else {
			sand = newPos
		}
	}

	return placed
}

// tryPositions tries all possible positions for the current sand location
// and returns the new position, whether the sand falls out (first bool), is at rest (second bool)
func tryPositions(positions map[Coordinate]bool, width, height int, sand *Coordinate, part1 bool) (Coordinate, bool, bool) {
	var result Coordinate
	if !part1 && sand.y+1 == height {
		return *sand, false, true
	}
	if sand.y+1 >= height && part1 {
		return result, true, false
	}
	if _, ok := positions[Coordinate{x: sand.x, y: sand.y + 1}]; !ok {
		result.x = sand.x
		result.y = sand.y + 1
		return result, false, false
	}
	if sand.x-1 < 0 && part1 {
		return result, true, false
	}
	if _, ok := positions[Coordinate{x: sand.x - 1, y: sand.y + 1}]; !ok {
		result.x = sand.x - 1
		result.y = sand.y + 1
		return result, false, false
	}
	if sand.x+1 >= width && part1 {
		return result, true, false
	}
	if _, ok := positions[Coordinate{x: sand.x + 1, y: sand.y + 1}]; !ok {
		result.x = sand.x + 1
		result.y = sand.y + 1
		return result, false, false
	}
	return *sand, false, true
}
