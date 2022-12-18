package solutions

import (
	"fmt"
	"math"
	"strings"

	go_ds "github.com/bsati/go-ds"
)

type Grid struct {
	points go_ds.Set[Tuple[int]]
	min_x  int
	max_x  int
	min_y  int
	max_y  int
}

func manhattanDistance(p1, p2 *Tuple[int]) int {
	return abs(p1.a-p2.a) + abs(p1.b-p2.b)
}

func (g *Grid) checkPositions(sensorX, sensorY, beaconX, beaconY int) {
	x := []int{sensorX, beaconX}
	y := []int{sensorY, beaconY}

	for _, posX := range x {
		if posX < g.min_x {
			g.min_x = posX
		}
		if posX > g.max_x {
			g.max_x = posX
		}
	}

	for _, posY := range y {
		if posY < g.min_y {
			g.min_y = posY
		}
		if posY > g.max_y {
			g.max_y = posY
		}
	}
}

type Sensor struct {
	position          Tuple[int]
	closestBeacon     Tuple[int]
	manhattanDistance int
}

type Day15 struct {
	grid    Grid
	sensors []Sensor
}

func newDay15() Problem {
	return &Day15{
		grid: Grid{
			min_x: math.MaxInt,
			max_x: math.MinInt,
			min_y: math.MaxInt,
			max_y: math.MinInt,
		},
	}
}

func (d *Day15) Parse(input string) {
	lines := strings.Split(input, "\n")
	sensors := []Sensor{}
	for _, line := range lines {
		var sensorX, sensorY, beaconX, beaconY int
		fmt.Sscanf(line, "Sensor at x=%d, y=%d: closest beacon is at x=%d, y=%d", &sensorX, &sensorY, &beaconX, &beaconY)
		d.grid.checkPositions(sensorX, sensorY, beaconX, beaconY)
		sensor := Tuple[int]{sensorX, sensorY}
		beacon := Tuple[int]{beaconX, beaconY}
		sensors = append(sensors, Sensor{
			position:          sensor,
			closestBeacon:     beacon,
			manhattanDistance: manhattanDistance(&sensor, &beacon),
		})
	}
	d.sensors = sensors
}

func (d *Day15) Part1() {
	row := 2000000

	positions := go_ds.NewSet[Tuple[int]]()

	for _, sensor := range d.sensors {
		offset := abs(row - sensor.position.b)
		if offset > sensor.manhattanDistance {
			continue
		}

		for i := sensor.position.a - sensor.manhattanDistance + offset; i <= sensor.position.a+sensor.manhattanDistance-offset; i++ {
			pos := Tuple[int]{i, row}
			positions.Add(pos)
		}
	}

	fmt.Println("part 1)", len(positions))
}

func (d *Day15) free(position *Tuple[int]) bool {
	for _, s := range d.sensors {
		if manhattanDistance(&s.position, position) <= s.manhattanDistance {
			return false
		}
	}
	return true
}

func (d *Day15) Part2() {
	min := 0
	max := 4000000

	signs := []Tuple[int]{
		{-1, -1},
		{-1, 1},
		{1, -1},
		{1, 1},
	}

	for _, sensor := range d.sensors {
		for dx := 0; dx <= sensor.manhattanDistance+1; dx++ {
			dy := (sensor.manhattanDistance + 1) - dx
			for _, sign := range signs {
				x := sensor.position.a + dx*sign.a
				y := sensor.position.b + dy*sign.b
				if x < min || x > max || y < min || y > max {
					continue
				}
				testPos := Tuple[int]{x, y}
				if d.free(&testPos) {
					fmt.Printf("x: %d, y: %d\n", x, y)
					fmt.Println("part 2)", x*4000000+y)
					return
				}
			}
		}
	}
}
