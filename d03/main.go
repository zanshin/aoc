package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("Day Three - Part One")

	fmt.Printf("Priority total: %d\n", rucksack(readData("input.txt")))
}

func rucksack(data []string) int {
	// a-z = 1 - 26, A-Z = 27 - 52
	priorities := map[string]int{}
	alphabet := "abcdefghijklmnopqrstuvwxyz"
	for letter := range alphabet {
		priorities[string(alphabet[letter])] = letter + 1
		priorities[strings.ToUpper(string(alphabet[letter]))] = letter + 27
	}

	total := 0
	for item := range data {
		half := len(data[item]) / 2
		first := data[item][:half]
		second := data[item][half:]
		// fmt.Printf("%s %s %s\n", data[item], first, second)

	outside:
		for x := 0; x < len(first); x++ {
			for y := 0; y < len(second); y++ {
				if first[x] == second[y] {
					// found the matching item, convert to value
					if v, found := priorities[string(first[x])]; found {
						fmt.Printf("Matching value found: %s - %d\n", string(first[x]), v)
						total = total + v
					}
					break outside
				}
			}
		}
	}
	return total
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
