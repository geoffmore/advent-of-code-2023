package main

import (
	"testing"
)

const (
	fileTest  = "sample.txt"
	BTest     = "sample2.txt"
	solutionA = 2
	solutionB = 6 // CHANGEME
)

func TestProblemA(t *testing.T) {
	if got := ProblemA(pattern, fileTest); got != solutionA {
		t.Errorf("Problem A with sample input should return value '%d'. Got '%d'", solutionA, got)
	}
}

func TestProblemB(t *testing.T) {
	if got := ProblemB(pattern, BTest); got != solutionB {
		t.Errorf("Problem B with sample input should return value '%d'. Got '%d'", solutionB, got)
	}
}
