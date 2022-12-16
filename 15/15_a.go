package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
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

func manhattan_dist(a Point, b Point) int {
	return int(math.Abs(float64(a.x)-float64(b.x)) + math.Abs(float64(a.y)-float64(b.y)))
}

func main() {
	x_squares_map := make(map[int]bool)
	sensors := make(map[Point]bool)
	beacons := make(map[Point]bool)

	readFile, err := os.Open("./input.txt")

	check(err)

	y_max := 2000000

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		str := fileScanner.Text()
		var s_x, s_y, b_x, b_y int
		fmt.Sscanf(str, "Sensor at x=%d, y=%d: closest beacon is at x=%d, y=%d", &s_x, &s_y, &b_x, &b_y)

		fmt.Println("Sensor:", s_x, s_y)
		sensors[Point{s_x, s_y}] = true

		fmt.Println("Beacon:", b_x, b_y)
		beacons[Point{b_x, b_y}] = true

		fmt.Println("Manhattan:", manhattan_dist(Point{s_x, s_y}, Point{b_x, b_y}))

		dist := manhattan_dist(Point{s_x, s_y}, Point{b_x, b_y})

		if dist < int(math.Abs(float64(s_y)-float64(y_max))) {
			continue
		}

		dist -= int(math.Abs(float64(s_y) - float64(y_max)))

		for i := s_x - dist; i <= s_x+dist; i++ {
			x_squares_map[i] = true
		}
	}

	for b := range beacons {
		if b.y == y_max {
			delete(x_squares_map, b.x)
		}
	}

	fmt.Println(len(x_squares_map))
}
