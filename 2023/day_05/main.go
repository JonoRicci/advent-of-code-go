// Package main solves the Advent of Code 2023 Day 05 problem.
package main

import (
	"fmt"
	"jonoricci/advent-of-code-go/common"
	"log"
	"math"
	"strconv"
	"strings"
	"time"

	"go.uber.org/zap"
)

// global variable for logging
var logger *zap.SugaredLogger

// Using global var for input header names
var sectionHeaders = []string{
	"seed-to-soil map:",
	"soil-to-fertilizer map:",
	"fertilizer-to-water map:",
	"water-to-light map:",
	"light-to-temperature map:",
	"temperature-to-humidity map:",
	"humidity-to-location map:",
}

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
		logger.Fatalf("Couldn't read input file: %v", err)
	}

	// Split into lines
	inputData := strings.Split(strings.Trim(input, " "), "\n")
	// Remove empty strings from the input
	values := common.RemoveEmptyStrings(inputData)

	// Execute Part 1
	part1, err := Part1(values)
	if err != nil {
		logger.Fatalf("Part 1: %v", err)
	}

	// Execute Part 2
	part2, err := Part2(values)
	if err != nil {
		logger.Fatalf("Part 2: %v", err)
	}

	logger.Infoln("Part 1:", part1)
	logger.Infoln("Part 2:", part2)
}

// Part1 treats each seed as an individual integer and finds the lowest location
// number corresponding to these seeds.
func Part1(input []string) (int, error) {
	start := time.Now()

	// Pass `false` to parseInputData to use seeds as individual ints
	seeds, parsedMaps, err := parseInputData(input, false)
	if err != nil {
		return 0, fmt.Errorf("error in Part1: %w", err)
	}

	lowestLocation := processSeeds(seeds, parsedMaps)

	logger.Infoln("Part 1 took:", time.Since(start))
	return lowestLocation, nil
}

// processSeeds processes a list of seeds through a series of maps defined in
// parsedMaps. It calculates and returns the lowest location number obtained
// from these mappings.
func processSeeds(seeds []int, parsedMaps map[string][]RangeMap) int {
	lowestLocation := math.MaxInt32

	for _, seed := range seeds {
		logger.Debug("Processing seed:", seed)
		location := seed

		for _, header := range sectionHeaders {
			logger.Debug("Processing header:", header)
			location = applyMapping(location, parsedMaps[header])
		}

		if location < lowestLocation {
			lowestLocation = location
			logger.Debug("Lowest location so far is:", lowestLocation)
		}
	}

	logger.Debug("Lowest location is:", lowestLocation)
	return lowestLocation
}

// applyMapping applies a single RangeMap to a seed number and returns the
// corresponding number in the destination range. If no mapping is found,
// it returns the original number.
func applyMapping(number int, mappings []RangeMap) int {
	for _, m := range mappings {
		if number >= m.SourceStart && number < m.SourceStart+m.Length {
			return m.DestStart + (number - m.SourceStart)
		}
	}
	return number // No mapping found, return original number
}

// parseInputData parses the input data into seeds and a series of mappings.
// It takes a flag seedRange to determine whether to treat seeds as individual
// numbers or as ranges. It returns a slice of seed numbers and a map of
// RangeMaps for each mapping step.
func parseInputData(input []string, seedRange bool) ([]int, map[string][]RangeMap, error) {
	// Get the seeds
	seeds, err := extractSeeds(input[0], seedRange)
	if err != nil {
		logger.Fatalln("Parsing seeds:", err)
	}

	// Map for storing the parsed maps
	parsedMaps := make(map[string][]RangeMap)

	var currentHeader string
	var currentSection []string

	for _, line := range input[1:] { // Skip first line containing seeds
		if isHeader(line, sectionHeaders) {
			if currentHeader != "" {
				// Parse previous section
				parsedMaps[currentHeader], err = parseMap(currentSection)
				if err != nil {
					logger.Fatalln("Parsing %S:", currentHeader, err)
				}
			}
			// Reset for new section
			currentHeader = line
			currentSection = []string{}
		} else {
			currentSection = append(currentSection, line)
		}
	}

	// Parse last section
	if currentHeader != "" {
		parsedMaps[currentHeader], err = parseMap(currentSection)
		if err != nil {
			return nil, nil, fmt.Errorf("parsing %s, %v", currentHeader, err)
		}
	}

	return seeds, parsedMaps, nil
}

