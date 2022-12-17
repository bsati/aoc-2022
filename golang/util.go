package main

import (
	"bufio"
	"os"
	"strings"
)

func readInput(filepath string) (string, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return "", err
	}

	scanner := bufio.NewScanner(file)
	var builder strings.Builder

	for scanner.Scan() {
		builder.WriteString(scanner.Text())
		builder.WriteRune('\n')
	}

	result := builder.String()

	return result[0 : len(result)-1], nil
}
