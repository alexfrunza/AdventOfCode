package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type node struct {
	x, y int
}

func showVisitedNodes(nodesList []node) {
	for i := 0; i < len(nodesList); i++ {
		fmt.Println("X:", nodesList[i].x, "Y:", nodesList[i].y)
	}
}

func isVisited(position node, nodesList []node) bool {
	for i := 0; i < len(nodesList); i++ {
		if nodesList[i] == position {
			return true
		}
	}
	return false
}

func addCoordinates(dest *node, adder *node) {
	dest.x += adder.x
	dest.y += adder.y
}

func isOnDiagonal(head *node, tail *node) bool {
	if head.x-1 == tail.x && head.y+1 == tail.y {
		// Case
		// . . .
		// . H .
		// T . .
		return true
	} else if head.x+1 == tail.x && head.y+1 == tail.y {
		// Case
		// . . .
		// . H .
		// T . .
		return true
	} else if head.x-1 == tail.x && head.y-1 == tail.y {
		// Case
		// T . .
		// . H .
		// . . .
		return true
	} else if head.x+1 == tail.x && head.y-1 == tail.y {
		// Case
		// . . T
		// . H .
		// . . .
		return true
	}

	return false
}

func isOnProximity(head *node, tail *node) bool {
	if isOnDiagonal(head, tail) {
		return true
	}
	if head.x == tail.x && head.y-1 == tail.y {
		// Case
		// . T .
		// . H .
		// . . .
		return true
	} else if head.x+1 == tail.x && head.y == tail.y {
		// Case
		// . . .
		// . H T
		// . . .
		return true
	} else if head.x == tail.x && head.y+1 == tail.y {
		// Case
		// . . .
		// . H .
		// . T .
		return true
	} else if head.x-1 == tail.x && head.y == tail.y {
		// Case
		// . . .
		// T H .
		// . . .
		return true
	} else if *head == *tail {
		return true
	}
	return false
}

func moveHead(headPosition, tailPosition *node, direction string, movements int, visitedNodes *[]node) {
	var modifier node
	switch direction {
	case "R":
		modifier.x = 1
	case "D":
		modifier.y = 1
	case "L":
		modifier.x = -1
	case "U":
		modifier.y = -1
	}

	for i := 0; i < movements; i++ {
		if *headPosition == *tailPosition {
			addCoordinates(headPosition, &modifier)
			continue
		}

		wasOnDiagonal := isOnDiagonal(headPosition, tailPosition)

		addCoordinates(headPosition, &modifier)

		if isOnProximity(headPosition, tailPosition) {
			continue
		} else {
			if !wasOnDiagonal {
				addCoordinates(tailPosition, &modifier)
			} else {
				switch direction {
				case "R":
					tailPosition.x = headPosition.x - 1
					tailPosition.y = headPosition.y
				case "D":
					tailPosition.x = headPosition.x
					tailPosition.y = headPosition.y - 1
				case "L":
					tailPosition.x = headPosition.x + 1
					tailPosition.y = headPosition.y
				case "U":
					tailPosition.x = headPosition.x
					tailPosition.y = headPosition.y + 1
				}

			}
		}

		if !isVisited(*tailPosition, *visitedNodes) {
			*visitedNodes = append(*visitedNodes, *tailPosition)
		}
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	var visitedNodes []node
	visitedNodes = append(visitedNodes, node{0, 0})

	headPosition := node{0, 0}
	tailPosition := node{0, 0}

	for scanner.Scan() {
		splittedText := strings.Split(scanner.Text(), " ")
		direction := splittedText[0]

		movements, err := strconv.Atoi(splittedText[1])

		if err != nil {
			log.Fatalln("The file isn't the good format", err)
		}

		moveHead(&headPosition, &tailPosition, direction, movements, &visitedNodes)
	}

	/// showVisitedNodes(visitedNodes)
	fmt.Println("The number of positions the tail visited at least once is: ", len(visitedNodes))
}
