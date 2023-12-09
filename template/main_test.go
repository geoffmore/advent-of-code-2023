package main

import (
	"github.com/geoffmore/advent-of-code-2023/aoclib"
	"testing"
)

const (
	fileTest  = "sample.txt"
	solutionA = 0 // CHANGEME
	solutionB = 0 // CHANGEME
)

func TestProblemA(t *testing.T) {
	lines, err := aoclib.ScanLines(fileTest)
	if err != nil {
		t.Errorf(err.Error())
	}
	if got := ProblemA(pattern, lines); got != solutionA {
		t.Errorf("Problem A with sample input should return value '%d'. Got '%d'", solutionA, got)
	}
}

func TestProblemB(t *testing.T) {
	lines, err := aoclib.ScanLines(fileTest)
	if err != nil {
		t.Errorf(err.Error())
	}
	if got := ProblemB(pattern, lines); got != solutionB {
		t.Errorf("Problem B with sample input should return value '%d'. Got '%d'", solutionB, got)
	}
}
