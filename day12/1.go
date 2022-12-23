package main

import (
	"bufio"
	"fmt"
	"os"
)

type Point struct {
	x, y              int
	distanceFromStart int
	next              *Point
}

type Tail struct {
	length      int
	first, last *Point
}

func top(t *Tail) *Point {
	return t.first
}

func pop(t *Tail) {
	if t.length == 0 {
		panic("You can't pop from an empty tail")
	}

	t.length--
	if t.length == 0 {
		t.first = nil
		t.last = nil
		return
	}

	t.first = t.first.next
}

func add(t *Tail, p Point) {
	if t.length == 0 {
		t.first = &p
		t.last = &p
		t.length = 1
		return
	}

	t.length++
	t.last.next = &p
	t.last = &p
}

type PointMap struct {
	visited           bool
	height            byte
	distanceFromStart int
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	var heightmap [][]PointMap
	var line int
	var startPosition Point
	var target Point

	for scanner.Scan() {
		heightmap = append(heightmap, []PointMap{})
		for i := 0; i < len(scanner.Text()); i++ {
			switch scanner.Text()[i] {
			case 'S':
				heightmap[line] = append(heightmap[line], PointMap{false, 'a', 0})
				startPosition = Point{i, line, 0, nil}
			case 'E':
				heightmap[line] = append(heightmap[line], PointMap{false, 'z', 0})
				target = Point{i, line, 0, nil}
			default:
				heightmap[line] = append(heightmap[line], PointMap{false, scanner.Text()[i], 0})
			}
		}
		line++
	}

	var t Tail
	add(&t, startPosition)
	for t.length > 0 {
		xM := []int{1, 0, -1, 0}
		yM := []int{0, 1, 0, -1}
		cP := t.first
		pop(&t)

		if heightmap[cP.y][cP.x].visited == false {
			heightmap[cP.y][cP.x].distanceFromStart = cP.distanceFromStart
			heightmap[cP.y][cP.x].visited = true
		} else if cP.distanceFromStart < heightmap[cP.y][cP.x].distanceFromStart {
			heightmap[cP.y][cP.x].distanceFromStart = cP.distanceFromStart
		} else {
			continue
		}

		if target.x == cP.x && target.y == cP.y {
			continue
		}

		for i := 0; i < 4; i++ {
			// If next direction exit from heightmap do nothing
			if !(0 <= cP.x+xM[i] && cP.x+xM[i] < len(heightmap[0]) &&
				0 <= cP.y+yM[i] && cP.y+yM[i] < len(heightmap)) {
				continue
			}

			if int(heightmap[cP.y][cP.x].height)-int(heightmap[cP.y+yM[i]][cP.x+xM[i]].height) < -1 {
				continue
			}

			add(&t, Point{cP.x + xM[i], cP.y + yM[i], cP.distanceFromStart + 1, nil})
		}
	}

	fmt.Println("The shortest distance from start to the target point is: ", heightmap[target.y][target.x].distanceFromStart)
}
