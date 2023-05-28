package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Node struct {
	next *Node
	info int
}

type Tail struct {
	length int
	first  *Node
	last   *Node
}

func addItem(t *Tail, info int) {
	var newNode Node
	newNode.info = info
	t.length++

	if t.length == 1 {
		t.first = &newNode
		t.last = &newNode
		return
	}

	t.last.next = &newNode
	t.last = &newNode
}

func top(t *Tail) *Node {
	return t.first
}

func removeItem(t *Tail) *Node {
	if t.length == 0 {
		log.Fatalln("You can't remove an element from an empty tail")
	}

	temp := t.first
	if t.length == 1 {
		t.length = 0
		t.first = nil
		t.last = nil
		return temp
	}

	t.length--
	t.first = t.first.next
	return temp
}

func inspectItem(m *Monkey) {
	var op1 int
	var op2 int
	var err error

	if m.opS1 == "old" {
		op1 = m.items.first.info
	} else {
		op1, err = strconv.Atoi(m.opS1)

		if err != nil {
			log.Fatalln("The file is not the corect format", err)
		}
	}

	if m.opS2 == "old" {
		op2 = m.items.first.info
	} else {
		op2, err = strconv.Atoi(m.opS2)

		if err != nil {
			log.Fatalln("The file is not the corect format", err)
		}

	}

	switch m.op {
	case "*":
		m.items.first.info = (op1 * op2) / 3
	case "+":
		m.items.first.info = (op1 + op2) / 3
	}
}

type Monkey struct {
	number         int
	items          Tail
	testDivisor    int
	inspectedItems int
	falseNumber    int
	trueNumber     int
	op             string
	opS1           string
	opS2           string
}

func trowItem(m *Monkey, monkeys []Monkey) {
	if m.items.first.info%m.testDivisor == 0 {
		addItem(&monkeys[m.trueNumber].items, m.items.first.info)
		removeItem(&m.items)
	} else {
		addItem(&monkeys[m.falseNumber].items, m.items.first.info)
		removeItem(&m.items)
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	var monkeys []Monkey

	for scanner.Scan() {
		var spplitedText []string
		var m Monkey

		// Take the monkey number
		spplitedText = strings.Split(scanner.Text(), " ")

		if len(spplitedText) < 1 || spplitedText[0] != "Monkey" {
			continue
		}

		mN := spplitedText[len(spplitedText)-1]
		monkeyNumber, err := strconv.Atoi(mN[:len(mN)-1])

		if err != nil {
			log.Fatalln("The file is not the corect format", err)
		}
		m.number = monkeyNumber

		// Take the starting items
		scanner.Scan()
		spplitedText = strings.Split(strings.ReplaceAll(scanner.Text(), ",", " "), " ")
		for i := 4; i < len(spplitedText); i += 2 {
			objInfo, err := strconv.Atoi(spplitedText[i])

			if err != nil {
				log.Fatalln("The file is not the corect format", err)
			}

			addItem(&m.items, objInfo)
		}

		// Inspect operation of the object
		scanner.Scan()
		spplitedText = strings.Split(scanner.Text(), " ")
		m.opS1 = spplitedText[len(spplitedText)-3]
		m.op = spplitedText[len(spplitedText)-2]
		m.opS2 = spplitedText[len(spplitedText)-1]

		// Test function of the monkey
		scanner.Scan()
		spplitedText = strings.Split(scanner.Text(), " ")

		divisor, err := strconv.Atoi(spplitedText[len(spplitedText)-1])

		if err != nil {
			log.Fatalln("The file is not the corect format", err)
		}
		m.testDivisor = divisor

		// Throw item to another monkey
		scanner.Scan()
		spplitedText = strings.Split(scanner.Text(), " ")

		trueNumber, err := strconv.Atoi(spplitedText[len(spplitedText)-1])
		if err != nil {
			log.Fatalln("The file is not the corect format", err)
		}

		scanner.Scan()
		spplitedText = strings.Split(scanner.Text(), " ")

		falseNumber, err := strconv.Atoi(spplitedText[len(spplitedText)-1])
		if err != nil {
			log.Fatalln("The file is not the corect format", err)
		}

		m.trueNumber = trueNumber
		m.falseNumber = falseNumber

		monkeys = append(monkeys, m)
	}

	for k := 0; k < 20; k++ {
		for i := 0; i < len(monkeys); i++ {
			cM := &monkeys[i]
			for cM.items.length > 0 {
				cM.inspectedItems++
				inspectItem(cM)
				trowItem(cM, monkeys)
			}
		}
	}

	var max1 int
	var max2 int
	for i := 0; i < len(monkeys); i++ {
		if monkeys[i].inspectedItems > max1 {
			max2 = max1
			max1 = monkeys[i].inspectedItems
		} else if monkeys[i].inspectedItems > max2 {
			max2 = monkeys[i].inspectedItems
		}
	}

	fmt.Println("The level of monkey business is: ", max1*max2)
}
