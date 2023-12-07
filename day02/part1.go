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
	file, err := os.Open("day02/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	bag := map[string]int{
		"red":   12,
		"green": 13,
		"blue":  14,
	}

	scanner := bufio.NewScanner(file)
	index := 1

	sum := 0
	for scanner.Scan() {
		game := strings.Split(scanner.Text(), ":")[1]

		numHandfuls := strings.Count(game, ";") + 1
		handfuls := strings.SplitN(game, ";", numHandfuls)

		possible := true
		for _, handful := range handfuls {
			numColors := strings.Count(handful, ",") + 1
			colorSets := strings.SplitN(handful, ",", numColors)

			for _, colorSet := range colorSets {
				colorSet = strings.Trim(colorSet, " ")

				color := strings.Split(colorSet, " ")[1]
				count, err := strconv.Atoi(strings.Split(colorSet, " ")[0])
				if err != nil {
					log.Fatal(err)
				}

				if count > bag[color] {
					possible = false
				}
			}

		}

		if possible {
			sum += index
		}
		index++
	}

	fmt.Println(sum)
}
