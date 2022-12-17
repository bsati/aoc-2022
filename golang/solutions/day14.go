package solutions

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

type Day14 struct {
	positions   map[Coordinate]bool
	min_x       int
	width       int
	height      int
	part1Result int
}

func newDay14() Problem {
	return &Day14{}
}

func (d *Day14) Parse(input string) {
	lines := strings.Split(input, "\n")
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

	d.positions = positions
	d.width = width
	d.height = height
	d.min_x = min_x
	d.printBoard()
}

func (d *Day14) Part1() {
	part1 := d.runner(true)
	d.part1Result = part1
	d.printBoard()
	fmt.Println("part 1)", part1)
}

func (d *Day14) Part2() {
	part2 := d.runner(false)
	d.printBoard()
	fmt.Println("part 2)", d.part1Result+part2+1)
}

func (d *Day14) runner(part1 bool) int {
	placed := 0

	sand := Coordinate{x: 500 - d.min_x, y: 0}

	h := d.height
	if !part1 {
		h += 1
	}

	for {
		newPos, invalid, rest := tryPositions(d.positions, d.width, h, &sand, part1)
		if invalid || !part1 && newPos.x == 500-d.min_x && newPos.y == 0 {
			break
		}
		if rest {
			d.positions[newPos] = true
			sand = Coordinate{x: 500 - d.min_x, y: 0}
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

func (d *Day14) printBoard() {
	min_x := math.MaxInt
	max_x := math.MinInt

	for k := range d.positions {
		if k.x < min_x {
			min_x = k.x
		}
		if k.x > max_x {
			max_x = k.x
		}
	}

	for j := 0; j < d.height; j++ {
		for i := min_x; i <= max_x; i++ {
			if s, ok := d.positions[Coordinate{x: i, y: j}]; ok {
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
