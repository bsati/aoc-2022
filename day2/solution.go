package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	file, _ := os.Open("./input1.txt")
	scanner := bufio.NewScanner(file)

	a(scanner)
	file.Seek(0, 0)
	scanner = bufio.NewScanner(file)
	b(scanner)
}

func a(scanner *bufio.Scanner) {
	var score int

	for scanner.Scan() {
		choices := strings.Split(scanner.Text(), " ")
		opponent := choices[0][0] - 'A'
		self := choices[1][0] - 'X'
		score += int(self) + 1
		if (self > opponent && !(self == 2 && opponent == 0)) || self == 0 && opponent == 2 {
			score += 6
		} else if self == opponent {
			score += 3
		}
	}

	fmt.Printf("a): %d\n", score)
}

func b(scanner *bufio.Scanner) {
	var score int

	for scanner.Scan() {
		lineSplit := strings.Split(scanner.Text(), " ")
		opponent := int(lineSplit[0][0]-'A') + 1
		outcome := int(lineSplit[1][0]-'X') + 1
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

	fmt.Printf("b): %d\n", score)
}
