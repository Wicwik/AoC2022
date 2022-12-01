package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
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
	var fileNums []int
	var fileSums []int

	for fileScanner.Scan() {
		str := fileScanner.Text()

		if str != "" {
			num, err := strconv.Atoi(str)
			check(err)

			fileNums = append(fileNums, num)
		} else {
			sum := sum(fileNums)

			fileSums = append(fileSums, sum)

			fileNums = nil
		}
	}

	sort.Sort(sort.Reverse(sort.IntSlice(fileSums)))
	fmt.Printf("%v", fileSums)
	fmt.Printf("%d", sum(fileSums[0:3]))
}
