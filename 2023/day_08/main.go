// Package main solves the Advent of Code 2023 Day 08 problem.
package main

import (
	"fmt"
	"jonoricci/advent-of-code-go/common"
	"log"
	"sort"
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
	sum := 0

	directions := parseDirections(input[0])
	nodes := parseNodes(input[1:])

	sum, err := navigateGhostNodes(directions, nodes)
	if err != nil {
		return 0, err
	}

	logger.Infoln("Part 2 took:", time.Since(start))
	return sum, nil
}

// navigateGhostNodes takes a slice of directions and a map of nodes, starting
// from all nodes ending with 'A'. It navigates through the nodes based on the
// directions until all current nodes end with 'Z'. The directions are applied
// cyclically.
func navigateGhostNodes(directions []string, nodes map[string][2]string) (int, error) {
	// Initial setup to identify start nodes and prepare data structures for memoization, etc.
	// Implement an efficient traversal algorithm, possibly using BFS, DFS, or a modified version suitable for this problem.
	// Incorporate cycle detection to avoid infinite loops in paths.
	// Use memoization to store and reuse results of subproblems.
	// Prune paths that do not contribute to the solution.
	// Return the step count when the condition is met.

	// Identify start nodes (end in "A")
	var startNodes []string
	for node := range nodes {
		if strings.HasSuffix(node, "A") {
			startNodes = append(startNodes, node)
		}
	}
	logger.Debugln("Starting nodes:", startNodes)

	// My attempt at BFS (Breadth First Search)
	// Queue for BFS
	type nodeState struct {
		nodes []string
		steps int
	}
	queue := []nodeState{{nodes: startNodes, steps: 0}}

	if len(startNodes) == 0 {
		logger.Debugln("No starting nodes ending with 'A' found")
		return -1, fmt.Errorf("no starting nodes found")
	}

	// Track visited states
	visited := make(map[string]struct{})

	directionLength := len(directions)
	for len(queue) > 0 {
		currentState := queue[0] // Deqeue first element
		queue = queue[1:]        // Queue is empty

		// Check if all nodes end with "Z"
		if endWithZ(currentState.nodes) {
			logger.Debugln("Solution found at steps:", currentState.steps)
			return currentState.steps, nil
		}

		// Serialise current state for cycle detection
		serialised := serialisedState(currentState.nodes)
		if _, exists := visited[serialised]; exists {
			logger.Debugln("Skipping visited state:", serialised)
			continue // skip if state already visited
		}
		visited[serialised] = struct{}{}

		// Generate next states
		nextNodes := make([]string, len(currentState.nodes))
		direction := directions[currentState.steps%directionLength]
		logger.Debugln("Applying direction:", direction, "at step:", currentState.steps)

		for i, node := range currentState.nodes {
			if nodePair, exists := nodes[node]; exists {
				nextNodes[i] = nodePair[directionIndex(direction)]
				// logger.Debugln("Node:", node, "Next Node:", nextNodes[i])
			} else {
				return -1, fmt.Errorf("invalid node: %s", node)
			}
		}
		// logger.Debugln("Adding new state to queue:", nodeState{nodes: nextNodes, steps: currentState.steps + 1})
		queue = append(queue, nodeState{nodes: nextNodes, steps: currentState.steps + 1})
		// logger.Debugln("Processing state:", currentState, "Queue size:", len(queue))
	}

	return -1, fmt.Errorf("solution not found")
}

// endWithZ checks if all nodes in the provided slice end with 'Z'. It returns
// true if every node in the slice has a name ending with 'Z', otherwise false.
func endWithZ(nodes []string) bool {
	for _, node := range nodes {
		if !strings.HasSuffix(node, "Z") {
			return false
		}
	}
	return true
}

// serialisedState ...
func serialisedState(nodes []string) string {
	sort.Strings(nodes) // Ensure consistent ordering for serialisation
	return strings.Join(nodes, ",")
}

// directionIndex ...
func directionIndex(direction string) int {
	if direction == "R" {
		return 1
	}
	return 0
}
