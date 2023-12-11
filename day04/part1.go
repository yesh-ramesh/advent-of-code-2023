package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("day04/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	totalPoints := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		card := strings.Split(line, ": ")[1]
		numbers := strings.Split(card, " | ")
		winning := strings.Fields(numbers[0])
		ours := strings.Fields(numbers[1])

		winningMap := map[int]bool{}
		for _, numStr := range winning {
			num, err := strconv.Atoi(numStr)
			if err != nil {
				log.Fatal()
			}

			winningMap[num] = true
		}

		cardPoints := 0
		for _, numStr := range ours {
			num, err := strconv.Atoi(numStr)
			if err != nil {
				log.Fatal(err)
			}

			if winningMap[num] {
				if cardPoints == 0 {
					cardPoints++
				} else {
					cardPoints *= 2
				}
			}
		}

		totalPoints += cardPoints
	}

	fmt.Println(totalPoints)
}
