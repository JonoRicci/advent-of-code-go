package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	var maxCalories int
	var currentCalories int
	for _, line := range lines {
		if line == "" {
			if currentCalories > maxCalories {
				maxCalories = currentCalories
			}
			currentCalories = 0
			continue
		}
		calories, err := strconv.Atoi(line)
		if err != nil {
			panic(err)
		}
		currentCalories += calories
	}

	if currentCalories > maxCalories {
		maxCalories = currentCalories
	}

	fmt.Printf("The Elf carrying the most Calories is carrying: %d Calories\n", maxCalories)
}