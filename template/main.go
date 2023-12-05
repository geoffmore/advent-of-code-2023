package main

import (
	"fmt"
	"github.com/geoffmore/advent-of-code-2023/aoclib"
	"regexp"
)

const pattern = "CHANGEME"
const file = "input.txt"

func main() {
	lines, err := aoclib.Scan(file)
	aoclib.PanicIf(err)

	fmt.Println("Problem A: ", ProblemA(pattern, lines))
	fmt.Println("Problem B: ", ProblemB(pattern, lines))
}

func ProblemA(pattern string, data [][]byte) int {
	var total int
	re := regexp.MustCompile(pattern)

	for _, line := range data {
		_, _ = line, re
	}
	return total
}

func ProblemB(pattern string, data [][]byte) int {
	var total int
	re := regexp.MustCompile(pattern)

	for _, line := range data {
		_, _ = line, re
	}
	return total
}
