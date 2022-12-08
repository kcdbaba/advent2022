// https://adventofcode.com/2022/day/8
package main

import (
	"fmt"
	"strconv"
	"unicode/utf8"

	"github.com/boljen/go-bitmap"
	"github.com/golang-collections/collections/stack"
	"github.com/kcdbaba/advent2022/input"
)

func print_grid(arr [][]uint8) {
	for _, row := range arr {
		fmt.Println(row)
	}
}

func print_scores(arr [][]int) {
	for _, row := range arr {
		fmt.Println(row)
	}
}

func grid_column(arr [][]uint8, col_idx int, reverse bool) []uint8 {
	num_rows := len(arr)
	var col = make([]uint8, num_rows)
	for i := 0; i < num_rows; i++ {
		idx := i
		if reverse {
			idx = num_rows - i - 1
		}
		col[i] = arr[idx][col_idx]
	}
	return col
}

func grid_row_reversed(arr [][]uint8, row_idx int) []uint8 {
	num_cols := len(arr[0])
	var row = make([]uint8, num_cols)
	for j := 0; j < num_cols; j++ {
		row[j] = arr[row_idx][num_cols-j-1]
	}
	return row
}

func find_valley_distances(arr []uint8, len int) []int {
	var arr_scores = make([]int, len)
	var elevated_indices stack.Stack
	for i := 1; i < len; i++ {
		if arr[i-1] < arr[i] {
			if elevated_indices.Peek() != nil {
				for elevated_indices.Peek() != nil &&
					arr[elevated_indices.Peek().(int)] < arr[i] {
					elevated_indices.Pop()
				}
				horizon_idx := 0
				if elevated_indices.Len() > 0 {
					horizon_idx = elevated_indices.Peek().(int)
					for elevated_indices.Peek() != nil &&
						arr[elevated_indices.Peek().(int)] == arr[i] {
						elevated_indices.Pop()
					}
				}
				arr_scores[i] = i - horizon_idx
			} else {
				arr_scores[i] = i
			}
		} else {
			if arr[i-1] >= arr[i] {
				elevated_indices.Push(i - 1)
			}
			arr_scores[i] = 1
		}
	}
	return arr_scores
}

func main() {
	lines, _ := input.FileToLines("input/8.1.txt")
	var rows = len(lines)
	var cols = utf8.RuneCountInString(lines[0])
	//fmt.Println(rows, cols)

	var grid = make([][]uint8, rows)
	for i := 0; i < rows; i++ {
		grid[i] = make([]uint8, cols)
	}
	for i := 0; i < rows; i++ {
		runes := []rune(lines[i])
		for j := 0; j < cols; j++ {
			tmp, _ := strconv.ParseUint(string(runes[j]), 10, 8)
			// add 1 for convenience while searching masima
			grid[i][j] = uint8(tmp) + 1
		}
	}

	var visible = make([]byte, rows*cols/8+1)
	for j := 0; j < cols; j++ {
		var top_down_peak uint8 = 0
		var bottom_up_peak uint8 = 0
		for i := 0; i < rows; i++ {
			if grid[i][j] > top_down_peak {
				bitmap.Set(visible, i*rows+j, true)
				top_down_peak = grid[i][j]
			}
			if grid[rows-i-1][j] > bottom_up_peak {
				bitmap.Set(visible, (rows-i-1)*rows+j, true)
				bottom_up_peak = grid[rows-i-1][j]
			}
		}
	}
	for i := 0; i < rows; i++ {
		var left_right_peak uint8 = 0
		var right_left_peak uint8 = 0
		for j := 0; j < cols; j++ {
			if grid[i][j] > left_right_peak {
				bitmap.Set(visible, i*rows+j, true)
				left_right_peak = grid[i][j]
			}
			if grid[i][cols-j-1] > right_left_peak {
				bitmap.Set(visible, i*rows+cols-j-1, true)
				right_left_peak = grid[i][cols-j-1]

			}
		}
	}

	sum_visible := 0
	for i := 0; i < rows*cols; i++ {
		if bitmap.Get(visible, i) {
			sum_visible += 1
		}
	}

	fmt.Println("Number of visible trees:", sum_visible)

	var max_scene_score int = 0
	var scene_score = make([][]int, rows)
	var brute_scene_score = make([][]int, rows)

	for i := 0; i < rows; i++ {
		scene_score[i] = make([]int, cols)
		brute_scene_score[i] = make([]int, cols)
		if i > 0 && i < rows-1 {
			for j := 1; j < cols-1; j++ {
				scene_score[i][j] = 1
				brute_scene_score[i][j] = 1
			}
		}
	}

	set_scene_score_col := func(col_scores []int, col_idx int, reverse bool) {
		for i := 0; i < rows; i++ {
			idx := i
			if reverse {
				idx = rows - i - 1
			}
			scene_score[idx][col_idx] *= col_scores[i]
		}
	}

	set_scene_score_row := func(scores []int, row_idx int, reverse bool) {
		for j := 0; j < cols; j++ {
			idx := j
			if reverse {
				idx = cols - j - 1
			}
			scene_score[row_idx][idx] *= scores[j]
		}
	}

	//print_grid(grid)
	for j := 0; j < cols; j++ {
		arr := grid_column(grid, j, false)
		col_scores := find_valley_distances(arr, rows)
		set_scene_score_col(col_scores, j, false)

		arr = grid_column(grid, j, true)
		col_scores = find_valley_distances(arr, rows)
		set_scene_score_col(col_scores, j, true)
	}
	for i := 0; i < rows; i++ {
		arr := grid[i]
		scores := find_valley_distances(arr, cols)
		set_scene_score_row(scores, i, false)

		arr = grid_row_reversed(grid, i)
		scores = find_valley_distances(arr, cols)
		set_scene_score_row(scores, i, true)
	}

	//print_scores(scene_score)
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if scene_score[i][j] > max_scene_score {
				max_scene_score = scene_score[i][j]
			}
		}
	}
	fmt.Println("Maximum scenic score is:", max_scene_score)

	for i := 1; i < rows-1; i++ {
		for j := 1; j < cols-1; j++ {
			var top = 0
			for iw := i - 1; iw >= 0; iw-- {
				if grid[iw][j] >= grid[i][j] {
					top = iw
					break
				}
			}
			brute_scene_score[i][j] *= i - top

			var bot = rows - 1
			for iw := i + 1; iw < rows; iw++ {
				if grid[iw][j] >= grid[i][j] {
					bot = iw
					break
				}
			}
			brute_scene_score[i][j] *= bot - i

			var left = 0
			for jw := j - 1; jw >= 0; jw-- {
				if grid[i][jw] >= grid[i][j] {
					left = jw
					break
				}
			}
			brute_scene_score[i][j] *= j - left

			var right = cols - 1
			for jw := j + 1; jw < cols; jw++ {
				if grid[i][jw] >= grid[i][j] {
					right = jw
					break
				}
			}
			brute_scene_score[i][j] *= right - j
		}
	}

	//print_scores(brute_scene_score)
	max_scene_score = 0
	var max_coords = make([]int, 2)
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if brute_scene_score[i][j] > max_scene_score {
				max_scene_score = brute_scene_score[i][j]
				max_coords = []int{i, j}
			}
		}
	}
	fmt.Println("Maximum brute force scenic score is:", max_scene_score)
	fmt.Println("Max scenic score found at coords:", max_coords)
}
