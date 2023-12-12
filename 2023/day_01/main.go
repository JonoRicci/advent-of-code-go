// Package main solves the Advent of Code 2023 Day 01 problem.
package main

import (
	"fmt"
	"jonoricci/advent-of-code-go/common"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
	"unicode"
)

// Part1 calculates the sum of two digit numbers from a slice of strings.
// Each number is formed by the first and last digit of each string.
func Part1(input []string) (int, error) {
	start := time.Now()
	sum := 0

	for _, line := range input {
		var firstDigit, lastDigit rune

		// Find first digit
		for _, char := range line {
			if unicode.IsDigit(char) {
				firstDigit = char
				break
			}
		}

		// Find the last digit
		for i := len(line) - 1; i >= 0; i-- {
			if unicode.IsDigit(rune(line[i])) {
				lastDigit = rune(line[i])
				break
			}
		}

		// Convert first and last digits to integers and add to sum
		firstDigitInt, _ := strconv.Atoi(string(firstDigit))
		lastDigitInt, _ := strconv.Atoi(string(lastDigit))
		// Multiply first digit by 10 to add numbers together
		// 1abc2 = 1 (10) + 2 (2) = 12.
		sum += firstDigitInt*10 + lastDigitInt
	}

	log.Println("[INFO] Part 1 took:", time.Since(start)) // Log time taken to execute
	return sum, nil
}

// Part2 calculates the sum of two digit numbers from a slice of strings.
// Each number is formed by the first and last digit, where digits can be
// integers or spelled-out words in a provided map.
func Part2(input []string) (int, error) {
	start := time.Now()

	numberMap := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
		"six":   6,
		"seven": 7,
		"eight": 8,
		"nine":  9,
	}
	sum := 0 // Initalise sum to zero
	for _, line := range input {
		firstDigit, lastDigit, err := searchLine(line, numberMap)
		if err != nil {
			log.Println("[ERROR]:", err)
			log.Fatal(err)
		}
		sum += firstDigit*10 + lastDigit
	}

	log.Println("[INFO] Part 2 took:", time.Since(start))
	return sum, nil
}

// searchLine extracts the first and last digits from a line.
// Returns the digits as integers or an error if no valid digit is found.
func searchLine(line string, numberMap map[string]int) (int, int, error) {
	// Create regex pattern from numberMap keys
	var pattern strings.Builder
	var err error

	pattern.WriteString("(")
	for key := range numberMap {
		pattern.WriteString(key)
		pattern.WriteString("|")
	}
	pattern.WriteString("\\d)")

	regex := regexp.MustCompile(pattern.String())

	// Find matches
	firstMatch := regex.FindStringSubmatch(line)
	var firstDigit, lastDigit int
	if len(firstMatch) > 0 {
		firstDigit, err = digitValue(firstMatch[0], numberMap)
		if err != nil {
			log.Println("[ERROR]:", err)
		}
	}

	// Iterate backward through the string to find the last digit
	for i := len(line) - 1; i >= 0; i-- {
		lastMatch := regex.FindStringSubmatch(line[i:])
		if len(lastMatch) > 0 {
			lastDigit, err = digitValue(lastMatch[0], numberMap)
			if err != nil {
				log.Println("[ERROR]:", err)
			}
			break
		}
	}

	return firstDigit, lastDigit, nil
}

// digitValue converts a string to it's numeric value.
// It uses a map for spelled-out numbers and standard conversion for digits.
func digitValue(match string, numberMap map[string]int) (int, error) {
	if digit, exists := numberMap[match]; exists {
		return digit, nil
	}
	digit, err := strconv.Atoi(match)
	if err != nil {
		return 0, fmt.Errorf("[ERROR]: Invalid digit: %s", match)
	}
	return digit, nil
}

func main() {
	// Set env var which dictates what input to use
	// Options are "", "PART_01", "PART_02"
	err := os.Setenv("ADVENT_OF_CODE_TEST", "")
	if err != nil {
		fmt.Println("Error setting environment variable:", err)
	}

	input, err := common.ReadInputFile()

	if err != nil {
		log.Println("[ERROR]:", err)
		log.Fatal(err)
	}

	// Split into lines
	inputData := strings.Split(strings.Trim(input, " "), "\n")
	// Remove empty strings from the input
	values := common.RemoveEmptyStrings(inputData)

	// Execute Part 1
	part1, err := Part1(values)
	if err != nil {
		log.Println("[ERROR]:", err)
		log.Fatal(err)
	}

	// Execute Part 2
	part2, err := Part2(values)
	if err != nil {
		log.Println("[ERROR]:", err)
		log.Fatal(err)
	}

	log.Println("[INFO] Part 1:", part1)
	log.Println("[INFO] Part 2:", part2)
}
