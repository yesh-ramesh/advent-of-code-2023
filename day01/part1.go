package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"unicode"
)

func main() {
	file, err := os.Open("day01/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	sum := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		numString := ""
		for _, r := range scanner.Text() {
			if unicode.IsDigit(r) {
				if len(numString) == 0 {
					numString += string(r)
				} else {
					numString = numString[:1]
					numString += string(r)
				}
			}
		}

		if len(numString) == 1 {
			numString += numString
		}

		num, err := strconv.Atoi(numString)
		if err != nil {
			log.Fatal(err)
		}

		sum += num
	}

	fmt.Println(sum)
}
