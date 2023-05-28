package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func findCaveDimensions(fileName string) (int, int, int) {
	f, err := os.Open(fileName)
	if err != nil {
		panic("There was an error opening the input file")
	}

	scanner := bufio.NewScanner(f)
	minX, maxX, maxY := 500, 500, 0

	for scanner.Scan() {
		spplitedLine := strings.Split(scanner.Text(), "->")
		for i := 0; i < len(spplitedLine); i++ {
			node := strings.Trim(spplitedLine[i], " ")
			spplitedNode := strings.Split(node, ",")

			x, err := strconv.Atoi(spplitedNode[0])

			if err != nil {
				panic("The file is not the correct format")
			}

			y, err := strconv.Atoi(spplitedNode[1])

			if err != nil {
				panic("The file is not the correct format")
			}

			if minX > x {
				minX = x
			}
			if maxX < x {
				maxX = x
			}
			if maxY < y {
				maxY = y
			}

		}
	}
	return minX, maxX, maxY
}

func fillPaths(fileName string, cM [][]int, offsetX int) {
	f, err := os.Open(fileName)
	if err != nil {
		panic("There was an error opening the input file")
	}
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		spplitedLine := strings.Split(scanner.Text(), "->")

		firstPoint := false
		var prevX, prevY int

		for i := 0; i < len(spplitedLine); i++ {
			node := strings.Trim(spplitedLine[i], " ")
			spplitedNode := strings.Split(node, ",")

			x, err := strconv.Atoi(spplitedNode[0])

			if err != nil {
				panic("The file is not the correct format")
			}

			y, err := strconv.Atoi(spplitedNode[1])

			if err != nil {
				panic("The file is not the correct format")
			}

			cM[y][x-offsetX] = 1
			if firstPoint {
				if x-offsetX == prevX {
					if y < prevY {
						for j := y; j <= prevY; j++ {
							cM[j][x-offsetX] = 1
						}
					} else {
						for j := y; j >= prevY; j-- {
							cM[j][x-offsetX] = 1
						}
					}
				} else {
					if x-offsetX < prevX {
						for j := x - offsetX; j <= prevX; j++ {
							cM[y][j] = 1
						}
					} else {
						for j := x - offsetX; j >= prevX; j-- {
							cM[y][j] = 1
						}
					}
				}
			} else {
				firstPoint = true
			}
			prevX, prevY = x-offsetX, y
		}
	}
}

func displayCaveMap(cM [][]int) {
	for i := 0; i < len(cM); i++ {
		for j := 0; j < len(cM[i]); j++ {
			switch cM[i][j] {
			case 0:
				fmt.Print(".")
			case 1:
				fmt.Print("#")
			case 2:
				fmt.Print("o")
			}
		}
		fmt.Println()
	}
}

type Point struct {
	x, y int
}

func pointInMap(p Point, limitX, limitY int) bool {
	return 0 <= p.x && p.x < limitX && 0 <= p.y && p.y < limitY
}

func addSand(caveMap [][]int, limitX, limitY int, startX int) int {
	var addedSand int
	var sandNotLeftMap bool = true

	for sandNotLeftMap {
		var sandMoving bool = true
		sandPosition := Point{startX, 0}

		for sandMoving {
			if !pointInMap(Point{sandPosition.x, sandPosition.y + 1}, limitX, limitY) {
				sandNotLeftMap = false
				break
			}

			if caveMap[sandPosition.y+1][sandPosition.x] == 0 {
				sandPosition.y++
				continue
			}

			if !pointInMap(Point{sandPosition.x - 1, sandPosition.y + 1}, limitX, limitY) {
				sandNotLeftMap = false
				break
			}

			if caveMap[sandPosition.y+1][sandPosition.x-1] == 0 {
				sandPosition.y++
				sandPosition.x--
				continue
			}

			if !pointInMap(Point{sandPosition.x + 1, sandPosition.y + 1}, limitX, limitY) {
				sandNotLeftMap = false
				break
			}

			if caveMap[sandPosition.y+1][sandPosition.x+1] == 0 {
				sandPosition.y++
				sandPosition.x++
				continue
			}

			sandMoving = false
			caveMap[sandPosition.y][sandPosition.x] = 2
			addedSand++
			// displayCaveMap(caveMap)
            // fmt.Println("Sand added: ", addedSand)
			// fmt.Println("#####################")
		}

	}
	return addedSand
}

func main() {
	var fileName = flag.String("in", "", "file name of paths")

	flag.Parse()

	*fileName = strings.Trim(*fileName, " ")

	if *fileName == "" {
		panic("You need to provide a file name, see use --help for details")
	}

	offsetX, maxX, maxLines := findCaveDimensions(*fileName)
	maxLines++

	sandStartX := 500 - offsetX
	maxColumnsMap := maxX - offsetX + 1

	caveMap := make([][]int, maxLines)
	for i := 0; i < maxLines; i++ {
		for j := 0; j < maxColumnsMap; j++ {
			caveMap[i] = append(caveMap[i], 0)
		}
	}

	fillPaths(*fileName, caveMap, offsetX)
    res := addSand(caveMap, maxColumnsMap, maxLines, sandStartX)

    fmt.Println("Units of sand: ", res)
}
