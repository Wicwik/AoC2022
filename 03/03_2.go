package main

import (
	"bufio"
	"fmt"
	"os"
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

	check(err)

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	result := 0
	var group []string
	for fileScanner.Scan() {
		str := fileScanner.Text()

		group = append(group, str)

		if len(group) == 3 {
			for _, character := range group[0] {
				first_i := strings.IndexRune(group[1], character)
				second_i := strings.IndexRune(group[2], character)
				if (first_i >= 0) && (second_i >= 0) {
					// println(first_i, second_i)

					if int(character) > 90 {
						result += (int(byte(character)) - 96)
					} else {
						result += (int(byte(character)) - 64 + 26)
					}
					println(byte(character))

					group = nil
					break
				}

			}
		}

		// fmt.Println(first, second)
	}

	fmt.Println(result)
}
