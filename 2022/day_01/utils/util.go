package utils

import (
	"bufio"
	"os"
)

// Function to get the input from a file
func GetInput() ([]string, error) {
	// Default filenamne
	filename := "puzzle-input"

	// Check if filename is passed as a command line argument
	if len(os.Args) > 1 {
		filename = os.Args[1]
	}

	// Read lines from the file
	lines, err := ReadLinesFromFile(filename)
	if err != nil {
		panic(err)
	}

	return lines, nil
}

// Function to read lines from a file
func ReadLinesFromFile(filename string) ([]string, error) {
	// Open the file
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Create a new Scanner for the file
	scanner := bufio.NewScanner(file)

	// Initialise a slice to hold the lines
	var lines []string

	// Read through the file line by line
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	// Check for errors during scanning
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return lines, nil
}
