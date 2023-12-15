// // Package main solves the Advent of Code 2023 Day 07 problem.
package main

import (
	"fmt"
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
		got := cardStrength(c.card)
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
			handType, _ := evaluateHand(tc.hand)
			if handType != tc.expected {
				t.Errorf("Expected handType for %s to be %d, got %d", tc.hand, tc.expected, handType)
			}
		})
	}
}
