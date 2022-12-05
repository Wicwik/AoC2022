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

		if (elf1[0] >= elf2[0] && elf1[1] <= elf2[1]) || (elf2[0] >= elf1[0] && elf2[1] <= elf1[1]) {
			count++
			// fmt.Println(elf1, elf2)
		}
	}

	fmt.Println(count)
}
