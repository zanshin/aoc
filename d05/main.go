package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type item struct {
	value string
	next  *item
}

type Stack struct {
	top  *item
	size int
}

func (stack *Stack) Len() int {
	return stack.size
}

func (stack *Stack) Push(value string) {
	stack.top = &item{
		value: value,
		next:  stack.top,
	}
	stack.size++
}

func (stack *Stack) Pop() (value string) {
	if stack.Len() > 0 {
		value = stack.top.value
		stack.top = stack.top.next
		stack.size--
		return
	}

	return ""
}

func (stack *Stack) Peek() {
	fmt.Printf("[%s]\n", stack.top.value)
}

func (stack *Stack) Contents() {
	item := stack.top
	for x := 0; x < stack.Len(); x++ {
		fmt.Printf("[%s] ", item.value)
		item = item.next
	}
	fmt.Println("")
}

var (
	cargo         = make([]Stack, 3)
	separator int = 0
	cargoBase int = 0
	moveStart int = 0
	stackNums int = 0
	numStacks int = 0
)

func main() {
	fmt.Println("AOC Day 5")

	crateMover9001 := true
	// run("input.txt", crateMover9001)
	fmt.Printf("Top crates: %s\n", run("input.txt", crateMover9001))
}

func run(filename string, crateMover9001 bool) string {
	data := readData(filename)
	separator = findSeparator(data)
	cargoBase = separator - 2
	stackNums = separator - 1
	moveStart = separator + 1

	numStacks = len(strings.ReplaceAll(data[stackNums], " ", ""))
	fmt.Printf("Number of stacks: %d\n", numStacks)
	cargo = make([]Stack, numStacks)

	initializeStacks(data)
	organizeCrates(data, crateMover9001)
	topCrates := ""
	for c := 0; c < numStacks; c++ {
		topCrates += string(cargo[c].Pop())
	}

	return topCrates

}

func initializeStacks(data []string) {
	for line := cargoBase; line >= 0; line-- {
		buildStacks(data[line])
	}

	// for x := 0; x < numStacks; x++ {
	// 	fmt.Printf("Cargo stack %d: ", x)
	// 	for cargo[x].Len() > 0 {
	// 		fmt.Printf("%s ", cargo[x].Pop())
	// 	}
	// 	fmt.Printf("\n")
	// }

	cargoContents()
}

func cargoContents() {
	for x := 0; x < numStacks; x++ {
		fmt.Printf("Cargo stack %d: ", x)
		cargo[x].Contents()
	}
}
func organizeCrates(data []string, crateMover9001 bool) {
	fmt.Println("")
	fmt.Println("organizeCrates")
	fmt.Printf("line: %d, length: %d\n", moveStart, len(data))
	for line := moveStart; line < len(data); line++ {
		getMoves(data[line], crateMover9001)
	}
	fmt.Println("done with moves")
}

func findSeparator(data []string) int {
	base := 0
	for line := range data {
		if data[line] == "" {
			base = line
		}
	}
	return base
}

func buildStacks(row string) {
	// len(row) - 1 = position of last crate
	// space between crates = 4
	// len(row) / 4 = number of cargo stacks
	// Don't save " " (blank) as a crate
	for x := 0; x <= (len(row) / 4); x++ {
		if (string(row[(x*4)+1])) != " " {
			cargo[x].Push(string(row[(x*4)+1]))
		}
	}
}

func getMoves(row string, crateMover9001 bool) {
	// move x from X to X
	// first x can be 1 or 2 digits
	// from and to will be 1 or 2 or 3
	// TODO: parse instructions into count, from, and to
	// TODO: call popPush
	fmt.Printf("\nmoves: %s\n", row)
	var re = regexp.MustCompile(`move (\d{1,2}) from (\d) to (\d)`)
	c, _ := strconv.Atoi(re.ReplaceAllString(row, "${1}"))
	f, _ := strconv.Atoi(re.ReplaceAllString(row, "${2}"))
	t, _ := strconv.Atoi(re.ReplaceAllString(row, "${3}"))
	// fmt.Printf("c, f, t: %d, %d, %d\n", c, f, t)
	popPush(c, f, t, crateMover9001)

}

func popPush(count, from, to int, crateMover9001 bool) {
	fmt.Printf("count: %d, from: %d, to: %d. CrateMover9001: %v\n", count, from, to, crateMover9001)
	craneStack := make([]Stack, 1)
	for x := 1; x <= count; x++ {
		fmt.Printf("x: %d\n", x)
		// fmt.Printf("crate: %q\n", cargo[from-1].Pop())
		cargo[from-1].Peek()
		if crateMover9001 {
			craneStack[0].Push(cargo[from-1].Pop())
		} else {
			cargo[to-1].Push(cargo[from-1].Pop())
		}
	}
	if crateMover9001 {
		for craneStack[0].Len() > 0 {
			cargo[to-1].Push(craneStack[0].Pop())
		}
	}
	cargoContents()
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
