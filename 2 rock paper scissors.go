// https://adventofcode.com/2022/day/2
package main

import (
	"fmt"
	"strings"

	"github.com/kcdbaba/advent2022/input"
)

func main() {
	scoring_map := map[string]map[string]int{
		"A": map[string]int{
			"X": 4,
			"Y": 8,
			"Z": 3,
		},
		"B": map[string]int{
			"X": 1,
			"Y": 5,
			"Z": 9,
		},
		"C": map[string]int{
			"X": 7,
			"Y": 2,
			"Z": 6,
		},
	}

	lines, _ := input.FileToLines("input/2.1.txt")
	split_chall_resp := func(l string) (string, string) {
		strs := strings.SplitN(l, " ", 2)
		return strs[0], strs[1]
	}

	var sum int = 0
	for _, line := range lines {
		challenge, response := split_chall_resp(line)
		sum = sum + scoring_map[challenge][response]
	}
	fmt.Println(sum)

	sum = 0
	for _, line := range lines {
		challenge, result := split_chall_resp(line)
		result_score := (int(result[0]) - int('X')) * 3
		hand_score := 0
		if result_score == 6 {
			hand_score = (int(challenge[0])-int('A')+1)%3 + 1
		} else if result_score == 3 {
			hand_score = int(challenge[0]) - int('A') + 1
		} else {
			hand_score = (int(challenge[0])-int('A')+2)%3 + 1
		}
		sum = sum + result_score + hand_score
	}
	fmt.Println(sum)
}
