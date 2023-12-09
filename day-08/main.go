package main

import (
	"fmt"
	"github.com/geoffmore/advent-of-code-2023/aoclib"
	"regexp"
)

// const pattern = "([[:alpha:]]{3}) = \\(([[:alpha:]]{3}), ([[:alpha:]]{3})\\)"
const pattern = "([[:alnum:]]{3})"
const file = "input.txt"

const (
	L = iota
	R
)

type problemStruct struct {
	instructions []int // Using iota instead of string here to make lookups easier(?)
	path         map[string][2]string
	startIdxs    []string
	endIdxs      []string
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
		k := string(match[0])

		// Problem B init
		if c := string(k[2]); c == "A" {
			p.startIdxs = append(p.startIdxs, k)
		} else if c == "Z" {
			p.endIdxs = append(p.endIdxs, k)
		}

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
			next = p.path[idx][dir]
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

	p := unmarshalProblem(pattern, file)

	idxs := p.startIdxs
	for {
		for _, dir := range p.instructions {
			// Assign all idxs
			for i := 0; i < len(idxs); i++ {
				// Brute force method. Can be improved by finding the LCD amongst each starting point (kinda)
				idxs[i] = p.path[idxs[i]][dir]
			}

			total++
			if endsWithAll(idxs, "Z") {
				return total
			} else {
				fmt.Println(idxs, total)
			}
		}
	}
	// Return -1 if a path to endIdx could not be found
	return -1
}

func endsWithAll(strs []string, s string) bool {
	if len(strs) == 0 {
		return false
	}
	for _, str := range strs {
		// Check for empty string, then check that the last char in str matches s
		if str == "" || string(str[len(str)-1]) != s {
			return false
		}
	}
	return true
}
