package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

func myVariant(enemy string, match string) string {
	switch match {
	case "X":
		switch enemy {
		case "A":
			return "C"
		case "B":
			return "A"
		case "C":
			return "B"
		}
	case "Z":
		switch enemy {
		case "A":
			return "B"
		case "B":
			return "C"
		case "C":
			return "A"
		}
	}

	// If I need to draw I will choose what enemy chose
	return enemy
}

func calculateScoreOfCompleteStrategy(fileName string) int {
	file, err := os.Open(fileName)

	if err != nil {
		log.Fatalln("The file with strategy can't be open =(", err)
	}

	var score int
	var pointsMatch map[string]int
	var pointsChoice map[string]int

	pointsChoice = map[string]int{
		"A": 1,
		"B": 2,
		"C": 3,
	}

	pointsMatch = map[string]int{
		"X": 0,
		"Y": 3,
		"Z": 6,
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		result := strings.Split(scanner.Text(), " ")
		enemy, matchType := result[0], result[1]

		me := myVariant(enemy, matchType)
		score += pointsMatch[matchType]
		score += pointsChoice[me]

	}

	return score
}

func main() {
	var fileName = flag.String("in", "", "file name, required parameter")
	flag.Parse()

	var helpMessage string = fmt.Sprintf("Use \"%s --help\" to see how the program works", os.Args[0])

	*fileName = strings.Trim(*fileName, " ")

	if *fileName == "" {
		log.Fatalln(helpMessage)
	}

	var score int
	score = calculateScoreOfCompleteStrategy(*fileName)

	fmt.Printf("The score is: %d\n", score)
}
