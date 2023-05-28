package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

func calculatePriority(x byte) int {
	if 'a' <= x && x <= 'z' {
		return int(x) - int('a') + 1
	}

	return int(x) - int('A') + 1 + 26
}

func fillSet(elements []byte, set map[int]bool) {
	for i := 0; i < len(elements); i++ {
		set[calculatePriority(elements[i])] = true
	}
}

func calculateSumOfPriorities(fileName string) int {

	file, err := os.Open(fileName)

	if err != nil {
		log.Fatalln("The file with rucksacks can't be open", err)
	}

	var sumOfPriorities int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		firstLine := scanner.Text()
		scanner.Scan()
		secondLine := scanner.Text()
		scanner.Scan()
		thirdLine := scanner.Text()

		firstSet := make(map[int]bool, 52)
		secondSet := make(map[int]bool, 52)
		thirdSet := make(map[int]bool, 52)

		fillSet([]byte(firstLine), firstSet)
		fillSet([]byte(secondLine), secondSet)
		fillSet([]byte(thirdLine), thirdSet)

		for i := 1; i <= 52; i++ {
			if firstSet[i] == true && secondSet[i] == true && thirdSet[i] == true {
				sumOfPriorities += i
				break
			}
		}
	}

	return sumOfPriorities
}

func main() {
	var fileName = flag.String("in", "", "file name, required flag")

	flag.Parse()

	var helpMessage string = fmt.Sprintf("Use \"%s --help\" to see how to use the program.", os.Args[0])

	*fileName = strings.Trim(*fileName, " ")

	if *fileName == "" {
		log.Fatalln(helpMessage)
	}

	var sumOfPriorities int

	sumOfPriorities = calculateSumOfPriorities(*fileName)

	fmt.Printf("Sum of priorities is: %d\n", sumOfPriorities)
}
