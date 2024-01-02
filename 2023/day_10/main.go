// Package main solves the Advent of Code 2023 Day 10 problem.
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

// Part1 finds the furthest distance in the loop from the start.
func Part1(input []string) (int, error) {
	start := time.Now()

	// Convert input into a grid of runes
	grid := make([][]rune, len(input))
	for i, line := range input {
		// logger.Debugln(line)
		grid[i] = []rune(line)
	}

	// Debug: Print the grid
	for _, line := range grid {
		logger.Debugln(string(line))
	}

	// Find starting position "S"
	startPos := findStartPos(grid)
	if startPos == nil {
		return 0, fmt.Errorf("start position not found")
	}

	// Debug: Print the start position
	logger.Debugln("Start Position:", startPos)

	// Walk through the loop to calculate distance
	distances := make(map[Pos]int)
	walkLoop(grid, *startPos, distances)

	// Find maximum distance
	maxDistance := findMaxDistance(distances)

	logger.Infoln("Part 1 took:", time.Since(start))
	return maxDistance, nil
}

// Pos is a position on the grid
type Pos struct {
	x, y int
}

// findStartPos locates starting position "S"
func findStartPos(grid [][]rune) *Pos {
	for y, row := range grid {
		for x, cell := range row {
			if cell == 'S' {
				return &Pos{x, y}
			}
		}
	}
	return nil
}

// walkLoop performs DFS to walk through the loop and updates distances.
func walkLoop(grid [][]rune, startPos Pos, distances map[Pos]int) {
	var dfs func(p Pos, dist int)
	dfs = func(p Pos, dist int) {
		// Debug: Print current position and distance
		logger.Debugln("Visiting:", p, "Distance:", dist)
		// If the position is out of bounds, return
		if p.x < 0 || p.x >= len(grid[0]) || p.y < 0 || p.y >= len(grid) {
			return
		}

		// If the position was visited with a shorter or equal path, return
		if existingDist, visited := distances[p]; visited && existingDist <= dist {
			return
		}

		// Update current position distance
		distances[p] = dist

		// Recursively explore adjacent pipes
		if grid[p.y][p.x] != '.' { // Ignore ground tiles
			for _, dir := range []Pos{{0, 1}, {1, 0}, {0, -1}, {-1, 0}} {
				nextPos := Pos{p.x + dir.x, p.y + dir.y}
				if isValidNextPos(grid, p, nextPos) {
					dfs(nextPos, dist+1)
				}
			}
		}
	}

	// Start DFS from initial position
	dfs(startPos, 0)
}

// isValidNextPos checks if moving from currPos to nextPos is valid based on the
// pipe rules.
func isValidNextPos(grid [][]rune, currPos, nextPos Pos) bool {
	// Check if next position is out of bounds
	if nextPos.x < 0 || nextPos.x >= len(grid[0]) || nextPos.y < 0 || nextPos.y >= len(grid) {
		return false
	}

	// Get the symbols of the current and next positions
	currSym := grid[currPos.y][currPos.x]
	nextSym := grid[nextPos.y][nextPos.x]

	// Define the movement direction
	dx := nextPos.x - currPos.x
	dy := nextPos.y - currPos.y

	// Debug: Print check of next position validity
	logger.Debugln("Checking move from", currPos, "to", nextPos, "Current:", string(currSym), "Next:", string(nextSym))

	// Check if movement is allowed based on the type of the current and next pipe
	switch currSym {
	case '|':
		return (dy != 0) && (nextSym == '|' || nextSym == '7' || nextSym == 'F' || nextSym == 'S')
	case '-':
		return (dx != 0) && (nextSym == '-' || nextSym == 'L' || nextSym == 'J' || nextSym == 'S')
	case 'L':
		return ((dy == -1 && nextSym == '|') || (dx == 1 && nextSym == '-')) || nextSym == 'S'
	case 'J':
		return ((dy == -1 && nextSym == '|') || (dx == -1 && nextSym == '-')) || nextSym == 'S'
	case '7':
		return ((dy == 1 && nextSym == '|') || (dx == -1 && nextSym == '-')) || nextSym == 'S'
	case 'F':
		return ((dy == 1 && nextSym == '|') || (dx == 1 && nextSym == '-')) || nextSym == 'S'
	case 'S':
		return inferAndCheckPipeAtS(grid, currPos, nextPos)
	}
	return false
}

func inferAndCheckPipeAtS(grid [][]rune, currPos, nextPos Pos) bool {
	var up, down, left, right rune
	if currPos.y > 0 {
		up = grid[currPos.y-1][currPos.x]
	}
	if currPos.y < len(grid)-1 {
		down = grid[currPos.y+1][currPos.x]
	}
	if currPos.x > 0 {
		left = grid[currPos.y][currPos.x-1]
	}
	if currPos.x < len(grid[0])-1 {
		right = grid[currPos.y][currPos.x+1]
	}

	dx := nextPos.x - currPos.x
	dy := nextPos.y - currPos.y

	// Horizontal movement check
	if dx != 0 {
		if left == '-' || right == '-' || left == 'J' || right == 'L' || left == '7' || right == 'F' {
			nextSym := grid[nextPos.y][nextPos.x]
			return nextSym == '-' || nextSym == 'L' || nextSym == 'J' || nextSym == '7' || nextSym == 'F' || nextSym == 'S'
		}
	}

	// Vertical movement check
	if dy != 0 {
		if up == '|' || down == '|' || up == 'L' || down == 'J' || up == 'F' || down == '7' {
			nextSym := grid[nextPos.y][nextPos.x]
			return nextSym == '|' || nextSym == 'L' || nextSym == 'J' || nextSym == '7' || nextSym == 'F' || nextSym == 'S'
		}
	}

	return false
}

// findMaxDistance finds the maximum distance from the start position.
func findMaxDistance(distances map[Pos]int) int {
	maxDist := 0
	for _, dist := range distances {
		if dist > maxDist {
			maxDist = dist
		}
	}
	return maxDist
}

// Part2 ...
func Part2(input []string) (int, error) {
	start := time.Now()
	sum := 0

	logger.Infoln("Part 2 took:", time.Since(start))
	return sum, nil
}
