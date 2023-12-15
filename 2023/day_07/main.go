// Package main solves the Advent of Code 2023 Day 07 problem.
package main

import (
	"fmt"
	"jonoricci/advent-of-code-go/common"
	"log"
	"strconv"
	"strings"
	"time"

	"go.uber.org/zap"
)

// global variable for logging
var logger *zap.SugaredLogger

// global variable for cards
var cardLabels = []string{
	"A",
	"K",
	"Q",
	"J",
	"T",
	"9",
	"8",
	"7",
	"6",
	"5",
	"4",
	"3",
	"2",
}

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

// Part1 ...
func Part1(input []string) (int, error) {
	start := time.Now()
	sum := 0

	// Parse input, get hands and bids
	for _, line := range input {
		logger.Debugln("Line:", line)
		splitLine := strings.Split(line, " ")
		hand := splitLine[0]
		bid, err := strconv.Atoi(splitLine[1])
		if err != nil {
			return 0, fmt.Errorf("failed to convert to int %v", err)
		}
		logger.Debugln("Hand:", hand)
		logger.Debugln("Bid:", bid)
	}

	// Work out hand type
	// Work out type order
	// Work out hand ranks

	logger.Infoln("Part 1 took:", time.Since(start))
	return sum, nil
}

// Part2 ...
func Part2(input []string) (int, error) {
	start := time.Now()
	sum := 0

	logger.Infoln("Part 2 took:", time.Since(start))
	return sum, nil
}
