// Package main solves the Advent of Code 2023 Day 09 problem.
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
	logger.Infoln("Part 1:", part1)

	// Execute Part 2
	part2, err := Part2(values)
	if err != nil {
		logger.Fatalln(err)
	}
	logger.Infoln("Part 2:", part2)
}

// Part1 takes a sequence of consecutively increasing ints and extrapolates the
// next value.
func Part1(input []string) (int, error) {
	start := time.Now()
	sum := 0

	sequences, err := parseInputToInts(input)
	if err != nil {
		return 0, err
	}

	for _, seq := range sequences {
		sum += extrapolateNextValue(seq)
	}

	logger.Infoln("Part 1 took:", time.Since(start))
	return sum, nil
}

// parseInputToInts takes line input into slices of integers.
func parseInputToInts(input []string) ([][]int, error) {
	var sequences [][]int
	for _, line := range input {
		var seq []int
		for _, numStr := range strings.Fields(line) {
			num, err := strconv.Atoi(numStr)
			if err != nil {
				return nil, err
			}
			seq = append(seq, num)
		}
		sequences = append(sequences, seq)
	}
	return sequences, nil
}

// extrapolateNextValue finds the next value of the sequence by adding the
// lower sequences together.
func extrapolateNextValue(seq []int) int {
	sequences := generateAllSequences(seq)
	logger.Debugln("Sequences:", sequences)
	for i := len(sequences) - 2; i >= 0; i-- {
		lastNum := sequences[i][len(sequences[i])-1]
		diff := sequences[i+1][len(sequences[i+1])-1]
		nextVal := lastNum + diff
		sequences[i] = append(sequences[i], nextVal)
	}
	logger.Debugln("Final next extrapolated value:", sequences[0][len(sequences[0])-1])
	return sequences[0][len(sequences[0])-1]
}

// generateAllSequences generates sequences down to the zero sequence.
func generateAllSequences(seq []int) [][]int {
	var sequences [][]int
	sequences = append(sequences, seq)

	for {
		lastSeq := sequences[len(sequences)-1]
		diff := calculateDifferences(lastSeq)
		sequences = append(sequences, diff)
		if allZeroes(diff) {
			break
		}
	}

	return sequences
}

// calculateDifferences calculates the differences between consecutive numbers
// in a sequence.
func calculateDifferences(seq []int) []int {
	var diffs []int
	for i := 1; i < len(seq); i++ {
		diffs = append(diffs, seq[i]-seq[i-1])
	}
	return diffs
}

// allZeroes takes a slice of ints and checks if they are all zeroes
func allZeroes(seq []int) bool {
	for _, num := range seq {
		if num != 0 {
			return false
		}
	}
	return true
}

// Part2 takes a sequence of consecutively increasing ints and extrapolates the
// previous value.
func Part2(input []string) (int, error) {
	start := time.Now()
	sum := 0

	sequences, err := parseInputToInts(input)
	if err != nil {
		return 0, err
	}

	for _, seq := range sequences {
		sum += extrapolatePreviousValue(seq)
	}

	logger.Infoln("Part 2 took:", time.Since(start))
	return sum, nil
}

// extrapolatePreviousValue ...
func extrapolatePreviousValue(seq []int) int {
	sequences := generateAllSequences(seq)
	logger.Debugln("Sequences:", sequences)

	// Add a zero at the beginning of the zero sequence
	zeroSeq := append([]int{0}, sequences[len(sequences)-1]...)
	sequences[len(sequences)-1] = zeroSeq

	for i := len(sequences) - 2; i >= 0; i-- {
		firstNum := sequences[i][0]
		diff := sequences[i+1][0]
		prevNum := firstNum - diff
		sequences[i] = append([]int{prevNum}, sequences[i]...)
	}

	logger.Debugln("Final previous extrapolated value:", sequences[0][0])
	return sequences[0][0]
}
