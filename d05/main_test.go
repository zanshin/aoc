package main

import (
	"testing"
)

func TestMain(t *testing.T) {
	// data := readData("data.txt")
	data := "data.txt"

	expected := "CMZ"
	result := run(data)

	if expected != result {
		t.Errorf("Top of stack contents incorrect. Expected %s, got %s\n", expected, result)
	}
}
