package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	readFile, err := os.Open("input.txt")

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

	maxCalories(data)
}

func maxCalories(data []string) int {
	sum := 0
	most := 0
	for i := range data {
		if data[i] != "" {
			val, _ := strconv.Atoi(data[i])
			sum += val
		} else {
			if sum > most {
				most = sum
			}
			sum = 0
		}
	}
	fmt.Println(most)
	return most
}
