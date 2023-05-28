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

func calculateSumOfPriorities(fileName string) int {

	file, err := os.Open(fileName)

	if err != nil {
		log.Fatalln("The file with rucksacks can't be open", err)
	}

	var sumOfPriorities int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := []byte(scanner.Text())
		lengthLine := len(line)
		charSet := make(map[int]bool, 52)

		for i := 0; i < lengthLine/2; i++ {
			charSet[calculatePriority(line[i])] = true
		}
		for i := lengthLine / 2; i < lengthLine; i++ {
			if charSet[calculatePriority(line[i])] {
				sumOfPriorities += calculatePriority(line[i])
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
