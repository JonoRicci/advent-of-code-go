// Package common provides utility functions shared across the project.
package common

import (
	"os"
)

// ReadInputFile reads contents of a file and returns them as a string.
// The function reads from `input.txt` by default unless the env var
// "ADVENT_OF_CODE_TEST" has been set to "TRUE".
func ReadInputFile() (string, error) {
	filePath := "input.txt"

	// Check for env var and determinte what input to use.
	envValue := os.Getenv("ADVENT_OF_CODE_TEST")
	if envValue == "TRUE" {
		filePath = "test.txt"
	} else if envValue == "PART_01" {
		filePath = "part01_test.txt"
	} else if envValue == "PART_02" {
		filePath = "part02_test.txt"
	}

	// Read the file specified by filePath.

	data, err := os.ReadFile(filePath)

	// Go error handling.
	if err != nil {
		return "", err
	}

	// Convert data from byte slice to a string and return.
	return string(data), nil
}
