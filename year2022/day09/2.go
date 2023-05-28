package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
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

func isOnArrow(target, follower *node) bool {
	if target.x == follower.x && target.y-1 == follower.y {
		// Case
		// . T .
		// . H .
		// . . .
		return true
	} else if target.x+1 == follower.x && target.y == follower.y {
		// Case
		// . . .
		// . H T
		// . . .
		return true
	} else if target.x == follower.x && target.y+1 == follower.y {
		// Case
		// . . .
		// . H .
		// . T .
		return true
	} else if target.x-1 == follower.x && target.y == follower.y {
		// Case
		// . . .
		// T H .
		// . . .
		return true
	}
	return false
}

func isOnProximity(head *node, tail *node) bool {
	if isOnDiagonal(head, tail) {
		return true
	}
	if isOnArrow(head, tail) {
		return true
	}
	if head.x == tail.x && head.y == tail.y {
		return true
	}
	return false
}

func distance(point1, point2 *node) float64 {
	x1 := float64(point1.x)
	x2 := float64(point2.x)
	y1 := float64(point1.y)
	y2 := float64(point2.y)
	return math.Sqrt((x1-x2)*(x1-x2) + (y1-y2)*(y1-y2))

}

func calcMinDistPoint(target, follower *node) (node, float64) {
	var minDistance float64 = 100
	var res node

	for i := -1; i <= 1; i++ {
		if i == 0 {
			continue
		}
		temp := *target
		temp.x += i
		if distance(&temp, follower) < minDistance {
			minDistance = distance(&temp, follower)
			res = temp
		}

		temp = *target
		temp.y += i
		if distance(&temp, follower) < minDistance {
			minDistance = distance(&temp, follower)
			res = temp
		}
	}

	return res, minDistance
}

func getArrowDirection(target, follower node) string {
	if target.x+1 == follower.x {
		return "R"
	}
	if target.x-1 == follower.x {
		return "L"
	}
	if target.y+1 == follower.y {
		return "D"
	}
	return "U"
}

func moveOnArrowDirection(target, follower *node, direction string) {
	switch direction {
	case "U":
		follower.y = target.y - 1
		follower.x = target.x
	case "D":
		follower.y = target.y + 1
		follower.x = target.x
	case "L":
		follower.y = target.y
		follower.x = target.x - 1
	case "R":
		follower.y = target.y
		follower.x = target.x + 1
	}
}

func moveHead(rope []node, direction string, visitedNodes *[]node) {
	ok := false

	prevLocation := rope[0]
	for i := 0; i < 9; i++ {
		wasOnArrow := isOnArrow(&prevLocation, &rope[i+1])
		wasOnDiagonal := isOnDiagonal(&prevLocation, &rope[i+1])

		if i == 0 {
			switch direction {
			case "U":
				rope[0].y--
			case "D":
				rope[0].y++
			case "L":
				rope[0].x--
			case "R":
				rope[0].x++
			}
		}

		if isOnProximity(&rope[i], &rope[i+1]) {
			break
		}

		newPrevLocation := rope[i+1]
		if wasOnArrow && ok {
			arrowDirection := getArrowDirection(prevLocation, rope[i+1])
			moveOnArrowDirection(&rope[i], &rope[i+1], arrowDirection)
			prevLocation = newPrevLocation
			continue
		}

		temp, distance := calcMinDistPoint(&rope[i], &rope[i+1])
		if distance == 1 {
			rope[i+1] = temp
		} else {
			rope[i+1] = prevLocation
		}
		if wasOnDiagonal {
			ok = true
		}
		prevLocation = newPrevLocation
	}

	if !isVisited(rope[9], *visitedNodes) {
		*visitedNodes = append(*visitedNodes, rope[9])
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	var visitedNodes []node
	visitedNodes = append(visitedNodes, node{0, 0})

	rope := make([]node, 10)

	for scanner.Scan() {
		splittedText := strings.Split(scanner.Text(), " ")
		direction := splittedText[0]

		movements, err := strconv.Atoi(splittedText[1])

		if err != nil {
			log.Fatalln("The file isn't the good format", err)
		}

		for j := 0; j < movements; j++ {
			moveHead(rope, direction, &visitedNodes)
		}
	}

	// showVisitedNodes(visitedNodes)
	fmt.Println("The number of positions the tail visited at least once is: ", len(visitedNodes))
}
