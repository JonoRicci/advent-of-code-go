// // Package main solves the Advent of Code 2023 Day 09 problem.
package main

import (
	"jonoricci/advent-of-code-go/common"
	"log"
	"strings"
	"testing"
)

// MockInput returns a slice of strings which is the puzzle input.
func MockInput() []string {
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

	return values
}

// TestPart1 ensures function produces the correct result.
// This works for either test or real puzzle input.
func TestPart1(t *testing.T) {
	expectedValues := []int{114, 1938800261}
	result, err := Part1(MockInput())
	if err != nil {
		t.Fatalf("Part1 returned an error: %v", err)
	}

	found := false
	for _, expected := range expectedValues {
		if result == expected {
			found = true
			break
		}
	}

	if !found {
		t.Errorf("Expected Part1 to return %d, got %d", expectedValues, result)
	}
}

// TestPart2 ensures function produces the correct result.
// This works for either test or real puzzle input.
func TestPart2(t *testing.T) {
	expectedValues := []int{2, 1112}
	result, err := Part2(MockInput())
	if err != nil {
		t.Fatalf("Part2 returned an error: %v", err)
	}

	found := false
	for _, expected := range expectedValues {
		if result == expected {
			found = true
			break
		}
	}

	if !found {
		t.Errorf("Expected Part2 to return %d, got %d", expectedValues, result)
	}
}
