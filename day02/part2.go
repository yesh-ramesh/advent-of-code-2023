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

	scanner := bufio.NewScanner(file)

	sum := 0
	for scanner.Scan() {
		game := strings.Split(scanner.Text(), ":")[1]

		numHandfuls := strings.Count(game, ";") + 1
		handfuls := strings.SplitN(game, ";", numHandfuls)

		max := map[string]int{
			"red":   0,
			"green": 0,
			"blue":  0,
		}
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

				if count > max[color] {
					max[color] = count
				}
			}
		}

		sum += max["red"] * max["blue"] * max["green"]
	}

	fmt.Println(sum)
}
