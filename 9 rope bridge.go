// https://adventofcode.com/2022/day/9
package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/kcdbaba/advent2022/input"
)

type Coordinates struct {
	x int
	y int
}
type Direction = Coordinates

var Up Direction = Direction{0, -1}
var Down Direction = Direction{0, 1}
var Left Direction = Direction{-1, 0}
var Right Direction = Direction{1, 0}

func (pos *Coordinates) Move(dir Direction) {
	pos.x += dir.x
	pos.y += dir.y
}

func ParseMove(s string) (Direction, int) {
	tokens := strings.SplitN(s, " ", 2)
	var dir Direction
	switch tokens[0] {
	case "U":
		dir = Up
	case "D":
		dir = Down
	case "L":
		dir = Left
	case "R":
		dir = Right
	}
	num, _ := strconv.Atoi(tokens[1])
	return dir, num
}

func sign(a int) int {
	if a < 0 {
		return -1
	} else if a > 0 {
		return 1
	}
	return 0
}

func MoveTail(head *Coordinates, tail *Coordinates) {
	xdist := head.x - tail.x
	ydist := head.y - tail.y
	if -1 <= xdist && xdist <= 1 && -1 <= ydist && ydist <= 1 {
		return
	} else if xdist == 0 {
		tail.y += sign(ydist)
	} else if ydist == 0 {
		tail.x += sign(xdist)
	} else {
		tail.x += sign(xdist)
		tail.y += sign(ydist)
	}
}

func main() {
	lines, _ := input.FileToLines("input/9.1.txt")
	var knots = make([]Coordinates, 10)
	var tail_coords = make([]map[Coordinates]struct{}, 2)
	for i := 0; i < 2; i++ {
		tail_coords[i] = make(map[Coordinates]struct{})
	}

	//fmt.Println(head, tail)
	for _, line := range lines {
		dir, steps := ParseMove(line)
		for i := 0; i < steps; i++ {
			knots[0].Move(dir)
			for i := 1; i < 10; i++ {
				MoveTail(&knots[i-1], &knots[i])
			}
			_, ok := tail_coords[0][knots[1]]
			if !ok {
				tail_coords[0][knots[1]] = struct{}{}
			}
			_, ok = tail_coords[1][knots[9]]
			if !ok {
				tail_coords[1][knots[9]] = struct{}{}
			}
			//fmt.Println(head, tail)
		}
	}

	for i := 0; i < 2; i++ {
		coords := make([]Coordinates, 0)
		for k := range tail_coords[i] {
			coords = append(coords, k)
		}
		fmt.Println(i+1, ":", len(coords))
	}
}
