// https://adventofcode.com/2022/day/4

package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/kcdbaba/advent2022/input"
)

func elf_sections(line string) [2][2]int {
	var sections [2][2]int
	strs := strings.SplitN(line, ",", 2)
	for i, str := range strs {
		bounds := strings.SplitN(str, "-", 2)
		sections[i][0], _ = strconv.Atoi(bounds[0])
		sections[i][1], _ = strconv.Atoi(bounds[1])
	}
	return sections
}

func is_overlap(s1 [2]int, s2 [2]int) (bool, bool) {
	var small, big [2]int
	if s1[1]-s1[0] <= s2[1]-s2[0] {
		small = s1
		big = s2
	} else {
		small = s2
		big = s1
	}
	start_compar := small[0] >= big[0] && small[0] <= big[1]
	end_compar := small[1] >= big[0] && small[1] <= big[1]
	//fmt.Println(s1, s2, start_compar, end_compar)
	return start_compar && end_compar, start_compar || end_compar
}

func main() {
	lines, _ := input.FileToLines("input/4.1.txt")
	var sections [][2][2]int
	var count, count_partial int
	for _, line := range lines {
		line_sects := elf_sections(line)
		sections = append(sections, line_sects)
		full, part := is_overlap(line_sects[0], line_sects[1])
		if full {
			count += 1
		}
		if part {
			count_partial += 1
		}
	}
	fmt.Println(count)
	fmt.Println(count_partial)
}
