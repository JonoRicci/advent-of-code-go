// Package main solves the Advent of Code 2023 Day 01 problem.
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

	// Check input length
	if len(input) != 2 {
		return 0, fmt.Errorf("input should only have two lines")
	}
	logger.Debugln("First line:", input[0])
	logger.Debugln("Second line:", input[1])

	// Remove "Time:" from the first line and "Distance:" from the second line
	timeLine := strings.TrimPrefix(input[0], "Time:")
	distanceLine := strings.TrimPrefix(input[1], "Distance:")

	// Parse times
	timeStrings := strings.Fields(timeLine)
	times := make([]int, len(timeStrings))
	for i, t := range timeStrings {
		logger.Debugln("Position:", i, "Time:", t)
		time, err := strconv.Atoi(t)
		if err != nil {
			return 0, err
		}
		times[i] = time
	}

	// Parse distances
	distanceStrings := strings.Fields(distanceLine)
	distances := make([]int, len(distanceStrings))
	for i, d := range distanceStrings {
		logger.Debugln("Position:", i, "Distance:", d)
		distance, err := strconv.Atoi(d)
		if err != nil {
			return 0, err
		}
		distances[i] = distance
	}

	// Calculate winning strategy
	totalWays := 1
	for i := 0; i < len(times); i++ {
		ways := 0
		for j := 0; j < times[i]; j++ {
			if ((times[i] - j) * j) > distances[i] {
				ways++
			}
		}
		totalWays *= ways
		logger.Debugln("Ways:", ways)
		logger.Debugln("Total ways:", totalWays)
	}

	logger.Infoln("Part 1 took:", time.Since(start))
	return totalWays, nil
}

// Part2 ...
func Part2(input []string) (int, error) {
	start := time.Now()

	// Check input length
	if len(input) != 2 {
		return 0, fmt.Errorf("input should only have two lines")
	}

	// Swap spaces for nothing
	timeLine := strings.ReplaceAll(input[0], " ", "")
	distanceLine := strings.ReplaceAll(input[1], " ", "")

	// Remove "Time:" and "Distance:"
	timeLine = strings.TrimPrefix(timeLine, "Time:")
	distanceLine = strings.TrimPrefix(distanceLine, "Distance:")
	logger.Debugln("Real time:", timeLine)
	logger.Debugln("Real distance:", distanceLine)

	// Convert time and distance to integers
	timeInt, err := strconv.Atoi(timeLine)
	if err != nil {
		return 0, err
	}
	distanceInt, err := strconv.Atoi(distanceLine)
	if err != nil {
		return 0, err
	}

	// Calculate the number of ways to win
	waysToWin := 0
	for i := 0; i < timeInt; i++ {
		if ((timeInt - i) * i) > distanceInt {
			waysToWin++
		}
	}

	logger.Infoln("Part 2 took:", time.Since(start))
	return waysToWin, nil
}
