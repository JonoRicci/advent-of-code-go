// Package main solves the Advent of Code 2023 Day 04 problem.
package main

import (
	"jonoricci/advent-of-code-go/common"
	"log"
	"strconv"
	"strings"
	"time"

	"go.uber.org/zap"
)

// global variable for logging
var logger *zap.SugaredLogger

func main() {
	// Load config file
	cfg, err := common.ReadConfig()
	if err != nil {
		log.Fatalf("Error reading config: %v", err)
	}

	// Initalise logging
	logger, err = common.InitialiseLogger(cfg)
	if err != nil {
		log.Fatalf("Error initialising logger: %v", err)
	}
	defer logger.Sync() // Flush any buffered log entries

	// Read puzzle input
	input, err := common.ReadInputFile(cfg)

	if err != nil {
		logger.Fatalln(err)
	}

	// Split into lines
	inputData := strings.Split(strings.Trim(input, " "), "\n")
	// Remove empty strings from the input
	values := common.RemoveEmptyStrings(inputData)

	// Execute Part 1
	part1, err := Part1(values)
	if err != nil {
		logger.Fatalln(err)
	}

	// Execute Part 2
	part2, err := Part2(values)
	if err != nil {
		logger.Fatalln(err)
	}

	logger.Infoln("Part 1:", part1)
	logger.Infoln("Part 2:", part2)
}

// Part1 processes a list of scratchcards, calculates the score for each card
// based on the number of matching numbers with the winning numbers, and returns
// the total score of all cards.
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
			logger.Fatalln(err)
		}

		yourNums, err := convertToIntSlice(yourNumsStr)
		if err != nil {
			logger.Fatalln(err)
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

	logger.Infoln("Part 1 took:", time.Since(start))
	return sum, nil
}

// convertToIntSlice converts a slice of strings to a slice of ints
func convertToIntSlice(strSlice []string) ([]int, error) {
	var intSlice []int
	for _, str := range strSlice {
		num, err := strconv.Atoi(str)
		if err != nil {
			logger.Fatalln(err)
		}
		intSlice = append(intSlice, num)
	}
	return intSlice, nil
}

// Part2 processes the scratchcards according to the new rules where each
// matching number wins additional scratchcards. It returns the total number of
// scratchcards, including both the original and the won copies.
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
			logger.Fatalln(err)
		}

		yourNums, err := convertToIntSlice(yourNumsStr)
		if err != nil {
			logger.Fatalln(err)
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

	logger.Infoln("Part 2 took:", time.Since(start))
	return totalCards, nil
}
