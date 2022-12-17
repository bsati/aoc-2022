package solutions

import (
	"fmt"
	"strings"
)

type Tuple[T any] struct {
	a T
	b T
}

type Day2 struct {
	choices []Tuple[byte]
}

func newDay2() Problem {
	return &Day2{}
}

func (d *Day2) Parse(input string) {
	split := strings.Split(input, "\n")
	choices := []Tuple[byte]{}
	for _, line := range split {
		first := line[0] - 'A'
		second := line[2] - 'X'
		choices = append(choices, Tuple[byte]{first, second})
	}
	d.choices = choices
}

func (d *Day2) Part1() {
	var score int

	for _, choice := range d.choices {
		score += int(choice.b) + 1
		if (choice.b > choice.a && !(choice.b == 2 && choice.a == 0)) || choice.b == 0 && choice.a == 2 {
			score += 6
		} else if choice.b == choice.a {
			score += 3
		}
	}

	fmt.Println("part 1)", score)
}

func (d *Day2) Part2() {
	var score int

	for _, choice := range d.choices {
		opponent := int(choice.a) + 1
		outcome := int(choice.b) + 1
		switch outcome {
		case 1:
			if opponent == 1 {
				score += 3
			} else {
				score += opponent - 1
			}
		case 2:
			score += opponent + 3
		case 3:
			if opponent == 3 {
				score += 1 + 6
			} else {
				score += opponent + 1 + 6
			}
		}
	}

	fmt.Println("part 2)", score)
}
