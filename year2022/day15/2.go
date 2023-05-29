package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

type point struct {
	x int
	y int
}

type pair struct {
	sP point
	bP point
}

func getData(fileName string) []pair {
	pairs := []pair{}

	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	pat := regexp.MustCompile(`([-]?\d+), y=([-]?\d+)`)
	for scanner.Scan() {
		if scanner.Text() != "" {
			matches := pat.FindAllStringSubmatch(scanner.Text(), -1)

			var x1, x2, y1, y2 int

			x1, err = strconv.Atoi(matches[0][1])
			if err != nil {
				log.Fatal(err)
			}

			y1, err = strconv.Atoi(matches[0][2])
			if err != nil {
				log.Fatal(err)
			}

			x2, err = strconv.Atoi(matches[1][1])
			if err != nil {
				log.Fatal(err)
			}

			y2, err = strconv.Atoi(matches[1][2])
			if err != nil {
				log.Fatal(err)
			}

			pairs = append(pairs, pair{
				sP: point{
					x: x1,
					y: y1,
				},
				bP: point{
					x: x2,
					y: y2,
				},
			})
		}
	}

	return pairs
}

func getLimitingPoints(pairs []pair) [2]point {
	var limitingPoints [2]point

	for i, p := range pairs {
		if i == 0 {
			if p.sP.x < p.bP.x {
				limitingPoints[0].x = p.sP.x
				limitingPoints[1].x = p.bP.x
			} else {
				limitingPoints[0].x = p.bP.x
				limitingPoints[1].x = p.sP.x
			}

			if p.sP.y < p.bP.y {
				limitingPoints[0].y = p.sP.y
				limitingPoints[1].y = p.bP.y
			} else {
				limitingPoints[0].y = p.bP.y
				limitingPoints[1].y = p.sP.y
			}
		} else {
			if limitingPoints[0].x > p.sP.x {
				limitingPoints[0].x = p.sP.x
			}

			if limitingPoints[1].x < p.sP.x {
				limitingPoints[1].x = p.sP.x
			}

			if limitingPoints[0].x > p.bP.x {
				limitingPoints[0].x = p.bP.x
			}

			if limitingPoints[1].x < p.bP.x {
				limitingPoints[1].x = p.bP.x
			}

			if limitingPoints[0].y > p.sP.y {
				limitingPoints[0].y = p.sP.y
			}

			if limitingPoints[1].y < p.sP.y {
				limitingPoints[1].y = p.sP.y
			}

			if limitingPoints[0].y > p.bP.y {
				limitingPoints[0].y = p.bP.y
			}

			if limitingPoints[1].y < p.bP.y {
				limitingPoints[1].y = p.bP.y
			}

		}
	}

	return limitingPoints
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func getDistance(p1, p2 point) int {
	return abs(p1.x-p2.x) + abs(p1.y-p2.y)
}

func min(x1, x2 int) int {
	if x1 < x2 {
		return x1
	}
	return x2
}

func max(x1, x2 int) int {
	if x1 > x2 {
		return x1
	}
	return x2
}

func positionWithoutBeacon(row int, upLimit int, pairs []pair) (int, [][2]int) {
	var result int

	// ranges without beacon
	var rwb [][2]int

	for _, p := range pairs {
		dist := getDistance(p.sP, p.bP)

		if p.sP.y-dist <= row && row <= p.sP.y+dist {
			distRow := abs(p.sP.y - row)
			distOccupied := (dist-distRow)*2 + 1
			pointLeft := p.sP.x - distOccupied/2
			if pointLeft < 0 {
				pointLeft = 0
			}
			pointRight := p.sP.x + distOccupied/2
			if pointRight > upLimit {
				pointRight = upLimit
			}

			rwb = append(rwb, [2]int{pointLeft, pointRight})
		}
	}

	ok := false
	for ok == false {
		ok = true
		for i := 0; i < len(rwb)-1; i++ {
			for j := i + 1; j < len(rwb); j++ {
				if rwb[i][0] <= rwb[j][0] && rwb[j][0] <= rwb[i][1] && rwb[i][1] <= rwb[j][1] {
					rwb[i][1] = rwb[j][1]
					rwb = append(rwb[:j], rwb[j+1:]...)
					ok = false
					continue
				}

				if rwb[i][0] <= rwb[j][0] && rwb[j][1] <= rwb[i][1] {
					rwb = append(rwb[:j], rwb[j+1:]...)
					ok = false
					continue
				}

				if rwb[j][0] <= rwb[i][0] && rwb[i][1] <= rwb[j][1] {
					ok = false
					rwb = append(rwb[:i], rwb[i+1:]...)
					continue
				}

				if rwb[j][0] <= rwb[i][0] && rwb[i][0] <= rwb[j][1] && rwb[j][1] <= rwb[i][1] {
					rwb[j][1] = rwb[i][1]
					rwb = append(rwb[:i], rwb[i+1:]...)
					ok = false
					continue
				}

			}
		}
	}

	for _, r := range rwb {
		result += abs(r[0]-r[1]) + 1
	}

	return result, rwb
}

func getPositionBeacon(upLimit int, pairs []pair) point {
	for i := 0; i <= upLimit; i++ {
		n, rwb := positionWithoutBeacon(i, upLimit, pairs)
		if n != upLimit+1 {
			return point{rwb[1][1] + 1, i}
		}
	}
	return point{}
}

func main() {
	pairs := getData("data.in")
	p := getPositionBeacon(4000000, pairs)
	fmt.Println(p.x * 4000000 + p.y)
}
