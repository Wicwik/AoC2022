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
	var crt [][]string
	current_row_pos := 0

	cycles_map := map[string]int{"addx": 2, "noop": 1}

	readFile, err := os.Open("./input.txt")

	register_x := 1
	cycle_counter := 0

	check(err)

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		str := fileScanner.Text()
		str_split := strings.Split(str, " ")

		if len(str_split) == 1 {
			cycle_counter++
			if (cycle_counter % 40) == 1 {
				crt = append(crt, []string{})
				current_row_pos = 0
				// fmt.Println(str)
			}

			if (register_x-1) == current_row_pos || register_x == current_row_pos || (register_x+1) == current_row_pos {
				crt[len(crt)-1] = append(crt[len(crt)-1], "#")
			} else {
				crt[len(crt)-1] = append(crt[len(crt)-1], ".")
			}

			current_row_pos++
		} else {
			to_add, err := strconv.Atoi(str_split[1])
			check(err)

			for i := 0; i < cycles_map[str_split[0]]; i++ {
				cycle_counter++

				if (cycle_counter % 40) == 1 {
					crt = append(crt, []string{})
					current_row_pos = 0
					// fmt.Println(str)
				}

				if (register_x-1) == current_row_pos || register_x == current_row_pos || (register_x+1) == current_row_pos {
					crt[len(crt)-1] = append(crt[len(crt)-1], "#")
				} else {
					crt[len(crt)-1] = append(crt[len(crt)-1], ".")
				}

				current_row_pos++

				if i == 1 {
					register_x += to_add
				}
			}
		}
	}

	for _, v := range crt {
		fmt.Println(v)
	}
}
