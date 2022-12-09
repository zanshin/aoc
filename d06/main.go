package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Day Six- Part One")

	marker := run("input.txt")
	fmt.Printf("The marker index is %d\n", marker)
}

func run(filename string) int {
	data := readData(filename)
	result := 0
	for line := range data {
		result = findMarker(data[line])
		fmt.Printf("result: %d\n", result)
	}
	return result
}

func findMarker(data string) int {
	// fmt.Printf("\ndata: %s length: %d\n", data, len(data))
	// var found bool = false
	// result := 0
	for p := 4; p < len(data); p++ {
		// fmt.Printf("\nindex: %d\n", p)
		if compareAll(string(data[p-4 : p])) {
			// fmt.Printf("Marker is %d\n", p)
			return p
		}
	}
	return 0
}

func compareAll(data string) bool {
	// fmt.Printf("inspecting: %s\n", data)
	x := 0
	if data[x] == data[x+1] {
		return false
	}
	if data[x] == data[x+2] {
		return false
	}
	if data[x] == data[x+3] {
		return false
	}

	x = 1
	if data[x] == data[x+1] {
		return false
	}
	if data[x] == data[x+2] {
		return false
	}

	x = 2
	if data[x] == data[x+1] {
		return false
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
