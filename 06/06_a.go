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

	for fileScanner.Scan() {
		str := fileScanner.Text()

		for i := 3; i < len(str); i++ {
			found := true
			for _, v := range str[i-3 : i+1] {
				// fmt.Println(str[i-3 : i+1])
				if strings.Count(str[i-3:i+1], string(v)) > 1 {
					found = false
					break
				}
			}

			if !found {
				continue
			} else {
				fmt.Println(i + 1)
				break
			}
		}
	}
}
