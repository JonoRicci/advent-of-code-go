// Package common provides utility functions shared across the project.
package common

import (
	"log/slog"
	"strconv"
)

// SumStrings takes a slice of strings, converts each to an integer, and returns
// the sum. If a string cannot be converted to an integer, the function logs a
// fatal error and stops execution.
func SumStrings(s []string) int {
	sum := 0 // Initialise sum to zero

	for _, str := range s {
		if str == "" {
			continue // Skip empty strings
		}

		// Attempt to convert string to an int
		num, err := strconv.Atoi(str)

		if err != nil {
			slog.Error("SumStrings Failed:", err)
		}

		// Add converted int to total sum
		sum += num
	}

	return sum
}

// SumInts takes a slice of integers and returns their sum.
func SumInts(i []int) int {
	sum := 0 // Initialise sum to zero

	for _, num := range i {
		sum += num // Add each number to the sum
	}

	return sum
}

// RemoveEmptyStrings takes a slice of strings and returns a new slice with all
// empty strings removed.
func RemoveEmptyStrings(s []string) []string {
	var r []string // Initialise an empty slice

	for _, str := range s {
		if str != "" {
			r = append(r, str) // Add non-empty strings to slice
		}
	}

	return r
}
