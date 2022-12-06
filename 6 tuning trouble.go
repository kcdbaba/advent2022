// https://adventofcode.com/2022/day/6

package main

import (
	"fmt"
	"strconv"

	"github.com/kcdbaba/advent2022/input"
)

type RingBuffer struct {
	size   int
	buffer []interface{}
	idx    int
}

func (r *RingBuffer) write_ring_buffer(item interface{}) {
	if len(r.buffer) < r.size {
		r.buffer = append(r.buffer, item)
	}
	r.buffer[r.idx] = item
	r.idx = (r.idx + 1) % r.size
}

func (r *RingBuffer) full_and_distinct() bool {
	if len(r.buffer) < r.size {
		return false
	}
	var items = make(map[interface{}]struct{})
	for i := 0; i < r.size; i++ {
		cand := r.buffer[i]
		if _, exists := items[cand]; exists {
			return false
		} else {
			items[cand] = struct{}{}
		}
	}
	return true
}

func NewRingBuffer(size int) *RingBuffer {
	newbuff := new(RingBuffer)
	newbuff.size = size
	newbuff.buffer = make([]interface{}, 0, size)
	return newbuff
}

func print_rune_buffer(r *RingBuffer) {
	for _, i := range r.buffer {
		fmt.Print(strconv.QuoteRune(i.(rune)))
	}
	fmt.Print("\n")
}

func main() {
	lines, _ := input.FileToLines("input/6.1.txt")
	packet_start_buff := NewRingBuffer(4)
	msg_start_buff := NewRingBuffer(14)

	var start_of_packet_idx int
	for i, chr := range []rune(lines[0]) {
		packet_start_buff.write_ring_buffer(chr)
		//print_rune_buffer(packet_start_buff)
		if packet_start_buff.full_and_distinct() {
			fmt.Println(i + 1)
			start_of_packet_idx = i - 3
			break
		}
	}
	for i, chr := range []rune(lines[0])[start_of_packet_idx:] {
		msg_start_buff.write_ring_buffer(chr)
		//print_rune_buffer(msg_start_buff)
		if msg_start_buff.full_and_distinct() {
			fmt.Println(start_of_packet_idx + i + 1)
			break
		}
	}
}
