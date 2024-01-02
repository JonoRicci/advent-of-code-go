// Package main solves the Advent of Code 2015 Day 01 problem.
package main

import (
	"fmt"
	"jonoricci/advent-of-code-go/common"
	"log"
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
	logger.Infoln("Part 1:", part1)

	// Execute Part 2
	part2, err := Part2(values)
	if err != nil {
		logger.Fatalln(err)
	}
	logger.Infoln("Part 2:", part2)
}

// Part1 calculates the final floor Santa will arrive on.
func Part1(input []string) (int, error) {
	start := time.Now()
	floor := 0

	// Only one line of input
	for _, line := range input {
		// Iterate over each character in the line
		for _, char := range line {
			switch char {
			case '(': // Go up a floor
				floor++
			case ')': // Go down a floor
				floor--
			}
		}
	}

	logger.Infoln("Part 1 took:", time.Since(start))
	return floor, nil
}

// Part2 finds position of the first character that causes Santa to enter the
// basement
func Part2(input []string) (int, error) {
	start := time.Now()
	floor := 0
	position := 0

	// Assume only one line of input
	for _, line := range input {
		for i, char := range line {
			position = i + 1 // First character is position 1 not 0
			switch char {
			case '(':
				floor++
			case ')':
				floor--
			}

			// Check if Santa has entered the basement
			if floor == -1 {
				logger.Infoln("Part 2 took:", time.Since(start))
				return position, nil
			}
		}
	}
	logger.Infoln("Part 2 took:", time.Since(start))
	return -1, fmt.Errorf("santa does not enter the basement")
}
