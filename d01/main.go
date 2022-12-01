package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	readFile, err := os.Open("input.txt")
	// readFile, err := os.Open("data.txt")

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	data := []string{}
	for fileScanner.Scan() {
		data = append(data, fileScanner.Text())
	}

	readFile.Close()

	most, top := maxCalories(data)
	fmt.Printf("The most any elf has is: %d. The top three elves have: %d\n", most, top)
}

func maxCalories(data []string) (int, int) {
	sum := 0
	most := 0
	elves := []int{}
	for i := range data {
		// fmt.Printf("Processing: %s\n", data[i])
		if data[i] != "" {
			val, _ := strconv.Atoi(data[i])
			sum += val
		} else {
			elves = append(elves, sum)
			if sum > most {
				most = sum
				sum = 0
			} else {
				sum = 0
			}
		}
	}
	// fmt.Printf("Last elf: %d\n", sum)
	elves = append(elves, sum)

	// fmt.Printf("Elves: %v\n", elves)
	one, two, three := topThree(elves)
	// fmt.Printf("Most: %d\n", most)
	// fmt.Printf("Top three: %d\n", one+two+three)
	// fmt.Printf("One: %d\n", one)
	// fmt.Printf("Two: %d\n", two)
	// fmt.Printf("Three: %d\n", three)

	return most, one + two + three
}

func topThree(elves []int) (int, int, int) {
	sort.Slice(elves, func(i, j int) bool {
		return elves[i] > elves[j]
	})

	// fmt.Printf("Sorted: %d\n", elves

	return elves[0], elves[1], elves[2]

}
