// https://adventofcode.com/2022/day/1
package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/kcdbaba/advent2022/input"
)

type FileType string

const (
	FILE FileType = "regular"
	DIR  FileType = "directory"
)

type Tree struct {
	parent   *Tree
	children map[string]*Tree
	name     string
	size     int
}

func NewTree() *Tree {
	var root = new(Tree)
	root.parent = nil
	root.name = ""
	root.children = make(map[string]*Tree)
	return root
}

func (parent *Tree) add_file(name string, size int) (err error) {
	if _, exists := parent.children[name]; exists {
		return errors.New("Name already exists")
	} else if name == "" {
		return errors.New("Name cannot be empty")
	}
	var file = new(Tree)
	file.parent = parent
	file.name = name
	file.size = size
	file.children = nil
	parent.children[name] = file

	ancestor := parent
	for ancestor != nil {
		ancestor.size += size
		ancestor = ancestor.parent
	}
	return nil
}

func (parent *Tree) add_dir(name string) (err error) {
	if _, exists := parent.children[name]; exists {
		return errors.New("Name already exists")
	} else if name == "" {
		return errors.New("Name cannot be empty")
	}
	var dir = new(Tree)
	dir.parent = parent
	dir.name = name
	dir.children = make(map[string]*Tree)
	parent.children[name] = dir
	return nil
}

func print_tree_recursive(t *Tree, indent string) {
	if t.children == nil {
		fmt.Println(indent, t.name, t.size)
	} else {
		if t.parent == nil {
			fmt.Println(indent, "/")
		} else {
			fmt.Println(indent, t.name, "/")
		}
		for _, child := range t.children {
			print_tree_recursive(child, indent+"  ")
		}
	}
}

func (root *Tree) print_tree() {
	print_tree_recursive(root, "")
}

func traverse_tree(line string, root *Tree, curr *Tree) *Tree {
	if line[:1] == "$" {
		tokens := strings.SplitN(line[2:], " ", 2)
		if tokens[0] == "cd" {
			switch tokens[1] {
			case "/":
				return root
			case "..":
				return curr.parent
			default:
				return curr.children[tokens[1]]
			}
		}
	} else {
		build_tree(line, curr)
	}
	return curr
}

func build_tree(line string, parent *Tree) {
	if line[:3] == "dir" {
		name := line[4:]
		parent.add_dir(name)
	} else {
		tokens := strings.SplitN(line, " ", 2)
		size, _ := strconv.Atoi(tokens[0])
		name := tokens[1]
		parent.add_file(name, size)
	}
}

func sum_size_limit_dirs(t *Tree, limit int) int {
	if t.children == nil {
		return 0
	}

	var sum int = 0
	if t.size <= limit {
		sum = t.size
	}
	for _, child := range t.children {
		sum += sum_size_limit_dirs(child, limit)
	}
	return sum
}

func min_dir_size(t *Tree, threshold int) int {
	var min_size = t.size
	for _, child := range t.children {
		if child.children != nil && child.size > threshold {
			child_size := min_dir_size(child, threshold)
			if child_size < min_size {
				min_size = child_size
			}
		}
	}
	return min_size
}

func main() {
	lines, _ := input.FileToLines("input/7.1.txt")
	root := NewTree()
	curr := root
	for _, line := range lines {
		curr = traverse_tree(line, root, curr)
	}
	//root.print_tree()
	//fmt.Println(root.size)

	const dir_size_limit = 100000
	fmt.Println("1. Sum of dir with size of atmost", dir_size_limit, ":",
		sum_size_limit_dirs(root, dir_size_limit))

	const total_space = 70000000
	const reqd_space = 30000000
	free_space := total_space - root.size
	to_del := reqd_space - free_space

	fmt.Println("Additional free space required:", to_del)
	fmt.Println("2. Smallest dir size to delete:", min_dir_size(root, to_del))
}
