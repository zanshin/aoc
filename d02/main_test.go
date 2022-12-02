package main

import (
	"fmt"
	"testing"
)

func TestRun(t *testing.T) {
	fmt.Println("Testing Run")
	data := readData("data.txt")

	expected := 15
	result := score(data)

	if expected != result {
		t.Errorf("Score is incorrect. Expected %d, got %d\n", expected, result)
	}

}
