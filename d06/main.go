package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Day Six- Part One")

	marker := run("input.txt", 4)
	fmt.Printf("The packet start index is %d\n", marker)

	marker = run("input.txt", 14)
	fmt.Printf("The message start index is %d\n", marker)
}

func run(filename string, size int) int {
	data := readData(filename)
	result := 0
	for line := range data {
		result = findMarker(data[line], size)
		fmt.Printf("result: %d\n", result)
	}
	return result
}

func findMarker(data string, size int) int {
	// fmt.Printf("\ndata: %s length: %d\n", data, len(data))
	for p := size; p < len(data); p++ {
		// fmt.Printf("\nindex: %d\n", p)
		if compareAll(string(data[p-size:p]), size) {
			// fmt.Printf("Marker is %d\n", p)
			return p
		}
	}
	return 0
}

func compareAll(data string, size int) bool {
	// fmt.Printf("inspecting: %s\n", data)

	for p := size - 1; p >= 0; p-- {
		for q := 0; q < p; q++ {
			// fmt.Printf("data[0]: %v, data[(size-1)-p]: %v\n", data[0], data[size-1-p])
			if data[p] == data[q] {
				return false
			}
		}
	}

	return true
}

func readData(filename string) []string {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Printf("could not open file due to this error: %s\n", err)
	}

	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)

	lines := []string{}
	for fileScanner.Scan() {
		lines = append(lines, fileScanner.Text())
	}

	if err = file.Close(); err != nil {
		fmt.Printf("could not close file due to this error: %s\n", err)
	}

	return lines
}
