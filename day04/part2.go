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

	cardCountMap := map[int]int{}

	scanner := bufio.NewScanner(file)
	index := 1
	for scanner.Scan() {
		line := scanner.Text()
		cardCountMap[index]++
		fmt.Printf("card: %d\n", index)

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

		numWinning := 0
		for _, numStr := range ours {
			num, err := strconv.Atoi(numStr)
			if err != nil {
				log.Fatal(err)
			}

			if winningMap[num] {
				numWinning++
			}
		}

		fmt.Printf("num wins: %d\n", numWinning)

		for i := index + 1; i <= index+numWinning; i++ {
			cardCountMap[i] += cardCountMap[index]
		}

		fmt.Printf("%v\n", cardCountMap)
		fmt.Printf("-------------\n")

		index++
	}

	total := 0
	for _, v := range cardCountMap {
		total += v
	}
	fmt.Println(total)
}
