package solutions

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

type Day14 struct {
	positions   map[Tuple[int]]bool
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
		path := Path{coordinates: []Tuple[int]{}}
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
			path.coordinates = append(path.coordinates, Tuple[int]{
				a: x,
				b: y,
			})
		}
		paths = append(paths, path)
	}

	height := max_y + 1
	width := max_x - min_x + 1

	positions := map[Tuple[int]]bool{}

	for i := 0; i < len(paths); i++ {
		for j := 0; j < len(paths[i].coordinates); j++ {
			paths[i].coordinates[j].a -= min_x
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

	sand := Tuple[int]{a: 500 - d.min_x, b: 0}

	h := d.height
	if !part1 {
		h += 1
	}

	for {
		newPos, invalid, rest := tryPositions(d.positions, d.width, h, &sand, part1)
		if invalid || !part1 && newPos.a == 500-d.min_x && newPos.b == 0 {
			break
		}
		if rest {
			d.positions[newPos] = true
			sand = Tuple[int]{a: 500 - d.min_x, b: 0}
			placed++
		} else {
			sand = newPos
		}
	}

	return placed
}

// tryPositions tries all possible positions for the current sand location
// and returns the new position, whether the sand falls out (first bool), is at rest (second bool)
func tryPositions(positions map[Tuple[int]]bool, width, height int, sand *Tuple[int], part1 bool) (Tuple[int], bool, bool) {
	var result Tuple[int]
	if !part1 && sand.b+1 == height {
		return *sand, false, true
	}
	if sand.b+1 >= height && part1 {
		return result, true, false
	}
	if _, ok := positions[Tuple[int]{a: sand.a, b: sand.b + 1}]; !ok {
		result.a = sand.a
		result.b = sand.b + 1
		return result, false, false
	}
	if sand.a-1 < 0 && part1 {
		return result, true, false
	}
	if _, ok := positions[Tuple[int]{a: sand.a - 1, b: sand.b + 1}]; !ok {
		result.a = sand.a - 1
		result.b = sand.b + 1
		return result, false, false
	}
	if sand.a+1 >= width && part1 {
		return result, true, false
	}
	if _, ok := positions[Tuple[int]{a: sand.a + 1, b: sand.b + 1}]; !ok {
		result.a = sand.a + 1
		result.b = sand.b + 1
		return result, false, false
	}
	return *sand, false, true
}

type Path struct {
	coordinates []Tuple[int]
}

func (p *Path) getCoordinatePairs() []Tuple[int] {
	result := []Tuple[int]{}
	result = append(result, p.coordinates...)
	if len(result) <= 1 {
		return result
	}
	for i := len(p.coordinates) - 1; i >= 1; i-- {
		target := p.coordinates[i-1]
		current := p.coordinates[i]
		if target.a != current.a {
			dir := -1
			if current.a < target.a {
				dir = 1
			}
			current.a += dir
			for ; current.a != target.a; current.a += dir {
				result = append(result, current)
			}
		} else {
			dir := -1
			if current.b < target.b {
				dir = 1
			}
			current.b += dir
			for ; current.b != target.b; current.b += dir {
				result = append(result, current)
			}
		}
	}
	return result
}

func (d *Day14) printBoard() {
	min_x := math.MaxInt
	max_x := math.MinInt

	for k := range d.positions {
		if k.a < min_x {
			min_x = k.a
		}
		if k.a > max_x {
			max_x = k.a
		}
	}

	for j := 0; j < d.height; j++ {
		for i := min_x; i <= max_x; i++ {
			if s, ok := d.positions[Tuple[int]{a: i, b: j}]; ok {
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
