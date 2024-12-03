package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	// Create maps for ranks
	cardRanks := map[rune]int{
		'2': 2,
		'3': 3,
		'4': 4,
		'5': 5,
		'6': 6,
		'7': 7,
		'8': 8,
		'9': 9,
		'T': 10,
		'J': 11,
		'Q': 12,
		'K': 13,
		'A': 14,
	}

	handRanks := map[string]int{
		"high-card":       0,
		"one-pair":        1,
		"two-pair":        2,
		"three-of-a-kind": 3,
		"full-house":      4,
		"four-of-a-kind":  5,
		"five-of-a-kind":  6,
	}

	var hands []string
	handToRank := make(map[string]string)
	handToBid := make(map[string]int)

	// Parse input and save
	file, err := os.Open("day07/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		split := strings.Fields(line)

		handStr := split[0]
		hands = append(hands, handStr)
		bid := split[1]
		handToBid[handStr], _ = strconv.Atoi(bid)

		// Create frequency map from hand
		hand := []rune(handStr)
		freqMap := make(map[rune]int)
		for _, card := range []rune(hand) {
			freqMap[card]++
		}

		// Use frequency map to determine hand type
		if len(freqMap) == 5 {
			// High card
			handToRank[handStr] = "high-card"
		} else if len(freqMap) == 4 {
			// One Pair
			handToRank[handStr] = "one-pair"
		} else if len(freqMap) == 3 {
			// Three of a Kind or Two Pair
			if freqMap[hand[0]] == 3 || freqMap[hand[1]] == 3 || freqMap[hand[2]] == 3 {
				handToRank[handStr] = "three-of-a-kind"
			} else {
				handToRank[handStr] = "two-pair"
			}
		} else if len(freqMap) == 2 {
			// Full House or Four of a Kind
			if freqMap[hand[0]] == 3 || freqMap[hand[0]] == 2 {
				handToRank[handStr] = "full-house"
			} else {
				handToRank[handStr] = "four-of-a-kind"
			}
		} else if len(freqMap) == 1 {
			// Five of a Kind
			handToRank[handStr] = "five-of-a-kind"
		}
	}

	// Sort hands based on hand type ranking and card ranking
	slices.SortFunc(hands, func(a, b string) int {
		if handRanks[handToRank[a]] > handRanks[handToRank[b]] {
			return 1
		} else if handRanks[handToRank[a]] < handRanks[handToRank[b]] {
			return -1
		} else {
			aRunes := []rune(a)
			bRunes := []rune(b)
			for i := 0; i < 5; i++ {
				if cardRanks[aRunes[i]] > cardRanks[bRunes[i]] {
					return 1
				} else if cardRanks[aRunes[i]] < cardRanks[bRunes[i]] {
					return -1
				}
			}

			return 0
		}
	})

	winnings := 0
	for i, hand := range hands {
		winnings += (i + 1) * handToBid[hand]
	}

	fmt.Println(winnings)
}
