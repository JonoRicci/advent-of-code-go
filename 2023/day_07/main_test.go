// // Package main solves the Advent of Code 2023 Day 07 problem.
package main

import (
	"fmt"
	"jonoricci/advent-of-code-go/common"
	"log"
	"strings"
	"testing"
)

// TestCardStrength ensures we get the correct strength rank back from a card.
func TestCardStrength(t *testing.T) {
	cases := []struct {
		card rune
		want int
	}{
		{'A', 13},
		{'K', 12},
		{'Q', 11},
		{'J', 10},
		{'T', 9},
		{'9', 8},
		{'8', 7},
		{'7', 6},
		{'6', 5},
		{'5', 4},
		{'4', 3},
		{'3', 2},
		{'2', 1},
	}

	for _, c := range cases {
		got := cardStrength(c.card, false)
		if got != c.want {
			t.Errorf("cardStrength(%q) == %d, want %d", c.card, got, c.want)
		}
	}
}

// TestEvaluateHand ensures we get the correct hand type of a hand
func TestEvaluateHand(t *testing.T) {
	type testCase struct {
		hand     string
		expected int
	}

	testCases := []testCase{
		{"AAAAA", 7}, // Five of a kind
		{"AA8AA", 6}, // Four of a kind
		{"23332", 5}, // Full house
		{"TTT98", 4}, // Three of a kind
		{"23432", 3}, // Two pair
		{"A23A4", 2}, // One pair
		{"23456", 1}, // High card
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("Hand: %s", tc.hand), func(t *testing.T) {
			handType, _ := evaluateHand(tc.hand, false)
			if handType != tc.expected {
				t.Errorf("Expected handType for %s to be %d, got %d", tc.hand, tc.expected, handType)
			}
		})
	}
}

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
	expectedValues := []int{6440, 249390788}
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
	expectedValues := []int{5905, 248750248}
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
