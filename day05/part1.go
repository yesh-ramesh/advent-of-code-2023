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
	file, err := os.Open("day05/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)

	scanner.Scan()
	seeds := strings.Fields(strings.Split(scanner.Text(), ":")[1])

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

	locations := []int{}
	for _, seedStr := range seeds {
		seed, err := strconv.Atoi(seedStr)
		if err != nil {
			log.Fatal(err)
		}

		soil := processMapping(seed, seedToSoil)
		fertilizer := processMapping(soil, soilToFertilizer)
		water := processMapping(fertilizer, fertilizerToWater)
		light := processMapping(water, waterToLight)
		temperature := processMapping(light, lightToTemperature)
		humidity := processMapping(temperature, temperatureToHumidity)
		location := processMapping(humidity, humidityToLocation)

		fmt.Printf("seed: %v\n", seed)
		fmt.Printf("soil: %v\n", soil)
		fmt.Printf("fertilizer: %v\n", fertilizer)
		fmt.Printf("water: %v\n", water)
		fmt.Printf("light: %v\n", light)
		fmt.Printf("temperature: %v\n", temperature)
		fmt.Printf("humidity: %v\n", humidity)
		fmt.Printf("location: %v\n", location)
		fmt.Println("-------------------------------")

		locations = append(locations, location)
	}

	min := math.MaxInt32
	for _, l := range locations {
		if l < min {
			min = l
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
