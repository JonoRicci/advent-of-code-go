// Package main solves the Advent of Code 2023 Day 02 problem.
package main

import (
	"jonoricci/advent-of-code-go/common"
	"log"
	"strconv"
	"strings"
	"time"
)

// // Part1 takes an array of strings representing the game input and returns
// the sum of the IDs of the games that are possible within the given cube
// constraints.
func Part1(input []string) (int, error) {
	start := time.Now()
	sum := 0

	for _, line := range input {
		// Split up input line into gameID and subsets
		parts := strings.Split(line, ": ")
		gameID, _ := strconv.Atoi(strings.Split(parts[0], " ")[1])
		subsets := strings.Split(parts[1], "; ")

		if checkGamePossible(subsets) {
			sum += gameID
		}
	}

	log.Println("[INFO] Part 1 took:", time.Since(start))
	return sum, nil
}

// checkGamePossible takes an array of subsets of cubes and returns true if the
// game is possible within the given cube constraints.
func checkGamePossible(subsets []string) bool {
	const maxRed, maxGreen, maxBlue = 12, 13, 14 // Given by problem input

	for _, subset := range subsets {
		red, green, blue := countCubes(subset)
		// Check if any color exceeds its maximum allowed cubes
		if red > maxRed || green > maxGreen || blue > maxBlue {
			return false // Game is not possible if any subset exceeds the constraints
		}
	}
	return true // Return true if all subsets are within the constraints
}

// countCubes takes a string representing a subset of cubes and returns the
// count of red, green, and blue cubes in that subset.
func countCubes(subset string) (int, int, int) {
	red, green, blue := 0, 0, 0 // Initalise at zero
	cubes := strings.Split(subset, ", ")

	for _, cube := range cubes {
		parts := strings.Split(cube, " ")
		count, _ := strconv.Atoi(parts[0]) // Parse the count of cubes
		colour := parts[1]                 // Get the colour of cubes

		// Increment count of cube colour
		switch colour {
		case "red":
			red += count
		case "green":
			green += count
		case "blue":
			blue += count
		}
	}
	return red, green, blue
}

// Part2 calculates the sum of the powers of the minimum sets of cubes needed
// for each game.
func Part2(input []string) (int, error) {
	start := time.Now()
	sum := 0

	for _, line := range input {
		// Split up input line into gameID and subsets
		parts := strings.Split(line, ": ")
		subsets := strings.Split(parts[1], "; ")
		// Find the minimum cubes required
		minRed, minGreen, minBlue := findMinimumSet(subsets)
		sum += minRed * minGreen * minBlue // Power of the set
	}

	log.Println("[INFO] Part 2 took:", time.Since(start))
	return sum, nil
}

// findMinimumSet returns the minimum number of coloured cubes needed for a game
func findMinimumSet(subsets []string) (int, int, int) {
	minRed, minGreen, minBlue := 0, 0, 0

	for _, subset := range subsets {
		red, green, blue := countCubes(subset)
		if red > minRed {
			minRed = red
		}
		if green > minGreen {
			minGreen = green
		}
		if blue > minBlue {
			minBlue = blue
		}
	}
	return minRed, minGreen, minBlue
}

func main() {
	input, err := common.ReadInputFile()

	if err != nil {
		log.Println("[ERROR]:", err)
		log.Fatal(err)
	}

	// Split into lines
	inputData := strings.Split(strings.Trim(input, " "), "\n")
	// Remove empty strings from the input
	values := common.RemoveEmptyStrings(inputData)

	// Execute Part 1
	part1, err := Part1(values)
	if err != nil {
		log.Println("[ERROR]:", err)
		log.Fatal(err)
	}

	// Execute Part 2
	part2, err := Part2(values)
	if err != nil {
		log.Println("[ERROR]:", err)
		log.Fatal(err)
	}

	log.Println("[INFO] Part 1:", part1)
	log.Println("[INFO] Part 2:", part2)
}
