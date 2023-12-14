// Package main solves the Advent of Code 2023 Day 03 problem.
package main

import (
	"jonoricci/advent-of-code-go/common"
	"log"
	"strconv"
	"time"
	"unicode"

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
	input, err := common.ReadInputFileAs2DSlice(cfg)

	if err != nil {
		logger.Fatalln(err)
	}

	// Make copy of the input so each function can modify it's input independently
	inputPart1 := copy2DSlice(input)
	inputPart2 := copy2DSlice(input)

	// Execute Part 1
	part1, err := Part1(inputPart1)
	if err != nil {
		logger.Fatalln(err)
	}

	// Execute Part 2
	part2, err := Part2(inputPart2)
	if err != nil {
		logger.Fatalln(err)
	}

	logger.Infoln("Part 1:", part1)
	logger.Infoln("Part 2:", part2)
}

// copy2DSlice makes a deep copy of the input so that each function can modify
// it's input independently.
func copy2DSlice(original [][]rune) [][]rune {
	copiedSlice := make([][]rune, len(original))
	for i := range original {
		copiedSlice[i] = make([]rune, len(original[i]))
		copy(copiedSlice[i], original[i])
	}
	return copiedSlice
}

// Part1 calculates the sum of all the numbers adjacent to a symbol in a given
// 2D slice.
func Part1(input [][]rune) (int, error) {
	start := time.Now()
	sum := 0

	for i, line := range input {
		// log.Println("[DEBUG]:", string(line))
		for j, char := range line {
			// log.Println("[DEBUG]:", i, j, string(char))
			if isNumber(char) && isAdjacentToSymbol(input, i, j) {
				fullNumStr := getFullNumber(input, i, j)
				num, err := strconv.Atoi(fullNumStr)
				if err != nil {
					logger.Fatalln(err)
				}
				// log.Println("[DEBUG]:", num, sum)
				sum += num

				// Mark number as processed to avoid double counting
				markNumberAsProcessed(input, i, j)
			}
		}
	}

	logger.Infoln("Part 1 took:", time.Since(start))
	return sum, nil
}

// isNumber simply checks if a given rune is an integer or not
func isNumber(char rune) bool {
	if unicode.IsDigit(char) {
		return true
	} else {
		return false
	}
}

// isAdjacentToSymbol checks if coordinates in a given 2D slice are ajacent to
// a symbol.
func isAdjacentToSymbol(input [][]rune, y int, x int) bool {
	// Define relative positions around given coordinate
	directions := []struct{ dy, dx int }{
		{-1, -1}, {-1, 0}, {-1, 1}, // Above
		{0, -1}, {0, 1}, // Sides
		{1, -1}, {1, 0}, {1, 1}, // Below
	}

	// Iterate over all adjacent positions
	for _, d := range directions {
		newY, newX := y+d.dy, x+d.dx
		if isValidPosition(newY, newX, len(input), len(input[0])) {
			adjacentChar := input[newY][newX]
			if !unicode.IsDigit(adjacentChar) && adjacentChar != '.' {
				return true
			}
		}
	}
	return false
}

// isValidPosition checks if given coordinates are within the boundaries of the
// grid.
func isValidPosition(y, x, maxY, maxX int) bool {
	if y >= 0 && y < maxY && x >= 0 && x < maxX {
		return true
	} else {
		return false
	}
}

// getFullNumber searches the left and right of an individual digit to find the
// full number, assuming '.' is a number separator.
func getFullNumber(input [][]rune, y, x int) string {
	numStr := string(input[y][x])

	// Scan left
	for i := x - 1; i >= 0 && unicode.IsDigit(input[y][i]); i-- {
		numStr = string(input[y][i]) + numStr
	}

	// Scan right
	for i := x + 1; i < len(input[y]) && unicode.IsDigit(input[y][i]); i++ {
		numStr += string(input[y][i])
	}

	return numStr
}

func markNumberAsProcessed(input [][]rune, y, x int) {
	// Set the current position to a non-digit character
	input[y][x] = 'x'

	// Scan and mark to the left
	for i := x - 1; i >= 0 && unicode.IsDigit(input[y][i]); i-- {
		input[y][i] = 'x'
	}

	// Scan and mark to the right
	for i := x + 1; i < len(input[y]) && unicode.IsDigit(input[y][i]); i++ {
		input[y][i] = 'x'
	}
}

// Part2 calculates the sum of gear ratios (two part numbers adjacent to a *
// symbol and multiplied together).
func Part2(input [][]rune) (int, error) {
	start := time.Now()
	sum := 0

	for y, line := range input {
		// log.Println("[DEBUG]:", line)
		for x, char := range line {
			// log.Println("[DEBUG]:", string(char))
			if char == '*' {
				nums, valid := getAdjacentNumbers(input, y, x)
				// log.Println("[DEBUG]:", string(char), nums, valid)
				if valid && len(nums) == 2 {
					num1, err1 := strconv.Atoi(nums[0])
					num2, err2 := strconv.Atoi(nums[1])
					if err1 != nil || err2 != nil {
						logger.Fatalln(err1, err2)
					}
					gearRatio := num1 * num2
					sum += gearRatio
				}
			}
		}
	}

	logger.Infoln("Part 2 took:", time.Since(start))
	return sum, nil
}

// getAdjacentNumbers finds the numbers adjacent to a given position.
// It returns numbers as strings and a boolean indicating if the position is
// valid.
func getAdjacentNumbers(input [][]rune, y, x int) ([]string, bool) {
	// Define relative positions around given coordinate
	directions := []struct{ dy, dx int }{
		{-1, -1}, {-1, 0}, {-1, 1}, // Above
		{0, -1}, {0, 1}, // Sides
		{1, -1}, {1, 0}, {1, 1}, // Below
	}

	var nums []string
	for _, d := range directions {
		newY, newX := y+d.dy, x+d.dx
		if isValidPosition(newY, newX, len(input), len(input[0])) && isNumber(input[newY][newX]) {
			numStr := getFullNumber(input, newY, newX)
			nums = append(nums, numStr)
			markNumberAsProcessed(input, newY, newX)
		}
	}
	return nums, true
}
