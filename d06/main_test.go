package main

import (
	"testing"
)

func TestMain(t *testing.T) {
	var tests = []struct {
		input    string
		expected int
	}{
		{"mjqjpqmgbljsphdztnvjfqwrcgsmlb", 7},
		{"bvwbjplbgvbhsrlpgdmjqwftvncz", 5},
		{"nppdvjthqldpwncqszvftbrmjlhg", 6},
		{"nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg", 10},
		{"zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw", 11},
	}

	for _, tt := range tests {
		result := findMarker(tt.input)

		if result != tt.expected {
			t.Errorf("Wrong marker index. Expected %d, got %d\n", tt.expected, result)
		}
	}
}
