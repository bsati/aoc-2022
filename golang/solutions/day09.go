package solutions

import (
	"fmt"
	"strconv"
	"strings"

	go_ds "github.com/bsati/go-ds"
)

type Day9 struct {
	lines                 []string
	visitedPositionsPart2 go_ds.Set[Tuple[int]]
}

func newDay9() Problem {
	return &Day9{
		visitedPositionsPart2: go_ds.NewSet[Tuple[int]](),
	}
}

func (d *Day9) Parse(input string) {
	d.lines = strings.Split(input, "\n")
}

func (d *Day9) Part1() {
	directionMap := map[byte]Tuple[int]{
		'R': {0, 1},
		'L': {0, -1},
		'U': {1, 0},
		'D': {-1, 0},
	}

	visitedPositionsPart1 := go_ds.NewSet[Tuple[int]]()
	currentDir := directionMap[d.lines[0][0]]
	counter, _ := strconv.Atoi(d.lines[0][2:])
	lineIndex := 1
	knots := []Tuple[int]{}
	for k := 0; k < 10; k++ {
		knots = append(knots, Tuple[int]{0, 0})
	}
	for {
		for ; counter > 0; counter-- {
			knots[0].Move(currentDir)

			for i := 1; i < len(knots); i++ {
				dx := knots[i-1].a - knots[i].a
				dy := knots[i-1].b - knots[i].b

				if !(dx > 1 || dx < -1 || dy > 1 || dy < -1) {
					continue
				}
				knots[i].Move(clamp(dx, dy))
				if i == 1 {
					visitedPositionsPart1.Add(knots[i])
				} else if i == 9 {
					d.visitedPositionsPart2.Add(knots[i])
				}
			}
		}
		if lineIndex == len(d.lines) {
			break
		}
		currentDir = directionMap[d.lines[lineIndex][0]]
		counter, _ = strconv.Atoi(d.lines[lineIndex][2:])
		lineIndex++
	}

	fmt.Println("part 1)", len(visitedPositionsPart1)+1)
}

func (d *Day9) Part2() {
	fmt.Println("part 2)", len(d.visitedPositionsPart2)+1)
}

func (p *Tuple[int]) Move(direction Tuple[int]) {
	p.a += direction.a
	p.b += direction.b
}

func (p *Tuple[int]) Add(other *Tuple[int]) Tuple[int] {
	return Tuple[int]{p.a + other.a, p.b + other.b}
}

func clampAxis(d int) int {
	if d > 1 {
		return 1
	} else if d < -1 {
		return -1
	}
	return d
}

func clamp(dx, dy int) Tuple[int] {
	return Tuple[int]{clampAxis(dy), clampAxis(dx)}
}
