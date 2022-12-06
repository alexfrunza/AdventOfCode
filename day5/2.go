package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

const (
	FIRST_STEP = iota
	CRATES_STEP
	PROCEDURES_STEP
)

type Node struct {
	info byte
	next *Node
}

type Stack struct {
	length int
	first  *Node
	last   *Node
}

func addElementAtStart(s *Stack, e byte) {
	if s.first == nil {
		s.length = 1
		var n Node
		n.info = e
		s.first = &n
		s.last = &n
		return
	}

	var n Node
	s.length++
	n.info = e
	s.first.next = &n
	s.first = &n
}

func addElement(s *Stack, e byte) {
	if s.last == nil {
		s.length = 1
		var n Node
		n.info = e
		s.first = &n
		s.last = &n
		return
	}

	var n Node
	s.length++
	n.info = e
	n.next = s.last
	s.last = &n
}

func popElement(s *Stack) byte {
	if s.last == nil {
		log.Fatalln("The stack was empty")
	}

	s.length--
	var n *Node = s.last
	s.last = n.next
	return n.info
}

func moveCrates(number int, fromStack *Stack, toStack *Stack) {
	if number > 0 {
		el := popElement(fromStack)
		moveCrates(number-1, fromStack, toStack)
		addElement(toStack, el)
	}
}

func findCrates(fileName string) string {
	file, err := os.Open(fileName)

	if err != nil {
		log.Fatalln("There was an error open the file with crates", err)
	}

	scanner := bufio.NewScanner(file)

	var stacks []Stack
	step := FIRST_STEP

	for scanner.Scan() {
		if step == FIRST_STEP {
			for crateInfoIndex := 1; crateInfoIndex < len(scanner.Text()); crateInfoIndex += 4 {
				var s Stack
				if crateInfo := scanner.Text()[crateInfoIndex]; crateInfo != ' ' {
					addElementAtStart(&s, crateInfo)
				}
				stacks = append(stacks, s)
			}
			step = CRATES_STEP
		} else if step == CRATES_STEP {
			if scanner.Text()[1] == '1' {
				scanner.Scan()
				step = PROCEDURES_STEP
				continue
			}
			for crateInfoIndex := 1; crateInfoIndex < len(scanner.Text()); crateInfoIndex += 4 {
				if crateInfo := scanner.Text()[crateInfoIndex]; crateInfo != ' ' {
					addElementAtStart(&stacks[(crateInfoIndex-1)/4], crateInfo)
				}
			}
		} else if step == PROCEDURES_STEP {
			step := scanner.Text()

			re := regexp.MustCompile(`[0-9]+`)
			splitStep := re.FindAllString(step, -1)

			cratesNumber, err := strconv.Atoi(splitStep[0])

			if err != nil {
				log.Fatalln("File bad", err)
			}

			fromStack, err := strconv.Atoi(splitStep[1])

			if err != nil {
				log.Fatalln("File bad", err)
			}

			toStack, err := strconv.Atoi(splitStep[2])

			if err != nil {
				log.Fatalln("File bad", err)
			}

			moveCrates(cratesNumber, &stacks[fromStack-1], &stacks[toStack-1])
		}
	}

	var crates []byte
	for i := 0; i < len(stacks); i++ {
		el := popElement(&stacks[i])
		crates = append(crates, el)
	}

	return string(crates)
}

func main() {
	var fileName = flag.String("in", "", "file name, required parameter")
	flag.Parse()
	var helpMessage string = fmt.Sprintf("Use \"%s --help\" to see how to use the program", os.Args[0])

	*fileName = strings.Trim(*fileName, " ")

	if *fileName == "" {
		log.Fatalln(helpMessage)
	}

	var crates string
	crates = findCrates(*fileName)
	fmt.Println("The crates which are on top are: ", crates)
}
