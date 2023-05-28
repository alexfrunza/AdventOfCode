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

func verifyElfIsTop3(top []Elf, elf Elf) {
	if elf.calories > top[0].calories {
		top[2] = top[1]
		top[1] = top[0]
		top[0] = elf
		return
	}

	if elf.calories > top[1].calories {
		top[2] = top[1]
		top[1] = elf
		return
	}

	if elf.calories > top[2].calories {
		top[2] = elf
		return
	}
}

func sumOfCaloriesElfs(elfs []Elf, number int) int {
	var sum int

	for i := 0; i < number; i++ {
		sum += elfs[i].calories
	}

	return sum
}

func calculateSumTop3Elfs(fileName string) int {
	readFile, err := os.Open(fileName)

	if err != nil {
		log.Fatalln(err)
	}

	defer readFile.Close()

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	var max [3]Elf
	var current Elf
	current.number = 1

	for fileScanner.Scan() {
		if fileScanner.Text() == "" {
			verifyElfIsTop3(max[:], current)

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
	verifyElfIsTop3(max[:], current)

	if current.number < 3 {
		log.Fatalln("Not enough elfs")
	}

	for i := 0; i < 3; i++ {
		fmt.Printf("Elf No.%d: Calories: %d, Number: %d\n", i+1, max[i].calories, max[i].number)
	}

	return sumOfCaloriesElfs(max[:], 3)
}

func main() {
	var sumTop3Elfs int

	var fileName = flag.String("in", "", "File name (Required flag)")
	var helpUsingProgram = fmt.Sprintf("Use \"%s --help\" to see how the program works", os.Args[0])

	flag.Parse()

	*fileName = strings.Trim(*fileName, " ")

	switch *fileName {
	case "":
		log.Fatalln(helpUsingProgram)
	default:
		sumTop3Elfs = calculateSumTop3Elfs(*fileName)
	}

	fmt.Printf("Sum of calories of top 3 elfs is: %d\n", sumTop3Elfs)
}
