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

func calcSignalStrength(register_x int, cycle_counter int) int {
	return (register_x * cycle_counter)
}

func main() {
	sum := 0
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
			if cycle_counter == 20 || ((cycle_counter-20)%40 == 0) {
				sum += calcSignalStrength(register_x, cycle_counter)
				// fmt.Println(str)
			}

		} else {
			to_add, err := strconv.Atoi(str_split[1])
			check(err)

			for i := 0; i < cycles_map[str_split[0]]; i++ {
				cycle_counter++

				if cycle_counter == 20 || ((cycle_counter-20)%40 == 0) {
					sum += calcSignalStrength(register_x, cycle_counter)
					// fmt.Println(str)
				}

				if i == 1 {
					register_x += to_add
				}
			}
		}
	}

	fmt.Println(sum)
}
