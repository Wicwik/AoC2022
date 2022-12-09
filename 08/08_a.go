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

func maxSlice(arr []int) int {
	m := 0

	for _, e := range arr {
		if e > m {
			m = e
		}
	}

	return m
}

func getColumn(board [][]int, columnIndex int) (column []int) {
	column = make([]int, 0)
	for _, row := range board {
		column = append(column, row[columnIndex])
	}
	return column
}

func main() {
	readFile, err := os.Open("./input.txt")

	check(err)

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	var arr [][]int

	for fileScanner.Scan() {
		str := fileScanner.Text()

		arr = append(arr, []int{})
		for _, v := range str {
			arr[len(arr)-1] = append(arr[len(arr)-1], int(v-'0'))
		}

		// fmt.Println(arr)
	}

	count := 0
	for i, row := range arr {
		for j, v := range row {
			if i == 0 || j == 0 || i == (len(arr)-1) || j == (len(row)-1) {
				count++
			} else {
				col := getColumn(arr, j)

				// fmt.Println("test: ", col[:2])

				if maxSlice(col[:i]) < v || maxSlice(col[i+1:]) < v || maxSlice(row[:j]) < v || maxSlice(row[j+1:]) < v {
					count++
					// fmt.Println(col[:i], col[i+1:])
					// fmt.Println(row[:j], row[j+1:])

					// fmt.Println(maxSlice(col[:i]), maxSlice(col[i+1:]), maxSlice(row[:j]), maxSlice(row[j+1:]), v)
				}
			}
		}
	}

	fmt.Println(count)
}
