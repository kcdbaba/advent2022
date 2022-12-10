// https://adventofcode.com/2022/day/10
package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/kcdbaba/advent2022/input"
)

type CPU struct {
	pointer      int
	x            int
	cycle        uint
	next_result  int
	skip_till    uint
	probe_cycles map[uint]struct{}
	signals      map[uint]int
	running      bool
	crt          [40][6]bool
}

func MakeCPU(probe_at_cycles []uint) *CPU {
	var cpu = new(CPU)
	cpu.x = 1
	cpu.pointer = 0
	cpu.cycle = 1
	cpu.probe_cycles = make(map[uint]struct{})
	for _, cycle := range probe_at_cycles {
		cpu.probe_cycles[cycle] = struct{}{}
	}
	cpu.signals = make(map[uint]int)
	cpu.skip_till = 0
	cpu.next_result = 1
	return cpu
}

func (cpu *CPU) run_cycle(lines []string) {
	if cpu.skip_till == cpu.cycle {
		cpu.x = cpu.next_result
		cpu.pointer += 1
	}

	_, ok := cpu.probe_cycles[cpu.cycle]
	if ok {
		cpu.signals[cpu.cycle] = cpu.x * int(cpu.cycle)
	}

	crt_x := (cpu.cycle - 1) / 40
	crt_y := (cpu.cycle - 1) % 40
	if int(crt_y) >= cpu.x-1 && int(crt_y) <= cpu.x+1 {
		cpu.crt[crt_y][crt_x] = true
	}

	if cpu.skip_till > cpu.cycle {
		// skip
	} else if cpu.pointer >= len(lines) {
		cpu.running = false
		return
	} else {
		tokens := strings.SplitN(lines[cpu.pointer], " ", 2)
		switch tokens[0] {
		case "noop":
			cpu.pointer += 1
		case "addx":
			add_arg, _ := strconv.Atoi(tokens[1])
			cpu.next_result = cpu.x + add_arg
			cpu.skip_till = cpu.cycle + 2
		}
	}

	cpu.cycle += 1
}

func main() {
	//lines, _ := input.FileToLines("input/10.test.txt")
	//var cpu = MakeCPU([]uint{1, 2, 3, 4, 5})
	lines, _ := input.FileToLines("input/10.1.txt")
	var cpu = MakeCPU([]uint{20, 60, 100, 140, 180, 220})

	cpu.running = true
	for cpu.running {
		cpu.run_cycle(lines)
	}

	fmt.Println(cpu.signals)
	sum_signals := 0
	for _, signals := range cpu.signals {
		sum_signals += signals
	}
	fmt.Println("1:", sum_signals)

	for j := 0; j < 6; j++ {
		for i := 0; i < 40; i++ {
			if cpu.crt[i][j] {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Print("\n")
	}
}
