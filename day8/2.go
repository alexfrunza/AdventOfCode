package main

import (
	"bufio"
	"fmt"
	"os"
)

// For debugging
func displayMatrix(m [][]uint8) {
	for i := 0; i < len(m); i++ {
		for j := 0; j < len(m[i]); j++ {
			fmt.Printf("%d ", m[i][j])
		}
		fmt.Printf("\n")
	}
}

type treePosition struct {
	x, y        int
	xMax, yMax  int
	treesHeight [][]uint8
}

func scenicScore(tP treePosition) int {
	var score int = 1
	var seenTrees int

	// upside
	seenTrees = 0
	for i := tP.x - 1; i >= 0; i-- {
		seenTrees++
		if tP.treesHeight[tP.x][tP.y] <= tP.treesHeight[i][tP.y] {
			break
		}
	}
	score *= seenTrees

	// downside
	seenTrees = 0
	for i := tP.x + 1; i < len(tP.treesHeight); i++ {
        seenTrees++
		if tP.treesHeight[tP.x][tP.y] <= tP.treesHeight[i][tP.y] {
			break
		}
	}
	score *= seenTrees

	// rightside
	seenTrees = 0
	for i := tP.y + 1; i < len(tP.treesHeight[tP.x]); i++ {
		seenTrees++
		if tP.treesHeight[tP.x][tP.y] <= tP.treesHeight[tP.x][i] {
			break
		}
	}
	score *= seenTrees

	// left side
	seenTrees = 0
	for i := tP.y - 1; i >= 0; i-- {
		seenTrees++
		if tP.treesHeight[tP.x][tP.y] <= tP.treesHeight[tP.x][i] {
			break
		}
	}
	score *= seenTrees

	return score
}

func calculateMaxScenicScore(tH [][]uint8) int {
	var result int

	for i := 1; i < len(tH)-1; i++ {
		for j := 1; j < len(tH[i])-1; j++ {
			if s := scenicScore(treePosition{i, j, len(tH), len(tH[i]), tH}); s > result {
				result = s
			}
		}
	}

	return result
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	var treesHeight [][]uint8
	var linesNumber int

	for ; scanner.Scan(); linesNumber++ {
		line := scanner.Text()
		treesHeight = append(treesHeight, []uint8{})
		for i := 0; i < len(line); i++ {
			treesHeight[linesNumber] = append(treesHeight[linesNumber], line[i]-'0')
		}
	}

	var maxScenicScore int
	maxScenicScore = calculateMaxScenicScore(treesHeight)
	fmt.Println("The highest scenic score is: ", maxScenicScore)
}
