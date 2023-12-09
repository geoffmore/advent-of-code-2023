// https://adventofcode.com/2023/day/1

package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

const file = "input.txt"

// https://github.com/google/re2/wiki/Syntax
// Does not correctly match twoneight as 3 matches, so cannot be used in regexp.FindAll
// const pattern = "zero|one|two|three|four|five|six|seven|eight|nine|[[:digit:]]"

// https://stackoverflow.com/questions/11430863
// Uses lookahead assertion, but is not supported in Golang b/c it isn't guaranteed O(n)
//const pattern = "(?=(zero|one|two|three|four|five|six|seven|eight|nine|[[:digit:]]))"

// Smaller, non-conflicting patterns used instead
const patternA = "one|four|five|six|seven|[[:digit:]]"
const patternB = "two|three|nine"
const patternC = "eight"

var nums = map[string]int{
	"zero":  0,
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
}

type reg []*regexp.Regexp

func (r reg) lineFunc(b []byte) int {
	// TODO - handle newlines so they don't cause panics
	var matches [][]int
	var i int
	var firstIndex, lastIndex = []int{-1, -1}, []int{-1, -1}
	var first, last, sum int
	for _, foo := range r {
		matches = foo.FindAllIndex(b, -1)
		// match starts at index 0 (inclusive) and ends at index 1 (exclusive)
		l := len(matches)
		if l == 0 {
			continue
		} else if l == 1 {
			if matches[0][0] < firstIndex[0] || firstIndex[0] == -1 {
				firstIndex = matches[0]
			}
			if matches[0][0] > lastIndex[0] || lastIndex[0] == -1 {
				lastIndex = matches[0]
			}
			// Check against both
		}
		if l >= 2 {
			if matches[0][0] < firstIndex[0] || firstIndex[0] == -1 {
				firstIndex = matches[0]
			}
			if matches[len(matches)-1][0] > lastIndex[0] || lastIndex[0] == -1 {
				lastIndex = matches[len(matches)-1]
			}
		}
	}
	if firstIndex[0] != -1 {
		i++
	}
	if lastIndex[0] != -1 {
		i++
	}

	first, err := byteoratoi(b[firstIndex[0]:firstIndex[1]])
	p(err)

	if i == 1 {
		sum = 11 * first

	} else {
		last, err = byteoratoi(b[lastIndex[0]:lastIndex[1]])
		p(err)
		sum = 10*first + last

	}
	return sum
}

func main() {
	var total int
	var regexps reg
	regexps = append(regexps, regexp.MustCompile(patternA))
	regexps = append(regexps, regexp.MustCompile(patternB))
	regexps = append(regexps, regexp.MustCompile(patternC))
	// r := regexp.MustCompile(pattern)

	data, err := os.ReadFile(file)
	p(err)

	// https://stackoverflow.com/questions/25691879
	lines := bytes.Split(data, []byte("\n"))
	for _, line := range lines {
		regexps.lineFunc(line)
		total += regexps.lineFunc(line)
	}

	fmt.Println(total)
}

func p(err error) {
	if err != nil {
		log.Panic(err)
	}
}

func b2i(b []byte) (int, error) {
	return strconv.Atoi(string(b))
}

func byteoratoi(b []byte) (int, error) {
	var i int
	s := string(b)
	// Check for word first
	i, ok := nums[s]
	if !ok {
		// If that fails, return the string -> int conversion
		return strconv.Atoi(string(b))
	}
	return i, nil
}
