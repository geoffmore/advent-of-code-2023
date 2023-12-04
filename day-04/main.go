package main

import (
	"fmt"
	"github.com/geoffmore/advent-of-code-2023/aoclib"
	"math"
	"regexp"
	"strconv"
	"strings"
)

const pattern = "Card[[:blank:]]{1,}(?P<cardNum>[[:digit:]]{1,}): (?P<winningNumsStr>.*) \\| (?P<inputNumsStr>.*)"
const file = "04_input.txt"

func main() {
	lines, err := aoclib.Scan(file)
	aoclib.PanicIf(err)

	fmt.Println(ProblemA(pattern, lines))
}

func ProblemA(pattern string, data [][]byte) int {
	var total int
	re := regexp.MustCompile(pattern)

	for _, line := range data {
		var points = 0
		// points := 1 * 2^^len(intersection - 1)
		groups := aoclib.PatternToMap(re, line)
		_ = groups
		// TODO - improve this by putting something into aoclib
		winningNums := strings2Ints(strings.Split(string(groups["winningNumsStr"]), " "))
		inputNums := strings2Ints(strings.Split(string(groups["inputNumsStr"]), " "))

		c := foo(inputNums, winningNums)
		if c > 0 {
			// Subtract 1 from c to account for the initial base of 1 in points
			c -= 1
			points = 1 * int(math.Pow(2, float64(c)))
		}
		total += points
	}
	return total
}

func strings2Ints(strs []string) []int {
	var result []int

	// Ideally a filtered list of input strings

	for _, s := range strs {
		if s != "" {
			v, err := strconv.Atoi(s)
			if err == nil {
				result = append(result, v)
			}
		}
	}

	return result
}

// TODO - make this compare on strings instead of ints since equality is checked and not values
// TODO - make this generic
// foo returns how many values in input exist in reference
func foo(input []int, reference []int) int {
	var result int

	// Build a map of desired values
	// TODO - generic type should be comparable
	m := make(map[int]bool)
	for _, v := range reference {
		m[v] = true
	}
	// Iterate over input and increment if a match is found
	for _, v := range input {
		if m[v] {
			result++
		}
	}
	return result
}
