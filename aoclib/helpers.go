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
