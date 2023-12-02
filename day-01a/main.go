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

const file = "01a_input.txt"
const pattern = "[[:digit:]]"

func main() {
	var sum int
	var matches [][]byte
	var i, j int

	r := regexp.MustCompile(pattern)

	data, err := os.ReadFile(file)
	p(err)

	// https://stackoverflow.com/questions/25691879
	lines := bytes.Split(data, []byte("\n"))
	for _, line := range lines {
		matches = r.FindAll(line, -1)
		if len(matches) == 0 {
			continue
		}

		i, err = b2i(matches[0])
		p(err)

		if len(matches) == 1 {
			// use the number twice
			// 10^1 * i + 1 * i = 11 * i
			sum += 11 * i
		} else {
			j, err = b2i(matches[len(matches)-1])
			p(err)
			sum += 10*i + j
		}

		_, _ = i, j
	}

	fmt.Println(sum)
}

func p(err error) {
	if err != nil {
		log.Panic(err)
	}
}

func b2i(b []byte) (int, error) {
	return strconv.Atoi(string(b))
}
