package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Point struct {
	x, y int
}

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

func fall_one(p *Point, grid_map map[Point]bool, max_y int) *Point {
	if p.y+1 == max_y {
		return nil
	}

	if !grid_map[Point{p.x, p.y + 1}] {
		return &Point{p.x, p.y + 1}
	} else {
		if !grid_map[Point{p.x - 1, p.y + 1}] {
			return &Point{p.x - 1, p.y + 1}
		} else if !grid_map[Point{p.x + 1, p.y + 1}] {
			return &Point{p.x + 1, p.y + 1}
		}
	}
	return nil
}

func main() {
	readFile, err := os.Open("./input.txt")

	check(err)

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	grid_map := make(map[Point]bool)
	max_y := 0

	for fileScanner.Scan() {
		str := fileScanner.Text()

		cords_str := strings.Split(str, " -> ")

		for i := 0; i < len(cords_str)-1; i++ {
			from_str := strings.Split(cords_str[i], ",")
			to_str := strings.Split(cords_str[i+1], ",")

			from_x, _ := strconv.Atoi(from_str[0])
			from_y, _ := strconv.Atoi(from_str[1])
			to_x, _ := strconv.Atoi(to_str[0])
			to_y, _ := strconv.Atoi(to_str[1])

			if to_y > max_y {
				max_y = to_y
			}

			if from_x > to_x {
				tmp := from_x
				from_x = to_x
				to_x = tmp
			}

			if from_y > to_y {
				tmp := from_y
				from_y = to_y
				to_y = tmp
			}

			for x := from_x; x <= to_x; x++ {
				for y := from_y; y <= to_y; y++ {
					grid_map[Point{x, y}] = true
				}
			}

			// fmt.Println(from_x, from_y)
			// fmt.Println(to_x, to_y)
		}
	}

	max_y += 2

	fmt.Println(len(grid_map))

	start := &Point{500, 0}
	count := 0

	for {
		queue := []*Point{start}
		var curr_pos *Point

		for len(queue) > 0 {
			curr_pos, queue = queue[0], queue[1:]
			next_pos := fall_one(curr_pos, grid_map, max_y)

			if next_pos == nil {
				if curr_pos == start {
					fmt.Println(max_y, count+1)
					return
				}

				grid_map[*curr_pos] = true
				break
			}

			// fmt.Println(next_pos)
			queue = append(queue, next_pos)
		}

		count++
	}
}
