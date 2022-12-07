package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type dir struct {
	name      string
	size      uint64
	parentdir *dir
	subdirs   map[string]*dir
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func sum(array []int) int {
	result := 0
	for _, v := range array {
		result += v
	}
	return result
}

var current_dir *dir
var root_dir dir
var dir_sizes []uint64

func update_parents(parent *dir, filesize uint64) {
	if parent != nil {
		fmt.Println(parent)

		parent.size += filesize
		update_parents(parent.parentdir, filesize)
	}
}

func tree_sum(current *dir, sum uint64) uint64 {

	for _, value := range current.subdirs {
		if value.size <= 100000 {
			sum += tree_sum(value, value.size)
		} else {
			sum += tree_sum(value, 0)
		}
	}
	return sum
}

func find_dirs_to_remove(current *dir, dirsize uint64) {
	if current.size >= dirsize {
		dir_sizes = append(dir_sizes, current.size)
	}

	for _, value := range current.subdirs {
		find_dirs_to_remove(value, dirsize)
	}
}

func parseCommand(command string, fileScanner *bufio.Scanner) {
	command_slc := strings.Split(command, " ")
	// fmt.Println(command_slc)

	if command_slc[0] == "ls" {
		for fileScanner.Scan() {
			str := fileScanner.Text()
			str_split := strings.Split(str, " ")
			// fmt.Println(str)

			if str[0] == '$' {
				command_slc = str_split[1:]
				break
			} else if str_split[0] == "dir" {
				current_dir.subdirs[str_split[1]] = &dir{str_split[1], 0, current_dir, make(map[string]*dir)}
			} else {
				filesize, err := strconv.Atoi(str_split[0])
				check(err)

				fmt.Println(current_dir)
				current_dir.size += uint64(filesize)

				update_parents(current_dir.parentdir, uint64(filesize))
			}
		}
	}

	// fmt.Println(command_slc)
	if command_slc[0] == "cd" {
		// fmt.Println("CD")
		if command_slc[1] == ".." {
			current_dir = current_dir.parentdir
		} else if command_slc[1] != "/" {
			// fmt.Println(command_slc[1])
			current_dir = current_dir.subdirs[command_slc[1]]
		}
	}
}

func main() {
	readFile, err := os.Open("./input.txt")

	check(err)

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	root_dir = dir{"/", 0, nil, make(map[string]*dir)}
	current_dir = &root_dir

	for fileScanner.Scan() {
		str := fileScanner.Text()

		if str[0] == '$' {
			parseCommand(str[2:], fileScanner)
		}
	}

	used_space := root_dir.size
	free_space := 70000000 - used_space
	to_free := 30000000 - free_space

	fmt.Println(to_free)

	fmt.Println(tree_sum(&root_dir, 0))

	find_dirs_to_remove(&root_dir, to_free)

	var m uint64
	for i, e := range dir_sizes {
		if i == 0 || e < m {
			m = e
		}
	}

	fmt.Println(m)
}
