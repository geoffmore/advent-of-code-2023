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
	p := unmarshalProblem(pattern, file)
	return p.pathLen("AAA")
}

func ProblemB(pattern string, file string) int {
	p := unmarshalProblem(pattern, file)

	idxs := p.startIdxs
	iters := make([]int, len(idxs))
	bases := make([]int, len(idxs))

	for i, idx := range idxs {
		// coincidence - iters[i] is always prime and len(p.instructions) is always prime, so LCD never works
		iters[i] = p.pathLen(idx)
		bases[i] = iters[i] / len(p.instructions) // This integer division makes no sense
	}

	v := len(p.instructions)
	for _, base := range bases {
		// Accommodate smaller datasets
		if base != len(p.instructions) {
			v *= base
		}
	}
	return v

	/*
		for {
			for _, dir := range p.instructions {
				// Assign all idxs
				for i := 0; i < len(idxs); i++ {
					// Brute force method. Can be improved by finding the LCD amongst each starting point (kinda)
					idxs[i] = p.path[idxs[i]][dir]
				}

				total++
				if c := endsWithAll(idxs, "Z"); c == len(idxs) {
					return total
				} else {
					if c > len(idxs)-2 {
						fmt.Printf("Current indexes are '%v'. Total is %d. Current matches are %d/%d\n", idxs, total, c, len(idxs))
					}
				}
			}
		}
	*/
}

// endsWith checks if string s ends with string c (assuming c is a char)
func endsWith(s, c string) bool {
	// Check for empty string, then check that the last char in str matches s
	return !(s == "" || string(s[len(s)-1]) != c)
}

func (p *problemStruct) pathLen(initialKey string) int {

	// initialKey not in map. Unable to iterate
	if _, ok := p.path[initialKey]; !ok {
		return -1
	}
	var total int
	startIdx := initialKey

	idx := startIdx
	var next string
	for {
		for _, dir := range p.instructions {
			next = p.path[idx][dir]
			idx = next

			total++
			if endsWith(idx, "Z") {
				return total
			}
		}
	}
}
