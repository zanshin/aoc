package main

import (
	"fmt"
	"testing"
)

func TestMain(t *testing.T) {
	fmt.Println("TestMain")
	data := []string{"1000", "2000", "3000", "", "4000", "", "5000", "6000", "", "7000", "8000", "9000", "", "10000"}
	expected := 24000
	topExpected := 45000

	most, top := maxCalories(data)
	if most != expected {
		t.Fatalf("Most calories not correct. Expected %d, got %d", expected, most)
	}

	if top != topExpected {
		t.Fatalf("Top three combined not correct. Expected %d, got %d\n", topExpected, top)
	}
}
