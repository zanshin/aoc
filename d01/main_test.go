package main

import (
	"fmt"
	"testing"
)

func TestMain(t *testing.T) {
	fmt.Println("TestMain")
	data := []string{"1000", "2000", "3000", "", "4000", "5000", "6000", "", "7000", "8000", "9000", "", "10000"}
	expected := 24000

	result := maxCalories(data)
	if result != expected {
		t.Fatalf("Result not correct. Expected %d, got %d", expected, result)
	}
}
