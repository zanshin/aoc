package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	run()
}

func run() {
	fmt.Println("Run")
	// result := score(readData("data.txt"))
	result := score(readData("input.txt"))
	fmt.Printf("The final score is %d\n", result)

	alternateResult := alternateScore(readData("input.txt"))
	fmt.Printf("The alternate score is %d\n", alternateResult)
}

func score(data []string) int {
	points := 0
	fmt.Printf("File size: %d\n", len(data))
	for row := range data {
		// fmt.Println(data[row])
		// A, X = rock and 1 point
		// B, Y = Paper and 2 points
		// C, Z = Scissors and 3 points
		// 0 points for loss, 3 points for tie, 6 points for win
		switch data[row] {
		case "A X":
			points = points + 4
		case "A Y":
			points = points + 8
		case "A Z":
			points = points + 3
		case "B X":
			points = points + 1
		case "B Y":
			points = points + 5
		case "B Z":
			points = points + 9
		case "C X":
			points = points + 7
		case "C Y":
			points = points + 2
		case "C Z":
			points = points + 6
		}
	}
	return points
}

func alternateScore(data []string) int {
	points := 0
	fmt.Printf("File size: %d\n", len(data))
	for row := range data {
		// fmt.Println(data[row])
		// A = 1 point
		// B = 2 points
		// C = 3 points
		// X = Lose
		// Y = Draw
		// Z = Win
		// 0 points for loss, 3 points for tie, 6 points for win
		switch data[row] {
		case "A X":
			points = points + 3
		case "A Y":
			points = points + 4
		case "A Z":
			points = points + 8
		case "B X":
			points = points + 1
		case "B Y":
			points = points + 5
		case "B Z":
			points = points + 9
		case "C X":
			points = points + 2
		case "C Y":
			points = points + 6
		case "C Z":
			points = points + 7
		}
	}
	return points
}

func readData(filename string) []string {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Printf("Could not open the file due to this %s error \n", err)
	}
	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)

	lines := []string{}
	for fileScanner.Scan() {
		lines = append(lines, fileScanner.Text())
	}
	if err = file.Close(); err != nil {
		fmt.Printf("Could not close the file due to this %s error \n", err)
	}

	return lines
}
