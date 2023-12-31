package main

import (
	"testing"
)

const (
	fileTest  = "sample.txt"
	solutionA = 114
	solutionB = 2
)

func TestProblemA(t *testing.T) {
	if got := ProblemA(pattern, fileTest); got != solutionA {
		t.Errorf("Problem A with sample input should return value '%d'. Got '%d'", solutionA, got)
	}
}

func TestProblemB(t *testing.T) {
	if got := ProblemB(pattern, fileTest); got != solutionB {
		t.Errorf("Problem B with sample input should return value '%d'. Got '%d'", solutionB, got)
	}
}
