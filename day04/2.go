package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func splitInPairs(row string) (string, string) {
	splitedRow := strings.Split(row, ",")

	if len(splitedRow) != 2 {
		log.Fatalln("The input is not as expected")
	}

	return splitedRow[0], splitedRow[1]
}

func splitPairs(pair string) (int, int) {
	splitedPair := strings.Split(pair, "-")

	if len(splitedPair) != 2 {
		log.Fatalln("The input is not as expected")
	}

	firstElement, err := strconv.Atoi(splitedPair[0])

	if err != nil {
		log.Fatalln("The input is not as expected")
	}

	secondElement, err := strconv.Atoi(splitedPair[1])

	if err != nil {
		log.Fatalln("The input is not as expected")
	}

	return firstElement, secondElement
}

func calculatePairsOverlappingAtAll(fileName string) int {

	file, err := os.Open(fileName)

	if err != nil {
		log.Fatalln("Failed to open the file with pairs", err)
	}

	var pairsOverlapping int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		p1, p2 := splitInPairs(scanner.Text())

		firstP1, secondP1 := splitPairs(p1)
		firstP2, secondP2 := splitPairs(p2)

		if firstP1 <= firstP2 && firstP2 <= secondP1 {
			pairsOverlapping++
		} else if firstP1 <= secondP2 && secondP2 <= secondP1 {
			pairsOverlapping++
		} else if firstP2 <= firstP1 && firstP1 <= secondP2 {
			pairsOverlapping++
		} else if firstP2 <= secondP1 && secondP1 <= secondP2 {
			pairsOverlapping++
		}
	}

	return pairsOverlapping
}

func main() {
	var fileName = flag.String("in", "", "file name, required flag")

	flag.Parse()

	var helpMessage string = fmt.Sprintf("Use \"%s --help\" to see how to use the program", os.Args[0])

	*fileName = strings.Trim(*fileName, " ")

	if *fileName == "" {
		log.Fatalln(helpMessage)
	}

	var pairsOverlapping int
	pairsOverlapping = calculatePairsOverlappingAtAll(*fileName)

	fmt.Println("Number of pairs overlapping at all is: ", pairsOverlapping)
}
