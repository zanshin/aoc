package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("Day Three - Part One")
	initPriorities()

	fmt.Printf("Priority total: %d\n", rucksack(readData("input.txt")))
	fmt.Printf("Badge total: %d\n", badges(readData("input.txt")))
}

var priorities = map[string]int{}

func initPriorities() {
	// a-z = 1 - 26, A-Z = 27 - 52
	alphabet := "abcdefghijklmnopqrstuvwxyz"
	for letter := range alphabet {
		priorities[string(alphabet[letter])] = letter + 1
		priorities[strings.ToUpper(string(alphabet[letter]))] = letter + 27
	}

}

func rucksack(data []string) int {
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

func badges(data []string) int {
	// Need to read two lines (a and B) call intersection and get c'
	// then read third line (c) and call intersetcion (c, c') to get answer
	// convert answer to value and add to total
	// repeat until end of file
	total := 0
	for x := 0; x < len(data); x = x + 3 {
		// fmt.Printf("x: %d\n", x)
		a := strings.Split(data[x], "")
		b := strings.Split(data[x+1], "")
		c := strings.Split(data[x+2], "")
		// fmt.Printf("a: %s\n", a)
		// fmt.Printf("b: %s\n", b)
		// fmt.Printf("c: %s\n", c)

		ab := intersection(a, b)
		ans := intersection(ab, c)
		// fmt.Printf("ab: %s, ans: %s\n", ab, ans[0])

		if v, found := priorities[ans[0]]; found {
			// fmt.Printf("Common element: %s - %d\n", strings.Join(ans, ""), v)
			total = total + v
		}
	}
	return total
}

func intersection(a, b []string) (c []string) {
	m := make(map[string]bool)

	for _, item := range a {
		m[item] = true
	}

	for _, item := range b {
		if _, ok := m[item]; ok {
			c = append(c, item)
		}
	}
	return
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
