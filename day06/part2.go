package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
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
	inputTime := strings.Join(strings.Fields(strings.Split(scanner.Text(), ":")[1]), "")
	time, err := strconv.Atoi(inputTime)
	if err != nil {
		log.Fatal(err)
	}

	scanner.Scan()
	inputDistance := strings.Join(strings.Fields(strings.Split(scanner.Text(), ":")[1]), "")
	recordDistance, err := strconv.Atoi(inputDistance)
	if err != nil {
		log.Fatal(err)
	}

	lowestPossible := math.MaxInt32
	highestPossible := math.MinInt32
	for i := 0; i <= time; i++ {
		distanceTraveled := (time - i) * i

		if distanceTraveled > recordDistance {
			fmt.Printf("Set a record by holding the button %d milliseconds. You'll travel %d millimeters which is greater than the record %d!\n", i, distanceTraveled, recordDistance)
			lowestPossible = i
			break
		}
	}

	for i := time; i >= 0; i-- {
		distanceTraveled := (time - i) * i

		if distanceTraveled > recordDistance {
			fmt.Printf("Set a record by holding the button %d milliseconds. You'll travel %d millimeters which is greater than the record %d!\n", i, distanceTraveled, recordDistance)
			highestPossible = i
			break
		}
	}

	possibleWins := highestPossible - lowestPossible + 1

	fmt.Printf("\nThere are %d ways to win this race\n\n", possibleWins)
}
