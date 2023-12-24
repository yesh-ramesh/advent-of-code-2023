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
	file, err := os.Open("day06/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)

	scanner.Scan()
	inputTimes := strings.Fields(strings.Split(scanner.Text(), ":")[1])
	scanner.Scan()
	inputDistances := strings.Fields(strings.Split(scanner.Text(), ":")[1])

	result := 1

	for index, time := range inputTimes {
		fmt.Printf("---- Processing Race %d ----\n", index+1)

		possibleWins := 0
		time, err := strconv.Atoi(time)
		if err != nil {
			log.Fatal(err)
		}

		recordDistance, err := strconv.Atoi(inputDistances[index])
		if err != nil {
			log.Fatal(err)
		}

		for i := 0; i <= time; i++ {
			distanceTraveled := (time - i) * i

			if distanceTraveled > recordDistance {
				fmt.Printf("Set a record by holding the button %d milliseconds. You'll travel %d millimeters which is greater than the record %d!\n", i, distanceTraveled, recordDistance)
				possibleWins++
			}
		}

		fmt.Printf("\nThere are %d ways to win this race\n\n", possibleWins)

		result *= possibleWins
	}

	fmt.Printf("Result %d\n", result)
}
