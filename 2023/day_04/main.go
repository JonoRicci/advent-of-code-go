// Package main solves the Advent of Code 2023 Day 04 problem.
package main

import (
	"fmt"
	"jonoricci/advent-of-code-go/common"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	// Set env var which dictates what input to use
	// Options are "", "PART_01", "PART_02"
	err := os.Setenv("ADVENT_OF_CODE_TEST", "")
	if err != nil {
		fmt.Println("Error setting environment variable:", err)
	}

	input, err := common.ReadInputFile()

	if err != nil {
		log.Fatalln("[ERROR]:", err)
	}

	// Split into lines
	inputData := strings.Split(strings.Trim(input, " "), "\n")
	// Remove empty strings from the input
	values := common.RemoveEmptyStrings(inputData)

	// Execute Part 1
	part1, err := Part1(values)
	if err != nil {
		log.Fatalln("[ERROR]:", err)
	}

	// Execute Part 2
	part2, err := Part2(values)
	if err != nil {
		log.Fatalln("[ERROR]:", err)
	}

	log.Println("[INFO] Part 1:", part1)
	log.Println("[INFO] Part 2:", part2)
}

// Part1 ...
func Part1(input []string) (int, error) {
	start := time.Now()
	sum := 0

	for _, line := range input {
		// Split up each line to get two slices
		colonIndex := strings.Index(line, ":")
		line = line[colonIndex+2:]
		parts := strings.Split(line, "|")

		winningNumsStr := strings.Fields(strings.TrimSpace(parts[0]))
		yourNumsStr := strings.Fields(strings.TrimSpace(parts[1]))

		winningNums, err := convertToIntSlice(winningNumsStr)
		if err != nil {
			log.Fatalln("[ERROR]:", err)
		}

		yourNums, err := convertToIntSlice(yourNumsStr)
		if err != nil {
			log.Fatalln("[ERROR]:", err)
		}

		// Create a map for winning numbers lookup
		winningNumsMap := make(map[int]bool)
		for _, num := range winningNums {
			winningNumsMap[num] = true
		}

		// Check for matches and calculate score
		score := 0
		for _, num := range yourNums {
			if winningNumsMap[num] {
				if score == 0 {
					score = 1
				} else {
					score *= 2
				}
			}
		}

		// Add total to score
		sum += score
	}

	log.Println("[INFO] Part 1 took:", time.Since(start))
	return sum, nil
}

// convertToIntSlice converts a slice of strings to a slice of ints
func convertToIntSlice(strSlice []string) ([]int, error) {
	var intSlice []int
	for _, str := range strSlice {
		num, err := strconv.Atoi(str)
		if err != nil {
			log.Fatalln("[ERROR]:", err)
		}
		intSlice = append(intSlice, num)
	}
	return intSlice, nil
}

// Part2 ...
func Part2(input []string) (int, error) {
	start := time.Now()

	// Store the number of copies for each card
	cardCopies := make([]int, len(input))
	for i := range cardCopies {
		cardCopies[i] = 1 // Each card starts with 1 copy (itself)
	}

	totalCards := 0

	for i := 0; i < len(input); i++ {
		// Split up each line to get two slices
		line := input[i]
		colonIndex := strings.Index(line, ":")
		line = line[colonIndex+2:]
		parts := strings.Split(line, "|")

		winningNumsStr := strings.Fields(strings.TrimSpace(parts[0]))
		yourNumsStr := strings.Fields(strings.TrimSpace(parts[1]))

		winningNums, err := convertToIntSlice(winningNumsStr)
		if err != nil {
			log.Fatalln("[ERROR]:", err)
		}

		yourNums, err := convertToIntSlice(yourNumsStr)
		if err != nil {
			log.Fatalln("[ERROR]:", err)
		}

		winningNumsMap := make(map[int]bool)
		for _, num := range winningNums {
			winningNumsMap[num] = true
		}

		// Count matches instead of calculating score
		matchCount := 0
		for _, num := range yourNums {
			if winningNumsMap[num] {
				matchCount++
			}
		}

		// For each match, add a copy to subsequent cards
		for j := 1; j <= matchCount; j++ {
			if (i + j) < len(cardCopies) {
				cardCopies[i+j] += cardCopies[i]
			}
		}
	}

	// Calculate total number of cards including copies
	for _, copies := range cardCopies {
		totalCards += copies
	}

	log.Println("[INFO] Part 2 took:", time.Since(start))
	return totalCards, nil
}
