package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Position struct {
	y int
	x int
}

func (p *Position) Move(direction Position) {
	p.x += direction.x
	p.y += direction.y
}

func (p *Position) Add(other *Position) Position {
	return Position{p.x + other.x, p.y + other.y}
}

func NewSet[T comparable]() Set[T] {
	return make(map[T]struct{})
}

type Set[T comparable] map[T]struct{}

func (s Set[T]) Add(value T) bool {
	new := true
	if _, ok := s[value]; ok {
		new = false
	}
	s[value] = struct{}{}
	return new
}

func (s Set[T]) Contains(value T) bool {
	_, ok := s[value]
	return ok
}

func clampAxis(d int) int {
	if d > 1 {
		return 1
	} else if d < -1 {
		return -1
	}
	return d
}

func clamp(dx, dy int) Position {
	return Position{clampAxis(dy), clampAxis(dx)}
}

func main() {
	file, _ := os.Open("./input.txt")
	scanner := bufio.NewScanner(file)

	lines := []string{}

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	directionMap := map[byte]Position{
		'R': {0, 1},
		'L': {0, -1},
		'U': {1, 0},
		'D': {-1, 0},
	}

	visitedPositionsPart1 := NewSet[Position]()
	visitedPositionsPart2 := NewSet[Position]()
	currentDir := directionMap[lines[0][0]]
	counter, _ := strconv.Atoi(lines[0][2:])
	lineIndex := 1
	knots := []Position{}
	for k := 0; k < 10; k++ {
		knots = append(knots, Position{0, 0})
	}
	for {
		for ; counter > 0; counter-- {
			knots[0].Move(currentDir)

			for i := 1; i < len(knots); i++ {
				dx := knots[i-1].x - knots[i].x
				dy := knots[i-1].y - knots[i].y

				if !(dx > 1 || dx < -1 || dy > 1 || dy < -1) {
					continue
				}
				knots[i].Move(clamp(dx, dy))
				if i == 1 {
					visitedPositionsPart1.Add(knots[i])
				} else if i == 9 {
					visitedPositionsPart2.Add(knots[i])
				}
			}
		}
		if lineIndex == len(lines) {
			break
		}
		currentDir = directionMap[lines[lineIndex][0]]
		counter, _ = strconv.Atoi(lines[lineIndex][2:])
		lineIndex++
	}

	// for j := 20; j >= -20; j-- {
	// 	for i := -20; i <= 20; i++ {
	// 		if visitedPositionsPart2.Contains(Position{j, i}) || i == 0 && j == 0 {
	// 			fmt.Print("#")
	// 		} else {
	// 			fmt.Print(".")
	// 		}
	// 	}
	// 	fmt.Print("\n")
	// }

	fmt.Printf("part 1) %d\n", len(visitedPositionsPart1)+1)
	fmt.Printf("part 2) %d", len(visitedPositionsPart2)+1)
}
