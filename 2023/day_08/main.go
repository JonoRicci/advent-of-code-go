// Package main solves the Advent of Code 2023 Day 08 problem.
package main

import (
	"fmt"
	"jonoricci/advent-of-code-go/common"
	"log"
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

// Part1 ...
func Part1(input []string) (int, error) {
	start := time.Now()
	sum := 0

	directons := parseDirections(input[0])
	nodes := parseNodes(input[1:])

	sum, err := navigateNodes(directons, nodes)
	if err != nil {
		return 0, err
	}

	logger.Infoln("Part 1 took:", time.Since(start))
	return sum, nil
}

// parseDirections takes a string of characters and splits each character into
// a slice
func parseDirections(input string) []string {
	return strings.Split(input, "")
}

// parseNodes takes a slice of nodes in the format 'AAA = (BBB, CCC)' and splits
// them into three parts, a node, a left element and a right element which are
// populated into a map
func parseNodes(input []string) map[string][2]string {
	nodes := make(map[string][2]string)
	for _, line := range input {
		parts := strings.Split(line, " = ")
		node := parts[0]
		// Extract both elements and remove parentheses.
		elements := strings.Split(parts[1][1:len(parts[1])-1], ", ")
		nodes[node] = [2]string{elements[0], elements[1]}
	}
	return nodes
}

// navigateNodes will iterate continuously through the directions
// interacting with the map of nodes to follow through the puzzle input until it
// finds the "ZZZ" node.
func navigateNodes(directions []string, nodes map[string][2]string) (int, error) {
	current := "AAA"
	steps := 0
	directionLength := len(directions)

	for current != "ZZZ" {
		direction := directions[steps%directionLength] // modulo ensures valid index
		// Steps will exceed directionLength. When moduluo used in a loop it can
		// cycle over a fixed range
		logger.Debugln("Current:", current, "Steps:", steps, "Direction:", direction)
		steps++
		if node, exists := nodes[current]; exists {
			if direction == "R" {
				current = node[1]
				logger.Debugln("Element:", node[1])
			} else if direction == "L" {
				current = node[0]
				logger.Debugln("Element:", node[0])
			}
		} else {
			return -1, fmt.Errorf("invlaid node: %s", current)
		}
	}
	return steps, nil
}

// Part2 ...
func Part2(input []string) (int, error) {
	start := time.Now()
	sum := 0

	logger.Infoln("Part 2 took:", time.Since(start))
	return sum, nil
}
