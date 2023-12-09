package main

import (
	"fmt"
	"github.com/geoffmore/advent-of-code-2023/aoclib"
	"regexp"
)

// const pattern = "([[:alpha:]]{3}) = \\(([[:alpha:]]{3}), ([[:alpha:]]{3})\\)"
const pattern = "([[:alpha:]]{3})"
const file = "input.txt"

const (
	L = iota
	R
)

type problemStruct struct {
	instructions []int // Using iota instead of string here to make lookups easier(?)
	path         map[string][2]string
}

func main() {

	fmt.Println("Problem A: ", ProblemA(pattern, file))
	fmt.Println("Problem B: ", ProblemB(pattern, file))
}

func unmarshalProblem(pattern, file string) problemStruct {
	var p problemStruct

	lines, err := aoclib.ScanLines(file)
	aoclib.PanicIf(err)
	re := regexp.MustCompile(pattern)

	// Grab the instructions
	for _, c := range lines[0] {
		if string(c) == "L" {
			p.instructions = append(p.instructions, L)

		} else if string(c) == "R" {
			p.instructions = append(p.instructions, R)
		}
	}

	// Map the path for each line
	p.path = make(map[string][2]string)
	for _, line := range lines[2:] {
		match := re.FindAll(line, -1)
		p.path[string(match[0])] = [2]string{string(match[1]), string(match[2])}
	}
	return p

}

func ProblemA(pattern string, file string) int {
	var total int

	startIdx, endIdx := "AAA", "ZZZ"
	p := unmarshalProblem(pattern, file)

	idx := startIdx
	var next string
	for {
		for _, dir := range p.instructions {
			a := p.path[idx]
			next = a[dir]
			idx = next

			total++
			if idx == endIdx {
				return total
			}
		}
	}
	// Return -1 if a path to endIdx could not be found
	return -1
}

func ProblemB(pattern string, file string) int {
	var total int
	re := regexp.MustCompile(pattern)

	for _, line := range file {
		_, _ = line, re
	}
	return total
}
