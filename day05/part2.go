package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
	"sync"
)

var wg sync.WaitGroup

func main() {
	file, err := os.Open("day05/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)

	scanner.Scan()
	seeds := [][2]int{}
	seedText := strings.Fields(strings.Split(scanner.Text(), ":")[1])
	for i := 0; i < len(seedText); i += 2 {
		seedRange := [2]int{}

		seedStart, err := strconv.Atoi(seedText[i])
		if err != nil {
			log.Fatal(err)
		}
		seedRange[0] = seedStart

		seedEnd, err := strconv.Atoi(seedText[i+1])
		if err != nil {
			log.Fatal(err)
		}
		seedRange[1] = seedEnd

		seeds = append(seeds, seedRange)
	}

	seedToSoil := parseMapping(scanner, "seed-to-soil")
	soilToFertilizer := parseMapping(scanner, "soil-to-fertilizer")
	fertilizerToWater := parseMapping(scanner, "fertilizer-to-water")
	waterToLight := parseMapping(scanner, "water-to-light")
	lightToTemperature := parseMapping(scanner, "light-to-temperature")
	temperatureToHumidity := parseMapping(scanner, "temperature-to-humidity")
	humidityToLocation := parseMapping(scanner, "humidity-to-location")

	fmt.Printf("seeds: %v\n", seeds)
	fmt.Printf("seedToSoil: %v\n", seedToSoil)
	fmt.Printf("soilToFertilizer: %v\n", soilToFertilizer)
	fmt.Printf("fertilizerToWater: %v\n", fertilizerToWater)
	fmt.Printf("waterToLight: %v\n", waterToLight)
	fmt.Printf("lightToTemperature: %v\n", lightToTemperature)
	fmt.Printf("temperatureToHumidity: %v\n", temperatureToHumidity)
	fmt.Printf("humidityToLocation: %v\n", humidityToLocation)

	fmt.Println("-------------------------------")

	ch := make(chan int, 100000)
	for _, seed := range seeds {
		for i := seed[0]; i <= seed[0]+seed[1]; i++ {
			wg.Add(1)
			i := i

			go func() {
				defer wg.Done()
				soil := processMapping(i, seedToSoil)
				fertilizer := processMapping(soil, soilToFertilizer)
				water := processMapping(fertilizer, fertilizerToWater)
				light := processMapping(water, waterToLight)
				temperature := processMapping(light, lightToTemperature)
				humidity := processMapping(temperature, temperatureToHumidity)
				location := processMapping(humidity, humidityToLocation)

				//fmt.Printf("seed: %v\n", seed)
				//fmt.Printf("soil: %v\n", soil)
				//fmt.Printf("fertilizer: %v\n", fertilizer)
				//fmt.Printf("water: %v\n", water)
				//fmt.Printf("light: %v\n", light)
				//fmt.Printf("temperature: %v\n", temperature)
				//fmt.Printf("humidity: %v\n", humidity)
				//fmt.Printf("location: %v\n", location)
				//fmt.Println("-------------------------------")

				ch <- location
			}()
		}
	}

	wg.Wait()
	close(ch)
	min := math.MaxInt32
	for loc := range ch {
		if loc < min {
			fmt.Printf("Checking location: %v, current min: %v\n", loc, min)
			min = loc
		}
	}
	fmt.Printf("Lowest Location: %v\n", min)
}

func parseMapping(scanner *bufio.Scanner, mappingName string) [][3]int {
	mapping := [][3]int{}
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, mappingName) {
			for scanner.Scan() && scanner.Text() != "" {
				line := scanner.Text()
				strs := strings.Fields(line)
				nums := [3]int{}
				for i := 0; i < 3; i++ {
					num, err := strconv.Atoi(strs[i])
					if err != nil {
						log.Fatal(err)
					}

					nums[i] = num
				}
				mapping = append(mapping, nums)
			}

			return mapping
		}
	}

	return nil
}

func processMapping(startNum int, mapping [][3]int) int {
	resultNum := -1
	unmapped := true
	for _, entry := range mapping {
		if startNum >= entry[1] && startNum <= entry[1]+entry[2]-1 {
			diff := startNum - entry[1]
			resultNum = entry[0] + diff
			unmapped = false
		}
	}
	if unmapped {
		resultNum = startNum
	}

	return resultNum
}
