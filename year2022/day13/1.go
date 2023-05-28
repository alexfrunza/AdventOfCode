package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type PacketData struct {
	integer int
	list    *ListData
	next    *PacketData
}

type ListData struct {
	length int
	first  *PacketData
	last   *PacketData
}

func emptyList() *ListData {
	var lD ListData
	return &lD
}

func pop(lD *ListData) {
	if lD.length == 0 {
		panic("You can't remove an element from an empty list")
	}
	lD.length--
	if lD.length == 0 {
		lD.first = nil
		lD.last = nil
	}

	lD.first = lD.first.next
}

func push(lD *ListData, pD *PacketData) {
	lD.length++
	if lD.length == 1 {
		lD.first = pD
		lD.last = pD
		return
	}

	lD.last.next = pD
	lD.last = pD
}

func getElement(lD *ListData, nr int) *PacketData {
	if nr < 0 || nr >= lD.length {
		panic("Index invalid")
	}
	pD := lD.first
	for i := 0; i < nr; i++ {
		pD = pD.next
	}
	return pD
}

func displayPacketData(pD *PacketData) {
	if pD.list == nil {
		fmt.Print(pD.integer)
	} else {
		switch pD.list.length {
		case 0:
			fmt.Print("[")
			fmt.Print("]")
		case 1:
			fmt.Print("[")
			displayPacketData(pD.list.first)
			fmt.Print("]")
		default:
			fmt.Print("[")
			for e := pD.list.first; e.next != nil; e = e.next {
				displayPacketData(e)
				fmt.Print(",")
			}
			displayPacketData(pD.list.last)
			fmt.Print("]")
		}

	}
}

// Returns the index of the closing bracket of the first bracket
func getIndexOfEndingBracket(text string) int {
	// Skip the first character because is excepted to be '['
	// The paranthesis is excepted to be correct
	var bracketCounter int = 1
	var i int
	for i = 1; i < len(text) && bracketCounter > 0; i++ {
		switch text[i] {
		case '[':
			bracketCounter++
		case ']':
			bracketCounter--
		}
	}
	return i - 1
}

func parseThePacket(text string) *PacketData {
	var res PacketData

	if text[0] != '[' {
		i, err := strconv.Atoi(text)

		if err != nil {
			panic("The file is not the correct format")
		}

		res.integer = i
	} else {
		res.list = emptyList()
		stopPosition := getIndexOfEndingBracket(text)

		for i := 1; i < stopPosition; {
			if text[i] == ',' {
				i++
				continue
			}

			var newElement *PacketData
			var endPosition int
			var nextValueI int

			if '0' <= text[i] && text[i] <= '9' {
				indexOfNextComma := strings.Index(text[i:], ",")

				switch indexOfNextComma {
				case -1:
					endPosition = len(text) - 1
					nextValueI = len(text)
				default:
					endPosition = i + indexOfNextComma
					nextValueI = i + indexOfNextComma
				}
			} else {
				endPosition = i + getIndexOfEndingBracket(text[i:]) + 1
				nextValueI = endPosition + 1
			}
			newElement = parseThePacket(text[i:endPosition])
			push(res.list, newElement)
			i = nextValueI
		}
	}

	return &res
}

// The function return:
// 1 if the pairs are in order
// 0 if the pairs are equals
// -1 if the pairs are not in order
func compareThePackets(firstPacket *PacketData, secondPacket *PacketData) int {
	if firstPacket.list == nil && secondPacket.list == nil {
		if firstPacket.integer < secondPacket.integer {
			return 1
		} else if firstPacket.integer == secondPacket.integer {
			return 0
		} else {
			return -1
		}
	} else if firstPacket.list != nil && secondPacket.list != nil {
		i := 0
		for i = 0; i < firstPacket.list.length && i < secondPacket.list.length; i++ {
			switch compareThePackets(getElement(firstPacket.list, i), getElement(secondPacket.list, i)) {
			case -1:
				return -1
			case 1:
				return 1
			}
		}

		if i == secondPacket.list.length && i < firstPacket.list.length {
			return -1
		}
		if firstPacket.list.length == secondPacket.list.length {
			return 0
		}
	} else if firstPacket.list == nil {
		var aux PacketData

		var aux2 PacketData
		aux2.integer = firstPacket.integer

		aux.list = emptyList()
		push(aux.list, &aux2)

		return compareThePackets(&aux, secondPacket)
	} else if secondPacket.list == nil {
		var aux PacketData

		var aux2 PacketData
		aux2.integer = secondPacket.integer

		aux.list = emptyList()
		push(aux.list, &aux2)

		return compareThePackets(firstPacket, &aux)
	}
	return 1
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	counter := 1
	res := 0

	for {
		// First packet of pair
		scanner.Scan()
		var firstPacket PacketData
		firstPacket = *parseThePacket(scanner.Text())
		// displayPacketData(&firstPacket)
		// fmt.Println()

		// Second packet of pair
		scanner.Scan()
		var secondPacket PacketData
		secondPacket = *parseThePacket(scanner.Text())
		// displayPacketData(&secondPacket)
		// fmt.Println()

		r := compareThePackets(&firstPacket, &secondPacket)
		if r == 1 {
			res += counter
		}
		counter++

		if scanner.Scan() == false {
			break
		}
	}

	fmt.Println("The sum of the indices of the pairs is: ", res)
}
