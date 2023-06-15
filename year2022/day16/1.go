package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

func dfs(roads [][]int, values []int, visited []bool, start int, oldValue int) {
	visited[start] = true
	for i := range roads {
		if roads[start][i] == 1 {
			if (oldValue+1 < values[i] && visited[i]) || !visited[i] {
				values[i] = oldValue + 1
				dfs(roads, values, visited, i, oldValue+1)
			}
		}
	}
}

func processInput(valves map[string]int, flowRates *[]int, roadLengths *[][]int) {
	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatalln("opening data file: ", err)
	}

	reLabel := regexp.MustCompile(`Valve (..) has`)
	reRate := regexp.MustCompile(`rate=(\d+);`)
	reGoto := regexp.MustCompile(`lead.? to valve.? (.*)`)
	reLabelGoto := regexp.MustCompile(`[A-Z][A-Z]`)

	var directAccessLabels [][]string

	scanner := bufio.NewScanner(file)
	valvesNumber := 0
	for ; scanner.Scan(); valvesNumber++ {
		// Get valve label
		label := reLabel.FindSubmatch([]byte(scanner.Text()))[1]
		valves[string(label)] = valvesNumber

		// Get the rate for a valve
		rateString := reRate.FindSubmatch([]byte(scanner.Text()))[1]

		rate, err := strconv.Atoi(string(rateString))
		if err != nil {
			log.Fatalln("converting rate to intger: ", err)
		}

		*flowRates = append(*flowRates, rate)

		stringGoto := string(reGoto.FindSubmatch([]byte(scanner.Text()))[1])
		directAccessLabels = append(directAccessLabels, reLabelGoto.FindAllString(stringGoto, -1))
	}
	if err := scanner.Err(); err != nil {
		log.Fatalln("reading from data file: ", err)
	}

	// Adjacency matrix for valves
	var availableRoads [][]int

	for i := 0; i < valvesNumber; i++ {
		currentRow := make([]int, valvesNumber)

		for _, entry := range directAccessLabels[i] {
			currentRow[valves[entry]] = 1
		}

		availableRoads = append(availableRoads, currentRow)
	}

	for i := 0; i < valvesNumber; i++ {
		visited := make([]bool, valvesNumber)
		values := make([]int, valvesNumber)
		dfs(availableRoads, values, visited, i, 0)
		*roadLengths = append(*roadLengths, values)
	}

}

func bkt(roadLengths [][]int, flowRates []int, bestValue *int, value *int, timeRemaining int, position int) {
	for i := 0; i < len(flowRates); i++ {
		if i == position {
			continue
		}

		if dt := timeRemaining - roadLengths[position][i] - 1; flowRates[i] != 0 && dt > 0 {
			flowRatesCopy := flowRates[i]
			*value += dt * flowRates[i]
			if *value > *bestValue {
				*bestValue = *value
			}
			flowRates[i] = 0
			bkt(roadLengths, flowRates, bestValue, value, dt, i)
			flowRates[i] = flowRatesCopy
			*value -= dt * flowRates[i]
		}
	}
}

func main() {
	var valves map[string]int
	var flowRates []int
	var roadLengths [][]int

	valves = make(map[string]int)

	processInput(valves, &flowRates, &roadLengths)

	var bestValue int
	var value int
	bkt(roadLengths, flowRates, &bestValue, &value, 30, valves["AA"])
	fmt.Println(bestValue)
}
