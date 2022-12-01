// https://adventofcode.com/2022/day/1
package main

import (
	"container/heap"
	"fmt"
	"strconv"

	"github.com/kcdbaba/advent2022/input"
	"github.com/kcdbaba/advent2022/utils"
)

func main() {
	lines, _ := input.FileToLines("input/1.1.txt")

	var cals, window int
	h := &utils.IntHeap{}
	for _, v := range lines {
		if v == "" {
			heap.Push(h, window)
			window = 0
		} else {
			cals, _ = strconv.Atoi(v)
			window = window - cals
		}
	}
	// 1.1
	var max, max2, max3 int
	max = heap.Pop(h).(int)
	fmt.Println(max * -1)
	// 1.2
	max2 = heap.Pop(h).(int)
	max3 = heap.Pop(h).(int)
	fmt.Println((max + max2 + max3) * -1)
}
