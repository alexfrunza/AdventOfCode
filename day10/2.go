package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func getMiddleSprite(register int) int {
	return (register - 1) % 40
}

func incrementCycle(cycleNumber *int, prevValueRegister int, screen [][]string) {
	if *cycleNumber > 241 {
		return
	}

	x, y := (*cycleNumber-1)%40, (*cycleNumber-1)/40
	xSprite := getMiddleSprite(prevValueRegister + 1)

	ok := false
	for i := -1; i <= 1; i++ {
		if xSprite+i == x {
			ok = true
		}
	}
	if ok {
		screen[y][x] = "#"
	} else {
		screen[y][x] = "."
	}

	*cycleNumber += 1
}

func main() {
	var registerX int = 1
	var cycleCounter int = 1
	var screen [][]string

	for i := 0; i < 6; i++ {
		screen = append(screen, make([]string, 40))
	}

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() && cycleCounter <= 240 {
		prevValueRegister := registerX

		splittedLine := strings.Split(scanner.Text(), " ")

		if len(splittedLine) == 1 {
			incrementCycle(&cycleCounter, prevValueRegister, screen)
		} else {
			value, err := strconv.Atoi(splittedLine[1])

			if err != nil {
				log.Fatalln("The file isn't the corect format", err)
			}

			incrementCycle(&cycleCounter, prevValueRegister, screen)
			incrementCycle(&cycleCounter, prevValueRegister, screen)
			registerX += value
		}

	}

	fmt.Println("The screen looks like this after execution:")
	for i := 0; i < 6; i++ {
		for j := 0; j < 40; j++ {
			fmt.Print(screen[i][j])
		}
		fmt.Println()
	}
}
