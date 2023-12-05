package main

import (
	"github.com/geoffmore/advent-of-code-2023/aoclib"
	"testing"
)

const (
	fileTest  = "04_sample.txt"
	solutionA = 13
	solutionB = 30
)

func TestProblemA(t *testing.T) {
	lines, err := aoclib.Scan(fileTest)
	if err != nil {
		t.Errorf(err.Error())
	}
	if got := ProblemA(pattern, lines); got != solutionA {
		t.Errorf("Problem A with sample input should return value '%d'. Got '%d'", solutionA, got)
	}
}

func TestProblemB(t *testing.T) {
	lines, err := aoclib.Scan(fileTest)
	if err != nil {
		t.Errorf(err.Error())
	}
	if got := ProblemB(pattern, lines); got != solutionB {
		t.Errorf("Problem A with sample input should return value '%d'. Got '%d'", solutionA, got)
	}
}
