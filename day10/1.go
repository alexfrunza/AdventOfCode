package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	var sumSignalStrength int
	var registerX int = 1
	var cycleCounter int = 1
	var targetCycle int = 20

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() && targetCycle <= 220 {
		prevValueRegister := registerX

		splittedLine := strings.Split(scanner.Text(), " ")

		if len(splittedLine) == 1 {
			cycleCounter += 1
		} else {
			value, err := strconv.Atoi(splittedLine[1])

			if err != nil {
				log.Fatalln("The file isn't the corect format", err)
			}

			cycleCounter += 2
			registerX += value
		}

		if cycleCounter == targetCycle {
			sumSignalStrength += targetCycle * registerX
			targetCycle += 40
		} else if cycleCounter > targetCycle {
			sumSignalStrength += targetCycle * prevValueRegister
			targetCycle += 40
		}
	}

	fmt.Println("The sum of six signal strength is: ", sumSignalStrength)
}
