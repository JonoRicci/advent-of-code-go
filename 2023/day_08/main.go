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
		logger.Warnln(err) // Switched to warn from fatal to pass part 02 test input
	}
	logger.Infoln("Part 1:", part1)

	// Execute Part 2
	part2, err := Part2(values)
	if err != nil {
		logger.Fatalln(err)
	}
	logger.Infoln("Part 2:", part2)
}

// Part1 navigates through the puzzle input to count the steps from "AAA" to
// "ZZZ".
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

// Part2 navigates through the puzzle input to count the steps using the Ghost
// method of navigation, which is to start simultaneously on all nodes ending
// in A and navigate through all of them simultaneously where the result is all
// nodes are on a step where each node ends in Z.
func Part2(input []string) (int, error) {
	start := time.Now()

	directions := parseDirections(input[0])
	nodes := parseNodes(input[1:])

	// Find individual path lengths
	var pathLengths []int
	for node := range nodes {
		if strings.HasSuffix(node, "A") {
			length, err := navigateIndividualPath(node, "Z", directions, nodes)
			if err != nil {
				return 0, err
			}
			pathLengths = append(pathLengths, length)
		}
	}

	// Calculate LCM of path lengths
	lcm, err := calculateLCM(pathLengths)
	if err != nil {
		return 0, err
	}

	logger.Infoln("Part 2 took:", time.Since(start))
	return lcm, nil
}

// navigateIndividualPath navigates from a given start node to an end node (that
// ends with 'Z').
func navigateIndividualPath(startNode, endSuffix string, directions []string, nodes map[string][2]string) (int, error) {
	steps := 0
	directionLength := len(directions)
	currentNode := startNode

	for !strings.HasSuffix(currentNode, endSuffix) {
		direction := directions[steps%directionLength]
		nextNode := nodes[currentNode][directionIndex(direction)]
		if nextNode == "" {
			return -1, fmt.Errorf("invalid node: %s", currentNode)
		}
		currentNode = nextNode
		steps++
	}
	return steps, nil
}

// directionIndex converts a direction character ('L' or 'R') into an index (0 or 1).
func directionIndex(direction string) int {
	if direction == "R" {
		return 1
	}
	return 0
}

// calculateLCM calculates the least common multiple of a slice of integers.
func calculateLCM(numbers []int) (int, error) {
	if len(numbers) == 0 {
		return 0, fmt.Errorf("empty slice provided")
	}
	lcm := numbers[0]
	for _, num := range numbers[1:] {
		lcm = lcmTwoNumbers(lcm, num)
	}
	return lcm, nil
}

// lcmTwoNumbers calculates the LCM of two numbers.
func lcmTwoNumbers(a, b int) int {
	return a * b / gcdTwoNumbers(a, b)
}

// gcdTwoNumbers calculates the greatest common divisor (GCD) of two numbers.
func gcdTwoNumbers(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}
