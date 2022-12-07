package main

import (
	"testing"
)

func TestMain(t *testing.T) {
	// data := readData("data.txt")
	data := "data.txt"

	crateMover9001 := false
	expected := "CMZ"
	result := run(data, crateMover9001)

	if expected != result {
		t.Errorf("Top of stack contents incorrect. Expected %s, got %s\n", expected, result)
	}

	crateMover9001 = true
	expected = "MCD"
	result = run(data, crateMover9001)

	if expected != result {
		t.Errorf("Top of stack contents incorrct. Expected %s, got %s\n", expected, result)
	}
}
