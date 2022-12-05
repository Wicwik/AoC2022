package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

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

func main() {
	readFile, err := os.Open("./input.txt")

	crates := [][]string{
		{"F", "H", "M", "T", "V", "L", "D"},
		{"P", "N", "T", "C", "J", "G", "Q", "H"},
		{"H", "P", "M", "D", "S", "R"},
		{"F", "V", "B", "L"},
		{"Q", "L", "G", "H", "N"},
		{"P", "M", "R", "G", "D", "B", "W"},
		{"Q", "L", "H", "C", "R", "N", "M", "G"},
		{"W", "L", "C"},
		{"T", "M", "Z", "J", "Q", "L", "D", "R"},
	}

	// crates := [][]string{
	// 	{"N", "Z"},
	// 	{"D", "C", "M"},
	// 	{"P"},
	// }

	check(err)

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		str := fileScanner.Text()
		str_splt := strings.Split(str, " ")

		take, err := strconv.Atoi(str_splt[1])
		check(err)

		from, err := strconv.Atoi(str_splt[3])
		check(err)

		to, err := strconv.Atoi(str_splt[5])
		check(err)

		// fmt.Println(crates, take, from, to)

		arr := crates[from-1]
		last := arr[:take]
		fmt.Println(arr[take:], last)
		crates[from-1] = arr[take:]

		tmp := []string{}
		for i, _ := range last {
			// fmt.Println(v)
			tmp = append(tmp, last[len(last)-1-i])
		}

		for _, v := range tmp {
			fmt.Println(v)
			crates[to-1] = append([]string{v}, crates[to-1]...)
		}
		// fmt.Println(crts)
	}

	for _, v := range crates {
		fmt.Print(v[0])
	}

	fmt.Println()
	fmt.Println(crates)
}
