package main

import (
	"github.com/geoffmore/advent-of-code-2023/aoclib"
	"testing"
)

const (
	fileTest  = "sample.txt"
	solutionA = 35
	solutionB = 0 // CHANGEME
)

func TestProblemA(t *testing.T) {
	data, err := aoclib.Scan(file)
	if err != nil {
		t.Errorf(err.Error())
	}
	if got := ProblemA(data); got != solutionA {
		t.Errorf("Problem A with sample input should return value '%d'. Got '%d'", solutionA, got)
	}
}

func TestProblemB(t *testing.T) {
	data, err := aoclib.Scan(file)
	if err != nil {
		t.Errorf(err.Error())
	}
	if got := ProblemB(data); got != solutionB {
		t.Errorf("Problem A with sample input should return value '%d'. Got '%d'", solutionA, got)
	}
}
