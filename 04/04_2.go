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

func sliceAtoi(sa []string) ([]int, error) {
	si := make([]int, 0, len(sa))
	for _, a := range sa {
		i, err := strconv.Atoi(a)
		if err != nil {
			return si, err
		}
		si = append(si, i)
	}
	return si, nil
}

func makeRange(min, max int) []int {
	a := make([]int, max-min+1)
	for i := range a {
		a[i] = min + i
	}
	return a
}

func main() {
	readFile, err := os.Open("./input.txt")

	check(err)

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	count := 0
	for fileScanner.Scan() {
		str := fileScanner.Text()
		pair := strings.Split(str, ",")

		elf1, err := sliceAtoi(strings.Split(pair[0], "-"))
		check(err)

		elf2, err := sliceAtoi(strings.Split(pair[1], "-"))
		check(err)

		arr1 := makeRange(elf1[0], elf1[1])
		arr2 := makeRange(elf2[0], elf2[1])
		m := make(map[int]bool)

		for _, v := range arr1 {
			m[v] = true
		}

		for _, v := range arr2 {
			if _, ok := m[v]; ok {
				count++
				break
			}
		}

		// if elf1[1] >= elf2[0] {
		// 	count++
		// 	fmt.Println(elf1, elf2)
		// }
	}

	fmt.Println(count)
}
