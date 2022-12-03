// https://adventofcode.com/2022/day/3

package main

import (
	"fmt"

	"github.com/boljen/go-bitmap"
	"github.com/kcdbaba/advent2022/input"
)

func priority(char rune) int {
	code := int(char) - 94
	sign := 1
	if code < 0 {
		sign = -1
		code = -code
	}
	code = (code - 3)
	code = code * sign
	if code < 0 {
		code = code + 52
	}
	return code + 1
}

func letter(code int) string {
	if code < 27 {
		return string(rune(code + 96))
	} else {
		return string(rune(code - 26 + 64))
	}
}

func main() {
	rucksacks, _ := input.FileToLines("input/3.1.txt")
	var both_compartments []int

	/*for _, char := range "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ" {
		fmt.Println(string(char), priority(char))
	}*/

	for _, sack := range rucksacks {
		bmap := []byte{0, 0, 0, 0, 0, 0, 0}
		l := len(sack)
		for _, chr := range sack[:l/2] {
			bitmap.Set(bmap, priority(chr)-1, true)
		}
		for _, chr := range sack[l/2:] {
			pri := priority(chr)
			if bitmap.Get(bmap, pri-1) {
				both_compartments = append(both_compartments, pri)
				break
			}
		}
	}

	sum := 0
	for _, v := range both_compartments {
		sum = sum + v
	}
	fmt.Println(sum)

	var badges []int
	for i := 0; i < len(rucksacks); i += 3 {
		var first2 [][]byte
		for _, sack := range rucksacks[i : i+2] {
			sack_bmap := []byte{0, 0, 0, 0, 0, 0, 0}
			for _, chr := range sack {
				bitmap.Set(sack_bmap, priority(chr)-1, true)
			}
			first2 = append(first2, sack_bmap)
		}

		sack := rucksacks[i+2]
		for _, chr := range sack {
			pri := priority(chr)
			if bitmap.Get(first2[0], pri-1) && bitmap.Get(first2[1], pri-1) {
				badges = append(badges, pri)
				//fmt.Println((i+1)/3, letter(pri), pri)
				break
			}
		}
	}

	sum = 0
	for _, v := range badges {
		sum = sum + v
	}
	fmt.Println(sum)
}
