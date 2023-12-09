package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

// const file = "sample.txt"

const file = "input.txt"

const patternA = "([[:digit:]]{1,})"

// https://github.com/bufbuild/protoc-gen-validate/issues/68
const patternB = "[*]"

// Point represents a cartesian coordinate
type Point [2]int

// Range represents a horizontal range of Point(s) as [x, y)
type Range [2]int

type Foo struct {
	y    int
	span Range
}

func ContainsPoint(points []Point, point Point) bool {
	for _, p := range points {
		if p == point {
			return true
		}
	}
	return false
}

func (f Foo) adjacency() []Point {
	var points, result, adj []Point

	for i := f.span[0]; i < f.span[1]; i++ {
		points = append(points, Point{i, f.y})
	}
	// Lazy adjacency calculation that doesn't de-duplicate, which also has an effect of overlapping numbers
	for _, point := range points {
		adj = append(adj, point.adjacency()...)
	}
	// Compare all adjacent points and add to result if they are not part of points or exist in map
	// https://stackoverflow.com/questions/66643946
	found := make(map[Point]bool)
	for _, point := range adj {
		if _, ok := found[point]; !ok {
			if !ContainsPoint(points, point) {
				result = append(result, point)
				found[point] = true
			}
		}
	}
	return result
}

func (f Foo) adjacencyFn(fn func(p Point) bool) []Point {
	var points, result, adj []Point

	for i := f.span[0]; i < f.span[1]; i++ {
		points = append(points, Point{i, f.y})
	}
	// Lazy adjacency calculation that doesn't de-duplicate, which also has an effect of overlapping numbers
	for _, point := range points {
		adj = append(adj, point.adjacency()...)
	}
	// Compare all adjacent points and add to result if they are not part of points or exist in map
	// https://stackoverflow.com/questions/66643946
	found := make(map[Point]bool)
	for _, point := range adj {
		if _, ok := found[point]; !ok {
			if fn(point) {
				if !ContainsPoint(points, point) {
					result = append(result, point)
					found[point] = true
				}
			}
		}
	}
	return result
}

func (r Range) adjacency() []Point {
	var result []Point

	return result

}

// adjacency generates the []Point surrounding a Point
func (p Point) adjacency() []Point {
	// func adjacency(point Point) []Point {
	var result []Point
	for i := p[0] - 1; i <= p[0]+1; i++ {
		for j := p[1] - 1; j <= p[1]+1; j++ {
			// Exclude the stuff at the point
			if !(p[0] == i && p[1] == j) {
				// if i != point[0] && j != point[1] {
				result = append(result, Point{i, j})
			}
		}
	}
	return result
}

func filterCharsA(s []string) bool {
	var result []string
	filter := map[string]bool{
		"1": false,
		"2": false,
		"3": false,
		"4": false,
		"5": false,
		"6": false,
		"7": false,
		"8": false,
		"9": false,
		"0": false,
		".": false,
	}
	for _, c := range s {
		if _, ok := filter[c]; !ok {
			result = append(result, c)
		}
	}
	return len(result) > 0
}

func p03a(re *regexp.Regexp, data [][]byte) int {
	var total int

	// Collect all numbers as []Point
	for i, line := range data {
		// Get the relative indices of numbers
		nums := re.FindAllIndex(line, -1)
		for _, num := range nums {
			// Gather adjacent points
			adjacentPoints := Foo{
				y:    i,
				span: Range{num[0], num[1]},
			}.adjacency()

			chars := UniqAdjacentChars(data, adjacentPoints)

			// Filter func for p03a
			if filterCharsA(chars) {
				n, err := strconv.Atoi(string(data[i][num[0]:num[1]]))
				p(err)
				total += n
			}
		}

	}
	return total
}

func p03b(pattern string, data [][]byte) int {
	var total int

	re := regexp.MustCompile(pattern)

	// Collect all numbers as []Point
	m := make(map[Point][]int)
	for i, line := range data {
		// Get the relative indices of numbers
		nums := re.FindAllIndex(line, -1)
		for _, num := range nums {
			// Gather adjacent points
			adjacentPoints := Foo{
				y:    i,
				span: Range{num[0], num[1]},
			}.adjacency()

			// Iterate over adjacentPoints and return Points where they match string
			filtered := Data(data).filterPoints(adjacentPoints, "*")

			// Gather results by point
			for _, point := range filtered {
				n, err := strconv.Atoi(string(data[i][num[0]:num[1]]))
				p(err)
				if _, ok := m[point]; !ok {
					m[point] = make([]int, 0)
				}
				m[point] = append(m[point], n)
			}
		}
	}
	for _, nums := range m {
		if len(nums) == 2 {
			total += nums[0] * nums[1]
		}
	}
	return total
}

func main() {

	re := regexp.MustCompile(patternA)

	lines, err := Scan(file)
	p(err)

	fmt.Println("A: ", p03a(re, lines))
	fmt.Println("B: ", p03b(patternA, lines))
}

func p(err error) {
	if err != nil {
		log.Panic(err)
	}
}

// Scan takes a file name, opens it, and returns the data from the file split by newlines
func Scan(file string) ([][]byte, error) {
	data, err := os.ReadFile(file)
	if err != nil {
		return nil, err
	}
	return bytes.Split(data, []byte("\n")), nil
}

// UniqAdjacentChars returns unique characters in data surrounding points. Is safe with OOB indices on points
func UniqAdjacentChars(data [][]byte, points []Point) []string {
	var result []string
	// Contains found strings
	m := make(map[string]bool)
	l, w := len(data), len(data[0])

	for _, point := range points {
		if (point[0] < w && point[0] >= 0) && (point[1] < l && point[1] >= 0) {
			s := string(data[point[1]][point[0]])
			if _, ok := m[s]; !ok {
				result = append(result, s)
				m[s] = true
			}
		}
	}
	return result
}

type Data [][]byte

func (d Data) filterPoints(points []Point, s string) []Point {
	var result []Point
	// Contains found strings
	m := make(map[string]bool)
	l, w := len(d), len(d[0])
	for _, point := range points {
		if (point[0] < w && point[0] >= 0) && (point[1] < l && point[1] >= 0) {
			c := string(d[point[1]][point[0]])
			if _, ok := m[s]; !ok {
				if c == s {
					result = append(result, point)
					m[s] = true
				}
			}
		}
		/*
			if true {
				v := string(d[point[1]][point[0]])
				_ = v
				//if d[point[1]][point[0]] == []byte(s) {
				result = append(result, point)
			}
		*/
	}
	return result
}
