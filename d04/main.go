package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Day Four - Part One")

	ctotal, ototal := run(readData("input.txt"))
	fmt.Printf("The number of contained pairs is %d, and the number of overlaps is %d\n", ctotal, ototal)
}

func run(data []string) (int, int) {
	// fmt.Printf("data: %v\n", data)
	ctotal := 0
	ototal := 0
	for x := range data {
		pairs := strings.Split(data[x], ",")

		c, o := examine(pairs[0], pairs[1])

		if c {
			ctotal = ctotal + 1
		}
		if o {
			ototal = ototal + 1
		}
	}
	return ctotal, ototal
}

func examine(a string, b string) (bool, bool) {
	aPair := strings.Split(a, "-")
	bPair := strings.Split(b, "-")

	// fmt.Printf("apair0: %v  ", aPair[0])
	// fmt.Printf("bpair0: %v\n", bPair[0])
	// fmt.Printf("apair1: %v  ", aPair[1])
	// fmt.Printf("bpair1: %v\n", bPair[1])

	c := contains(aPair[0], aPair[1], bPair[0], bPair[1])
	o := overlaps(aPair[0], aPair[1], bPair[0], bPair[1])

	return c, o
}

func contains(a1, a2, b1, b2 string) bool {
	if between(a1, b1, b2) && between(a2, b1, b2) {
		// fmt.Println("A is contained in B")
		return true
	}
	if between(b1, a1, a2) && between(b2, a1, a2) {
		// fmt.Println("B is contained in A")
		return true
	}

	return false

}

func overlaps(a1, a2, b1, b2 string) bool {
	if between(a1, b1, b2) || between(a2, b1, b2) {
		// fmt.Println("A is contained in B")
		return true
	}
	if between(b1, a1, a2) || between(b2, a1, a2) {
		// fmt.Println("B is contained in A")
		return true
	}
	return false
}

func between(x, y, z string) bool {
	a, _ := strconv.Atoi(x)
	b, _ := strconv.Atoi(y)
	c, _ := strconv.Atoi(z)
	if (b <= a) && (a <= c) {
		return true
	}
	return false
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
