package main

import (
	"encoding/json"
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

func compare(left, right any) int {
	lefts, leftok := left.([]any)
	rights, rightok := right.([]any)

	if !leftok && !rightok {
		return int(left.(float64) - right.(float64))
	} else if !leftok {
		lefts = []any{left}
	} else if !rightok {
		rights = []any{right}
	}

	for i := 0; i < len(lefts) && i < len(rights); i++ {
		cmp := compare(lefts[i], rights[i])

		if cmp != 0 {
			return cmp
		}
	}

	return len(lefts) - len(rights)
}

func main() {
	readFile, err := os.ReadFile("./input.txt")

	check(err)

	in_right_order := 0

	for i, v := range strings.Split(string(readFile), "\n\n") {
		str_pair := strings.Split(v, "\n")

		var left, right any

		err := json.Unmarshal([]byte(str_pair[0]), &left)
		check(err)

		err = json.Unmarshal([]byte(str_pair[1]), &right)
		check(err)

		if compare(left, right) <= 0 {
			in_right_order += i + 1
		}
	}

	fmt.Println(in_right_order)
}
