package main

import (
	"bufio"
	"fmt"
	"math"
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

func biggerSlice(arr []int, v int) int {
	for i, e := range arr {
		if e >= v {
			return i
		}
	}

	return -1
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

func reverse(numbers []int) []int {
	newNumbers := make([]int, 0, len(numbers))
	for i := len(numbers) - 1; i >= 0; i-- {
		newNumbers = append(newNumbers, numbers[i])
	}
	return newNumbers
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

	max_scenic := 0
	for i, row := range arr {
		for j, v := range row {
			scenic_score := 1
			if i == 0 || j == 0 || i == (len(arr)-1) || j == (len(row)-1) {
				continue
			} else {
				col := getColumn(arr, j)

				// fmt.Println(col[:i], col[i+1:])
				// fmt.Println(row[:j], row[j+1:])

				sub_arr := col[:i]
				m := maxSlice(sub_arr)
				// fmt.Print(m, " ")
				if m < v {
					scenic_score *= len(sub_arr)
				} else {
					m_i := biggerSlice(reverse(sub_arr), v)
					m_i += i + 1
					scenic_score *= int(math.Abs(float64(m_i - i)))
					// fmt.Println(int(math.Abs(float64(m_i - i))))
					// fmt.Println(m_i, i)
				}

				sub_arr = col[i+1:]
				m = maxSlice(sub_arr)
				// fmt.Print(m, " ")
				if m < v {
					scenic_score *= len(sub_arr)
				} else {
					m_i := biggerSlice(sub_arr, v)
					m_i += i + 1
					scenic_score *= int(math.Abs(float64(m_i - i)))
					// fmt.Println(int(math.Abs(float64(m_i - i))))
				}

				sub_arr = row[:j]
				m = maxSlice(sub_arr)
				// fmt.Print(m, " ")
				if m < v {
					scenic_score *= len(sub_arr)
				} else {
					m_i := biggerSlice(reverse(sub_arr), v)
					m_i += j + 1
					scenic_score *= int(math.Abs(float64(m_i - j)))
					// fmt.Println(int(math.Abs(float64(m_i - j))))
				}

				sub_arr = row[j+1:]
				m = maxSlice(sub_arr)
				// fmt.Print(m, "\n")
				if m < v {
					scenic_score *= len(sub_arr)
				} else {
					m_i := biggerSlice(sub_arr, v)
					m_i += j + 1
					scenic_score *= int(math.Abs(float64(m_i - j)))
					// fmt.Println(int(math.Abs(float64(m_i - j))))
				}

				// fmt.Println(scenic_score)

				if scenic_score > max_scenic {
					max_scenic = scenic_score
				}
			}
		}
	}

	fmt.Println(max_scenic)
}
