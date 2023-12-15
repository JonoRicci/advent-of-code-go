// Package main solves the Advent of Code 2023 Day 07 problem.
package main

import (
	"fmt"
	"jonoricci/advent-of-code-go/common"
	"log"
	"sort"
	"strconv"
	"strings"
	"time"

	"go.uber.org/zap"
)

// global variable for logging
var logger *zap.SugaredLogger

// global variable for card strength
// Add a map to define the strength of each card
var cardStrengthMap = map[rune]int{
	'A': 13,
	'K': 12,
	'Q': 11,
	'J': 10,
	'T': 9,
	'9': 8,
	'8': 7,
	'7': 6,
	'6': 5,
	'5': 4,
	'4': 3,
	'3': 2,
	'2': 1,
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

	// Execute Part 2
	part2, err := Part2(values)
	if err != nil {
		logger.Fatalln(err)
	}

	logger.Infoln("Part 1:", part1)
	logger.Infoln("Part 2:", part2)
}

// Part1 calculates the total winnings based on Camel Cards game rules
func Part1(input []string) (int, error) {
	start := time.Now()
	sum := 0

	// Create slice of structs to hold hand data
	type handData struct {
		hand       string
		bid        int
		handType   int
		sortedHand []string
	}
	hands := make([]handData, 0)

	// Parse and evaluate each hand and bid
	for _, line := range input {
		// logger.Debugln("Line:", line)
		splitLine := strings.Split(line, " ")
		hand := splitLine[0]
		bid, err := strconv.Atoi(splitLine[1])
		if err != nil {
			return 0, fmt.Errorf("failed to convert bid to int: %v", err)
		}

		// Get the hand type and sort the hand left to right by card strength
		handType, sortedHand := evaluateHand(hand)
		hands = append(hands, handData{hand, bid, handType, sortedHand})
		// logger.Debugln("Hand:", hand, "Bid:", bid, "Type:", handType, "Sorted:", sortedHand)
	}

	// Sort hands based on type and card strength.
	// Useing an anonymous function here, known as lambda function in Python. Returns a true if hand is higher.
	sort.Slice(hands, func(i, j int) bool {
		// Compare hand type and return highest hand.
		if hands[i].handType != hands[j].handType {
			return hands[i].handType > hands[j].handType
		}

		// If the hand types are the same, then compare higher cards.
		// Iterate through each card in the sorted order.
		for k := 0; k < len(hands[i].hand) && k < len(hands[j].hand); k++ {
			if cardStrength(rune(hands[i].hand[k])) != cardStrength(rune(hands[j].hand[k])) {
				return cardStrength(rune(hands[i].hand[k])) > cardStrength(rune(hands[j].hand[k]))
			}
		}
		logger.Warnln("Potential tie detected between hands:", hands[i].hand, "and", hands[j].hand)
		return false
	})

	// Calculate total winnings
	for i, hd := range hands {
		rank := len(hands) - i
		sum += hd.bid * rank
		logger.Debugln("Ranked Hand:", hd.hand, "Bid:", hd.bid, "Rank:", rank, "Score:", hd.bid*rank, "Sum:", sum)
	}

	logger.Infoln("Part 1 took:", time.Since(start))
	return sum, nil
}

// evaluateHand determines the type and strength of each hand
func evaluateHand(hand string) (int, []string) {
	counts := make(map[rune]int)
	var sortedCards []rune

	for _, card := range hand {
		counts[card]++
	}

	var handType int
	switch len(counts) {
	case 5: // High card or straight
		handType = 1
	case 4: // One pair
		handType = 2
	case 3: // Two pair or three of a kind
		for _, count := range counts {
			if count == 3 {
				handType = 4
				break
			}
		}
		if handType == 0 {
			handType = 3 // Two pair
		}
	case 2: // Full house or four of a kind
		for _, count := range counts {
			if count == 4 {
				handType = 6
				break
			}
		}
		if handType == 0 {
			handType = 5 // Full house
		}
	case 1: // Five of a kind
		handType = 7
	}

	// Populate sortedCards with all cards, sorted by frequency and rank
	for card, count := range counts {
		for i := 0; i < count; i++ {
			sortedCards = append(sortedCards, card)
		}
	}

	// Sort the sortedCards based on frequency and rank
	sort.Slice(sortedCards, func(i, j int) bool {
		countI := counts[sortedCards[i]]
		countJ := counts[sortedCards[j]]
		if countI == countJ {
			return cardStrength(sortedCards[i]) > cardStrength(sortedCards[j]) // Compare by card strength if frequency is the same
		}
		return countI > countJ // Compare by frequency
	})

	// Convert sortedCards to []string
	var sortedCardsStr []string
	for _, card := range sortedCards {
		sortedCardsStr = append(sortedCardsStr, string(card))
	}

	return handType, sortedCardsStr
}

// cardStrength returns the strength of a card based on its rank
func cardStrength(card rune) int {
	return cardStrengthMap[card]
}

// Part2 ...
func Part2(input []string) (int, error) {
	start := time.Now()
	sum := 0

	logger.Infoln("Part 2 took:", time.Since(start))
	return sum, nil
}
