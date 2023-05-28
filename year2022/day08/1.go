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

func isVisible(tP treePosition) bool {
	var isVisible bool

	// upside
	isVisible = true
	for i := tP.x - 1; i >= 0; i-- {
		if tP.treesHeight[tP.x][tP.y] <= tP.treesHeight[i][tP.y] {
			isVisible = false
			break
		}
	}

	if isVisible == true {
		return true
	}

	// downside
	isVisible = true
	for i := tP.x + 1; i < len(tP.treesHeight); i++ {
		if tP.treesHeight[tP.x][tP.y] <= tP.treesHeight[i][tP.y] {
			isVisible = false
			break
		}
	}

	if isVisible == true {
		return true
	}

	// rightside
	isVisible = true
	for i := tP.y + 1; i < len(tP.treesHeight[tP.x]); i++ {
		if tP.treesHeight[tP.x][tP.y] <= tP.treesHeight[tP.x][i] {
			isVisible = false
			break
		}
	}

	if isVisible == true {
		return true
	}

	// left side
	isVisible = true
	for i := tP.y - 1; i >= 0; i-- {
		if tP.treesHeight[tP.x][tP.y] <= tP.treesHeight[tP.x][i] {
			isVisible = false
			break
		}
	}

	if isVisible == true {
		return true
	}

	return false
}

func calculateVisibleTrees(tH [][]uint8) int {
	var result int

	// Number of trees from the edge
	result = 2*(len(tH)+len(tH[0])) - 4

	for i := 1; i < len(tH) - 1; i++ {
		for j := 1; j < len(tH[i]) - 1; j++ {
			if isVisible(treePosition{i, j, len(tH), len(tH[i]), tH}) == true {
				result++
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

	var visibleTrees int
	visibleTrees = calculateVisibleTrees(treesHeight)
	fmt.Println("The number of visible trees is: ", visibleTrees)
}
