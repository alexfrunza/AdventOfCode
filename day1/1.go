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

type Elf struct {
	number   int
	calories int
}

func calculateMaximumCalories(fileName string) (int, int) {
	readFile, err := os.Open(fileName)

	if err != nil {
		log.Fatalln(err)
	}

	defer readFile.Close()

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	var current Elf
	current.number = 1

	var max Elf
	max.number = 0
	max.calories = 0

	for fileScanner.Scan() {
		if fileScanner.Text() == "" {
			if current.calories > max.calories {
				max.number = current.number
				max.calories = current.calories
			}

			current.number++
			current.calories = 0
		} else {
			caloriesFromLine, err := strconv.Atoi(fileScanner.Text())

			if err != nil {
				log.Fatalln("The file isn't good", err)
			}

			current.calories += caloriesFromLine
		}
	}
	if current.calories > max.calories {
		max.number = current.number
		max.calories = current.calories
	}

	return max.calories, max.number
}

func main() {
	var maximumCaloriesOnElf int
	var numberOfMaxElf int

	var fileName = flag.String("in", "", "File name (Required flag)")
	var helpUsingProgram = fmt.Sprintf("Use \"%s --help\" to see how the program works", os.Args[0])

	flag.Parse()

	*fileName = strings.Trim(*fileName, " ")

	switch *fileName {
	case "":
		log.Fatalln(helpUsingProgram)
	default:
		maximumCaloriesOnElf, numberOfMaxElf = calculateMaximumCalories(*fileName)
	}

	fmt.Printf("Maximum calories of an elf: %d\n", maximumCaloriesOnElf)
	fmt.Printf("Number of the Elf: %d\n", numberOfMaxElf)
}
