package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"unicode"
)

const DAY3_PART1_GRID_SIZE = 10

func main() {
	file, err := os.Open("day03/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	// BRUTE FORCE
	// Convert the input into a 2D array or runes
	grid := [DAY3_PART1_GRID_SIZE][DAY3_PART1_GRID_SIZE]rune{}
	scanner := bufio.NewScanner(file)
	i := 0
	for scanner.Scan() {
		line := scanner.Text()
		j := 0
		for _, c := range line {
			grid[i][j] = c
			j++
		}
		i++
	}

	// Go through the grid capturing digits and flag if a digit neighbors a symbol
	// If it does, then the number formed by the digits is a valid part number
	partNumbers := []int{}
	numberStr := ""
	symbolAdjacent := false
	sum := 0
	for i := 0; i < DAY3_PART1_GRID_SIZE; i++ {
		for j := 0; j < DAY3_PART1_GRID_SIZE; j++ {
			c := grid[i][j]
			if unicode.IsDigit(c) {
				numberStr = numberStr + string(c)

				// check if the current digit neighbors a symbol
				neighbors := []rune{}

				if i == 0 && j == 0 { // top left corner
					neighbors = append(neighbors, grid[i+1][j])
					neighbors = append(neighbors, grid[i][j+1])
					neighbors = append(neighbors, grid[i+1][j+1])
				} else if i == 0 && j == DAY3_PART1_GRID_SIZE-1 { // top right corner
					neighbors = append(neighbors, grid[i+1][j])
					neighbors = append(neighbors, grid[i][j-1])
					neighbors = append(neighbors, grid[i+1][j-1])
				} else if i == DAY3_PART1_GRID_SIZE-1 && j == 0 { // bottom left corner
					neighbors = append(neighbors, grid[i-1][j])
					neighbors = append(neighbors, grid[i][j+1])
					neighbors = append(neighbors, grid[i-1][j+1])
				} else if i == DAY3_PART1_GRID_SIZE-1 && j == DAY3_PART1_GRID_SIZE-1 { // bottom right corner
					neighbors = append(neighbors, grid[i-1][j])
					neighbors = append(neighbors, grid[i][j-1])
					neighbors = append(neighbors, grid[i-1][j-1])
				} else if i == 0 { // top edge
					neighbors = append(neighbors, grid[i+1][j])
					neighbors = append(neighbors, grid[i][j-1])
					neighbors = append(neighbors, grid[i][j+1])
					neighbors = append(neighbors, grid[i+1][j+1])
					neighbors = append(neighbors, grid[i+1][j-1])
				} else if i == DAY3_PART1_GRID_SIZE-1 { // bottom edge
					neighbors = append(neighbors, grid[i-1][j])
					neighbors = append(neighbors, grid[i][j-1])
					neighbors = append(neighbors, grid[i][j+1])
					neighbors = append(neighbors, grid[i-1][j+1])
					neighbors = append(neighbors, grid[i-1][j-1])
				} else if j == 0 { // left edge
					neighbors = append(neighbors, grid[i-1][j])
					neighbors = append(neighbors, grid[i+1][j])
					neighbors = append(neighbors, grid[i][j+1])
					neighbors = append(neighbors, grid[i+1][j+1])
					neighbors = append(neighbors, grid[i-1][j+1])
				} else if j == DAY3_PART1_GRID_SIZE-1 { // right edge
					neighbors = append(neighbors, grid[i-1][j])
					neighbors = append(neighbors, grid[i+1][j])
					neighbors = append(neighbors, grid[i][j-1])
					neighbors = append(neighbors, grid[i+1][j-1])
					neighbors = append(neighbors, grid[i-1][j-1])
				} else if i > 0 && j < DAY3_PART1_GRID_SIZE-1 { // not corner, not edge
					neighbors = append(neighbors, grid[i-1][j])
					neighbors = append(neighbors, grid[i+1][j])
					neighbors = append(neighbors, grid[i][j-1])
					neighbors = append(neighbors, grid[i][j+1])
					neighbors = append(neighbors, grid[i+1][j+1])
					neighbors = append(neighbors, grid[i+1][j-1])
					neighbors = append(neighbors, grid[i-1][j+1])
					neighbors = append(neighbors, grid[i-1][j-1])
				} else if i == 0 && j == 0 { // top left corner
					neighbors = append(neighbors, grid[i+1][j])
					neighbors = append(neighbors, grid[i][j+1])
					neighbors = append(neighbors, grid[i+1][j+1])
				} else if i == 0 && j == DAY3_PART1_GRID_SIZE-1 { // top right corner
					neighbors = append(neighbors, grid[i+1][j])
					neighbors = append(neighbors, grid[i][j-1])
					neighbors = append(neighbors, grid[i+1][j-1])
				} else if i == DAY3_PART1_GRID_SIZE-1 && j == 0 { // bottom left corner
					neighbors = append(neighbors, grid[i-1][j])
					neighbors = append(neighbors, grid[i][j+1])
					neighbors = append(neighbors, grid[i-1][j+1])
				} else if i == DAY3_PART1_GRID_SIZE-1 && j == DAY3_PART1_GRID_SIZE-1 { // bottom right corner
					neighbors = append(neighbors, grid[i-1][j])
					neighbors = append(neighbors, grid[i][j-1])
					neighbors = append(neighbors, grid[i-1][j-1])
				} else if i == 0 { // top edge
					neighbors = append(neighbors, grid[i+1][j])
					neighbors = append(neighbors, grid[i][j-1])
					neighbors = append(neighbors, grid[i][j+1])
					neighbors = append(neighbors, grid[i+1][j+1])
					neighbors = append(neighbors, grid[i+1][j-1])
				} else if i == DAY3_PART1_GRID_SIZE-1 { // bottom edge
					neighbors = append(neighbors, grid[i-1][j])
					neighbors = append(neighbors, grid[i][j-1])
					neighbors = append(neighbors, grid[i][j+1])
					neighbors = append(neighbors, grid[i-1][j+1])
					neighbors = append(neighbors, grid[i-1][j-1])
				} else if j == 0 { // left edge
					neighbors = append(neighbors, grid[i-1][j])
					neighbors = append(neighbors, grid[i+1][j])
					neighbors = append(neighbors, grid[i][j+1])
					neighbors = append(neighbors, grid[i+1][j+1])
					neighbors = append(neighbors, grid[i-1][j+1])
				} else if j == DAY3_PART1_GRID_SIZE-1 { // right edge
					neighbors = append(neighbors, grid[i-1][j])
					neighbors = append(neighbors, grid[i+1][j])
					neighbors = append(neighbors, grid[i][j-1])
					neighbors = append(neighbors, grid[i+1][j-1])
					neighbors = append(neighbors, grid[i-1][j-1])
				}

				for _, r := range neighbors {
					if !unicode.IsDigit(r) && r != '.' {
						symbolAdjacent = true
					}
				}
			} else {
				if symbolAdjacent {
					number, err := strconv.Atoi(numberStr)
					if err != nil {
						log.Fatal(err)
					}

					partNumbers = append(partNumbers, number)
					symbolAdjacent = false
				}
				numberStr = ""
			}
		}
	}

	for _, partNumber := range partNumbers {
		sum += partNumber
	}
	fmt.Println(sum)
}
