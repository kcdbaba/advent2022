// https://adventofcode.com/2022/day/5

package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/golang-collections/collections/stack"
	"github.com/kcdbaba/advent2022/input"
)

type Move struct {
	num  int
	from int
	to   int
}

func print_rune_stacks(stacks []stack.Stack) {
	var max int
	for i := 0; i < len(stacks); i++ {
		if stacks[i].Len() > max {
			max = stacks[i].Len()
		}
	}
	var out = make([]string, max)
	for i := 0; i < len(stacks); i++ {
		var reverse = stack.New()
		crate := stacks[i].Pop()
		for crate != nil {
			reverse.Push(crate)
			crate = stacks[i].Pop()
		}
		for i := 0; i < max; i++ {
			crate := reverse.Pop()
			if crate == nil {
				crate = ' '
			}
			out[i] = out[i] + string(crate.(rune))
		}
	}
	for i := max - 1; i >= 0; i-- {
		fmt.Println(out[i])
	}
}

func main() {
	lines, _ := input.FileToLines("input/5.1.txt")
	var initial [][]rune
	var moves []Move
	for _, line := range lines {
		if line == "" {
		} else if line[:4] == "move" {
			splits := strings.SplitN(line[5:], " from ", 2)
			num, _ := strconv.Atoi(splits[0])
			splits = strings.SplitN(splits[1], " to ", 2)
			from, _ := strconv.Atoi(splits[0])
			to, _ := strconv.Atoi(splits[1])
			moves = append(moves, Move{num, from, to})
		} else {
			initial = append(initial, []rune(line))
		}
	}

	stacks := make([]stack.Stack, 9)
	for i := len(initial) - 2; i >= 0; i-- {
		for j := 0; j < 9; j++ {
			crate := initial[i][j*4+1]
			if crate != ' ' {
				stacks[j].Push(crate)
			}
		}
	}
	//fmt.Println(stacks, moves)
	//print_rune_stacks(stacks)

	copy_stacks := make([]stack.Stack, len(stacks))
	copy(copy_stacks, stacks)

	for _, move := range moves {
		for i := 0; i < move.num; i++ {
			crate := stacks[move.from-1].Pop()
			stacks[move.to-1].Push(crate)
		}
	}

	for i := 0; i < 9; i++ {
		fmt.Print(string(stacks[i].Peek().(rune)))
	}
	fmt.Println()

	for _, move := range moves {
		var temp = stack.New()
		for i := 0; i < move.num; i++ {
			crate := copy_stacks[move.from-1].Pop()
			temp.Push(crate)
		}
		for i := 0; i < move.num; i++ {
			crate := temp.Pop()
			copy_stacks[move.to-1].Push(crate)
		}

	}

	for i := 0; i < 9; i++ {
		fmt.Print(string(copy_stacks[i].Peek().(rune)))
	}
	fmt.Println()
}
