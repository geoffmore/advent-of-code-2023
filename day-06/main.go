package main

import (
	"bytes"
	"fmt"
	"github.com/geoffmore/advent-of-code-2023/aoclib"
	"regexp"
)

const pattern = "[[:digit:]]{1,}"
const file = "input.txt"

func main() {
	lines, err := aoclib.ScanLines(file)
	aoclib.PanicIf(err)

	fmt.Println("Problem A: ", ProblemA(pattern, lines))
	fmt.Println("Problem B: ", ProblemB(pattern, lines))
}

// probFn takes a time and distance and calculates the ints that satisfy the condition x * (t - x) > d
func probFn(t int, d int) []int {
	var result []int
	for i := 0; i < t; i++ {
		if i*(t-i) > d {
			result = append(result, i)
		}
	}
	return result
}

func ProblemA(pattern string, data [][]byte) int {
	var total int = 1

	re := regexp.MustCompile(pattern)

	t := re.FindAll(data[0], -1)
	times := aoclib.BytesToInts(t)

	d := re.FindAll(data[1], -1)
	dists := aoclib.BytesToInts(d)

	if len(times) != len(dists) {
		return -1
	}

	for i := 0; i < len(t); i++ {
		total *= len(probFn(times[i], dists[i]))
	}

	return total
}

func ProblemB(pattern string, data [][]byte) int {
	re := regexp.MustCompile(pattern)

	t := re.FindAll(bytes.ReplaceAll(data[0], []byte(" "), nil), -1)
	time, _ := aoclib.B2i(t[0])
	d := re.FindAll(bytes.ReplaceAll(data[1], []byte(" "), nil), -1)
	dist, _ := aoclib.B2i(d[0])

	return len(probFn(time, dist))
}
