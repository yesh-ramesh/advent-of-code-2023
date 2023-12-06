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
	file, err := os.Open("day01/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	m := map[int]string{
		0: "zero",
		1: "one",
		2: "two",
		3: "three",
		4: "four",
		5: "five",
		6: "six",
		7: "seven",
		8: "eight",
		9: "nine",
	}

	sum := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		digits := [2]int{-1, -1}

		first := math.MaxInt32
		last := math.MinInt32
		for k, v := range m {
			i := strings.Index(line, strconv.Itoa(k))
			if i != -1 && i < first {
				first = i
				digits[0] = k
			}

			i = strings.Index(line, v)
			if i != -1 && i < first {
				first = i
				digits[0] = k
			}

			i = strings.LastIndex(line, strconv.Itoa(k))
			if i != -1 && i > last {
				last = i
				digits[1] = k
			}

			i = strings.LastIndex(line, v)
			if i != -1 && i > last {
				last = i
				digits[1] = k
			}
		}

		sum += digits[0]*10 + digits[1]
	}

	fmt.Println(sum)
}
