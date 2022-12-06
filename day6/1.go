package main

import (
	"bufio"
	"fmt"
	"os"
)

func verifyChar(s string) bool {
	for i := 0; i < len(s)-1; i++ {
		for j := i + 1; j < len(s); j++ {
			if s[i] == s[j] {
				return false
			}
		}
	}
	return true
}

func main() {
	// This program take the data from stdin
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	stream := scanner.Text()
	var charsNeeded int

	for i := 0; i < (len(stream) - 3); i++ {
		if verifyChar(stream[i : 4+i]) {
			charsNeeded = i + 4
			break
		}
	}

	fmt.Printf("We need to process %d, before first start-of-packet marker\n", charsNeeded)
}
