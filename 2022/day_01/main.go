// Package main provides a simple tool to determine the maximum number of
// calories in a sequence.
package main

import (
	"day_01/utils"
	"fmt"
	"strconv"
)

// Main function, called when the program is executed
func main() {
	lines, err := utils.GetInput()
	if err != nil {
		panic(err)
	}
	maxCalories := MostCalories(lines)

	// Print out the maximum number of calories
	fmt.Printf("The Elf carrying the most Calories is carrying: %d Calories\n", maxCalories)
}

// Function to determine the maximum number of calories in a sequence
func MostCalories(lines []string) int {
	// Initialize two variables to hold the maximum and current number of calories
	var maxCalories int
	var currentCalories int

	// Iterate over each line in the lines slice
	for _, line := range lines {
		// If the line is empty, check if the current number of calories is greater than the maximum
		// If it is, set the maximum to the current number
		// Reset the current number of calories to 0
		// Skip to the next iteration
		if line == "" {
			if currentCalories > maxCalories {
				maxCalories = currentCalories
			}
			currentCalories = 0
			continue
		}

		// Try to convert the line to an integer
		// If an error occurs, the application panics
		calories, err := strconv.Atoi(line)
		if err != nil {
			panic(err)
		}

		// Add the number of calories from the line to the current total
		currentCalories += calories
	}

	// After all lines have been processed, check if the current number of calories is greater than the maximum
	// If it is, set the maximum to the current number
	if currentCalories > maxCalories {
		maxCalories = currentCalories
	}

	return maxCalories
}
