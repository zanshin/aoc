package main

import (
	"testing"
)

func TestMain(t *testing.T) {
	initPriorities()
	data := readData("data.txt")

	expected := 157
	result := rucksack(data)

	if expected != result {
		t.Errorf("Priority sum incorrect. Expect %d, got %d\n", expected, result)
	}

	expectedBadge := 70
	badgeResult := badges(data)

	if expectedBadge != badgeResult {
		t.Errorf("Badge sum incorrect. Expected %d, got %d\n", expectedBadge, badgeResult)
	}
}
