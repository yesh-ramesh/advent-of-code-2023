package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"regexp"
	"strconv"
)

const DAY3_PART2_GRID_SIZE = 140

type part struct {
	number string
	row    int
	col    int
}

func main() {
	// Grab the input
	file, err := os.Open("day03/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)

	rows := [DAY3_PART2_GRID_SIZE]string{}
	index := 0
	for scanner.Scan() {
		line := scanner.Text()
		rows[index] = line
		index++
		fmt.Println(line)
	}

	// Save all numbers and gears along with their grid positions
	digitRegex := regexp.MustCompile("[0-9]+")
	gearRegex := regexp.MustCompile("[*]")

	numbers := []part{}
	gears := []part{}
	for rowIndex, row := range rows {
		numberMatches := digitRegex.FindAllStringIndex(row, -1)
		gearMatches := gearRegex.FindAllStringIndex(row, -1)

		for _, number := range numberMatches {
			n := part{number: row[number[0]:number[1]], row: rowIndex, col: number[0]}
			numbers = append(numbers, n)
		}

		for _, gear := range gearMatches {
			g := part{number: row[gear[0]:gear[1]], row: rowIndex, col: gear[0]}
			gears = append(gears, g)
		}
	}

	// Go through all gear and numbers, check for neighbors, and calculate sum of gear ratios
	sum := 0
	for _, gear := range gears {
		neighbors := []int{}
		for _, number := range numbers {
			if math.Abs(float64(gear.row-number.row)) <= 1 &&
				number.col <= gear.col+len(gear.number) &&
				gear.col <= number.col+len(number.number) {

				num, err := strconv.Atoi(number.number)
				if err != nil {
					log.Fatal(err)
				}
				neighbors = append(neighbors, num)
			}
		}

		if len(neighbors) == 2 {
			sum += neighbors[0] * neighbors[1]
		}
	}

	fmt.Printf("%v \n\n", sum)
}
