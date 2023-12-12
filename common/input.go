// Package common provides utility functions shared across the project.
package common

import (
	"bufio"
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

// ReadInputFileAs2DSlice reads contents of a file and returns them as a 2D
// slice of runes. Useful for taking input as a 2D grid with coordinates.
// Will remove empty lines.
func ReadInputFileAs2DSlice() ([][]rune, error) {
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
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines [][]rune
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		if line != "" { // Check if line is not empty
			lines = append(lines, []rune(line))
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return lines, nil
}
