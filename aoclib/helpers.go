package aoclib

import (
	"bytes"
	"log"
	"os"
	"strconv"
)

type Problem func(pattern string, data [][]byte) int

func PanicIf(err error) {
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

func B2i(b []byte) (int, error) {
	return strconv.Atoi(string(b))
}

// ContainCount finds the number of elem in reference that exist in input and returns an int. Limited to comparable because of map generation
func ContainCount[K comparable](input []K, ref []K) int {
	// func ContainCount(input []int, reference []int) int {
	var result int

	// Build a map of desired values
	m := make(map[K]bool)
	for _, v := range ref {
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
