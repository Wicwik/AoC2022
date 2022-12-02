package main

import (
	"bufio"
	"fmt"
	"os"
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

	points_map := make(map[string]int)
	wld_map := make(map[string]int)
	points := 0

	points_map["A"] = 1
	points_map["B"] = 2
	points_map["C"] = 3

	wld_map["X"] = -1
	wld_map["Y"] = 0
	wld_map["Z"] = 1

	check(err)

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		str := fileScanner.Text()

		to_add := points_map[str[0:1]] + wld_map[str[2:3]]

		if to_add == 0 || to_add == 3 {
			points += 3
		} else {
			points += (to_add % 3)
		}

		if str[2:3] == "Z" {
			points += 6
		}

		if str[2:3] == "Y" {
			points += 3
		}
	}

	fmt.Println(points)
}