// RangeMap defines a mapping from a source range to a destination range in
// the puzzle. It specifies where the source range starts, where the destination
// range starts, and the length of these ranges.
type RangeMap struct {
	SourceStart int
	DestStart   int
	Length      int
}

// extractSeeds processes the first line of input to extract seed numbers.
// It handles two formats: individual seeds or ranges of seeds, based on the
// isRangeFormat flag. It returns a slice of all seed numbers to be processed.
func extractSeeds(input string, isRangeFormat bool) ([]int, error) {
	parts := strings.SplitN(input, ":", 2)
	if len(parts) != 2 {
		return nil, fmt.Errorf("invalid seed input format: %s", input)
	}

	numberParts := strings.Fields(parts[1])
	var seeds []int

	if isRangeFormat {
		// Process as ranges (Part Two format)
		for i := 0; i < len(numberParts); i += 2 {
			start, err := strconv.Atoi(numberParts[i])
			if err != nil {
				return nil, err
			}
			logger.Debug("Found seed:", start)
			length, err := strconv.Atoi(numberParts[i+1])
			if err != nil {
				return nil, err
			}
			logger.Debug("Seed has length of:", length)
			for j := 0; j < length; j++ {
				seeds = append(seeds, start+j)
			}
		}
	} else {
		// Process each number as an individual seed (Part One format)
		for _, part := range numberParts {
			seed, err := strconv.Atoi(part)
			if err != nil {
				return nil, err
			}
			logger.Debug("Found seed:", seed)
			seeds = append(seeds, seed)
		}
	}
	return seeds, nil
}

// isHeader checks if a given line in the input matches any of the predefined
// section headers. It is used to identify the start of a new mapping section
// in the input.
func isHeader(line string, sectionHeaders []string) bool {
	for _, header := range sectionHeaders {
		if line == header {
			return true
		}
	}
	return false
}

// parseMap parses a slice of strings representing lines of a map section.
// It converts each line into a RangeMap struct, which defines the mappings
// from one category to another. It returns a slice of RangeMaps for the
// section.
func parseMap(lines []string) ([]RangeMap, error) {
	var maps []RangeMap
	for _, line := range lines {
		parts := strings.Fields(line)
		if len(parts) != 3 {
			return nil, fmt.Errorf("invalid map line: %s", line)
		}
		sourceStart, err1 := strconv.Atoi(parts[1])
		destStart, err2 := strconv.Atoi(parts[0])
		length, err3 := strconv.Atoi(parts[2])
		if err1 != nil || err2 != nil || err3 != nil {
			return nil, fmt.Errorf("error parsing map line: %s", line)
		}
		maps = append(maps, RangeMap{
			SourceStart: sourceStart,
			DestStart:   destStart,
			Length:      length,
		})
	}
	return maps, nil
}

// Part2 treats the seeds as ranges, expands these ranges into individual seeds,
// and finds the lowest location number. This took me at least 5 minutes to
// execute so performance is not the best.
func Part2(input []string) (int, error) {
	start := time.Now()

	// Pass `true` to parseInputData to use seeds as a range
	seeds, parsedMaps, err := parseInputData(input, true)
	if err != nil {
		return 0, fmt.Errorf("error in Part2: %w", err)
	}

	lowestLocation := processSeeds(seeds, parsedMaps)

	logger.Infoln("Part 2 took:", time.Since(start))
	return lowestLocation, nil
}
