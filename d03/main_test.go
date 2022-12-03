package main

import (
	"testing"
)

func TestMain(t *testing.T) {
	data := readData("data.txt")

	expected := 157
	result := rucksack(data)

	if expected != result {
		t.Errorf("Priority sum incorrect. Expect %d, got %d\n", expected, result)
	}
}
