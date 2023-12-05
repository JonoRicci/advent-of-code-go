// Package main solves the Advent of Code 2023 Day 01 problem.
package main

import (
	"jonoricci/advent-of-code-go/common"
	"log"
	"strconv"
	"strings"
	"time"
)

// Part1 processes the input strings and calculates a sum based on the problem's
// logic. It takes a slice of strings as input and returns an integer sum and an
// error. The function also logs the time taken to execute.
func Part1(input []string) (int, error) {
	start := time.Now() // Record the start time for execution duration tracking

	sum := 0 // Initialise sum to zero

	for _, line := range input {
		chars := strings.Split(line, "") // Split line into individual characters
		ints := []int{}                  // Initialise a slice to store integers

		// Convert each character to an integer
		for _, char := range chars {
			charInt, err := strconv.Atoi(char) // Convert character to integer

			if err != nil {
				continue // Skip characters that can't be converted to integers
			}

			ints = append(ints, charInt) // Add the converted integer to the slice
		}

		// Retrieve the first and last digits from the slice of integers
		firstDigit := ints[0]
		lastDigit := ints[len(ints)-1]

		// Calculate the contribution of this line to the sum
		sum += firstDigit*10 + lastDigit
	}

	log.Println("Part 1 took:", time.Since(start)) // Log time taken to execute

	return sum, nil // Return the calculated sum and nil error
}

func main() {
	input, err := common.ReadInputFile() // Read input from file

	if err != nil {
		log.Fatal(err) // Log and terminate if there's an error reading the file
	}

	// Split into lines
	inputData := strings.Split(strings.Trim(input, " "), "\n")
	// Remove empty strings from the input
	values := common.RemoveEmptyStrings(inputData)

	// Execute Part 1 of the problem
	part1, err := Part1(values)

	if err != nil {
		log.Fatal(err)
	}

	log.Println("Part 1:", part1) // Log the result of Part 1
}
