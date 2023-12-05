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

	fmt.Println("Problem A: ", ProblemA(pattern, lines))
	fmt.Println("Problem B: ", ProblemB(pattern, lines))
}

func ProblemA(pattern string, data [][]byte) int {
	var total int
	re := regexp.MustCompile(pattern)

	for _, line := range data {
		var points = 0
		// points := 1 * 2^^len(intersection - 1)
		groups := aoclib.PatternToMap(re, line)
		// TODO - improve this by putting something into aoclib
		winningNums := strings2Ints(strings.Split(string(groups["winningNumsStr"]), " "))
		inputNums := strings2Ints(strings.Split(string(groups["inputNumsStr"]), " "))

		c := aoclib.ContainCount(inputNums, winningNums)
		if c > 0 {
			// Subtract 1 from c to account for the initial base of 1 in points
			c -= 1
			points = 1 * int(math.Pow(2, float64(c)))
		}
		total += points
	}
	return total
}

func ProblemB(pattern string, data [][]byte) int {
	var total int
	// Maps card num to next card nums, which can be used for lookups
	var m = make(map[int][]int)
	var called = make(map[int]int) // map of how many times a key, cardNum calls other things (including itself)

	// Iterate on every card traversal
	re := regexp.MustCompile(pattern)

	// Likely better to iterate backwards because we can use the last line to avoid calculating the first line recursively. Should be similar to Pascal numbers
	for i := len(data); i > 0; i-- {
		line := data[i-1]
		groups := aoclib.PatternToMap(re, line)
		// cardNum needs to be an int because of matching logic
		cardNum, err := aoclib.B2i(groups["cardNum"])
		aoclib.PanicIf(err)
		// TODO - improve this by putting something into aoclib
		winningNums := strings2Ints(strings.Split(string(groups["winningNumsStr"]), " "))
		inputNums := strings2Ints(strings.Split(string(groups["inputNumsStr"]), " "))

		c := aoclib.ContainCount(inputNums, winningNums)

		// Recursive
		/*
			f(1) = 1 + f(2) + f(3) + f(4) + f(5)
			f(2) = 1 + f(3) + f(4)
			f(3) = 1 + f(4) + f(5)
			f(4) = 1 + f(5)
			f(5) = 1
			f(6) = 1
		*/

		// Iterative
		/* 	m[n] = 1
		 	m[n-1] = 1
		 	...
			m[6] = 1
			m[5] = 1
			m[4] = 1 + m[5] = 3
			m[3] = 1 + m[4] + m[5] = 5
			m[2] = 1 + m[3] + m[4] = 8
			m[1] = 1 + m[2] + m[3] + m[4] + m[5] = 1+ 8 + 5 + 3 + 1 = 36
			then subtract len(m) to get 30
		*/
		called[cardNum]++ // set the called count for the current number to 1 (from 0)

		if _, ok := m[cardNum]; !ok {
			m[cardNum] = []int{}      // Alternative to make([]int, c) because 2nd arg is length (w/ 0 val), not cap
			for j := c; j >= 1; j-- { // Find the biggest number first, because that is guaranteed to be in the map
				newNum := cardNum + j
				m[cardNum] = append(m[cardNum], newNum)
			}
			for j := c; j >= 1; j-- { // All of these nums exist in m because num > cardNum and cardNum goes n -> 1
				num := cardNum + j
				// TODO - get help understanding how I arrived at this
				called[cardNum] += called[num]
			}
		}
	}

	for _, v := range called {
		total += v
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
