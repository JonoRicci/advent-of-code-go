// Package main solves the Advent of Code 20XX Day 0X problem.
package main

import (
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

// Part1 ...
func Part1(input []string) (int, error) {
	start := time.Now()
	sum := 0

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
