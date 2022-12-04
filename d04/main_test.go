package main

import (
	"testing"
)

func TestMain(t *testing.T) {
	data := readData("data.txt")

	expected := 2
	overlapExpected := 4
	result, overlapResult := run(data)

	if expected != result {
		t.Errorf("Number of contained pairs incorrect. Expected %d, got %d\n", expected, result)
	}

	if overlapExpected != overlapResult {
		t.Errorf("Number of overlaps is incorrect. Expected %d, got %d\n", overlapExpected, overlapResult)
	}
}
