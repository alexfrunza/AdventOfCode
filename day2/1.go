package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

func calculateScoreOfStrategy(fileName string) int {
	file, err := os.Open(fileName)

	if err != nil {
		log.Fatalln("The file with strategy can't be open =(", err)
	}

	var score int
	var winningVariant map[string]string
	var pointsVariant map[string]int
	var equalVariant map[string]string

	winningVariant = map[string]string{
		"X": "C",
		"Y": "A",
		"Z": "B",
	}
	pointsVariant = map[string]int{
		"X": 1,
		"Y": 2,
		"Z": 3,
	}
	equalVariant = map[string]string{
		"X": "A",
		"Y": "B",
		"Z": "C",
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		result := strings.Split(scanner.Text(), " ")
		enemy, me := result[0], result[1]

		score += pointsVariant[me]
		if equalVariant[me] == enemy {
			score += 3
		} else if winningVariant[me] == enemy {
			score += 6
		}
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
	score = calculateScoreOfStrategy(*fileName)

	fmt.Printf("The score is: %d\n", score)
}
