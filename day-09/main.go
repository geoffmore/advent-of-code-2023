package main

import (
	"fmt"
	"github.com/geoffmore/advent-of-code-2023/aoclib"
	"regexp"
)

const pattern = "(-?[[:digit:]]{1,})"
const file = "input.txt"

func main() {
	lines, err := aoclib.ScanLines(file)
	aoclib.PanicIf(err)

	fmt.Println("Problem A: ", ProblemA(pattern, file))
	fmt.Println("Problem B: ", ProblemB(pattern, lines))
}

// problem stores a set of numbers in elem 0 and their diffs in successive slices
type problem [][]int

// isValid checks if each successive array is 1 smaller than the last and whether the last index array contains all 0s
func (p problem) isValid() bool {
	return false
}

func (p problem) isComplete() bool {
	if len(p) == 0 {
		// Guessing that an empty line isn't complete
		return false
	}
	// Lazy way to check completeness by just looking at the last index
	for _, i := range p[len(p)-1] {
		if i != 0 {
			return false
		}
	}
	return true
}

// problemStruct contains all problem read from input file
type problemStruct struct {
	problems []problem
}

func diffs(nums []int) []int {
	var result []int
	if len(nums) == 1 {
		return result
	}
	for i := 1; i < len(nums); i++ {
		result = append(result, nums[i]-nums[i-1])
	}
	return result
}

// unmarshalProblem takes a file and regex pattern and results an array of problem
func unmarshalProblem(pattern, file string) problemStruct {
	var p problemStruct

	lines, err := aoclib.ScanLines(file)
	aoclib.PanicIf(err)
	re := regexp.MustCompile(pattern)

	for _, line := range lines {
		v := aoclib.BytesToInts(re.FindAll(line, -1))
		p.problems = append(p.problems, problem{v})
	}
	return p
}

func funcA(p problemStruct) int {
	var result int
	for i := 0; i < len(p.problems); i++ {
		var j int
		for !p.problems[i].isComplete() {
			p.problems[i] = append(p.problems[i], diffs(p.problems[i][j]))
			j++
		}
	}
	// each problem is complete now
	for _, prob := range p.problems {
		zeroesIdx := len(prob) - 1
		lastIdx := len(prob[zeroesIdx]) - 1
		// v should always start at 0
		v := prob[zeroesIdx][lastIdx]
		// Add the last idx of each above row to v
		for i := len(prob) - 1; i >= 0; i-- {
			// Add the last num of current index of prob[i] asc
			v += prob[i][len(prob[i])-1]
		}
		result += v
	}
	return result
}

func ProblemA(pattern string, file string) int {
	p := unmarshalProblem(pattern, file)
	return funcA(p)
}

func ProblemB(pattern string, data [][]byte) int {
	var total int
	re := regexp.MustCompile(pattern)

	for _, line := range data {
		_, _ = line, re
	}
	return total
}
