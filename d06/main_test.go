package main

import (
	"testing"
)

func TestMain(t *testing.T) {
	var tests = []struct {
		input   string
		packet  int
		message int
	}{
		{"mjqjpqmgbljsphdztnvjfqwrcgsmlb", 7, 19},
		{"bvwbjplbgvbhsrlpgdmjqwftvncz", 5, 23},
		{"nppdvjthqldpwncqszvftbrmjlhg", 6, 23},
		{"nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg", 10, 29},
		{"zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw", 11, 26},
	}

	for _, tt := range tests {
		result := findMarker(tt.input, 4)

		if result != tt.packet {
			t.Errorf("Wrong marker index. Expected %d, got %d\n", tt.packet, result)
		}

		messageStart := findMarker(tt.input, 14)
		if messageStart != tt.message {
			t.Errorf("Wrong message index. Expected %d, got %d\n", tt.message, messageStart)
		}
	}
}
