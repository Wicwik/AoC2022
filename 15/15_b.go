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

type rangeValues struct {
	x_values [][2]int
}

func correctIdx(arr [][2]int, itm [2]int) int {
	for i := len(arr) - 1; i >= 0; i-- {
		a := arr[i]
		if itm[0] == a[0] && itm[1] > a[1] {
			return i + 1
		}
		if itm[0] > a[0] {
			return i + 1
		}
	}
	return 0
}

func (r *rangeValues) insert(n [2]int) {
	idx := correctIdx(r.x_values, n)
	if len(r.x_values) == 0 {
		r.x_values = [][2]int{n}
	} else if len(r.x_values) == idx {
		r.x_values = append(r.x_values, n)
	} else {
		r.x_values = append(r.x_values[:idx+1], r.x_values[idx:]...)
		r.x_values[idx] = n
	}
	newVals := [][2]int{r.x_values[0]}
	for _, i := range r.x_values[1:] {
		if newVals[len(newVals)-1][1] >= i[0]-1 {
			newVals[len(newVals)-1][1] = int(math.Max(float64(newVals[len(newVals)-1][1]), float64(i[1])))
		} else {
			newVals = append(newVals, i)
		}
	}
	r.x_values = newVals
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
	sensors := make(map[Point]bool)
	beacons := make(map[Point]bool)

	readFile, err := os.Open("./input.txt")

	check(err)

	available_squares := map[int]rangeValues{}
	max := 4000000

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
		i := 0

		for x := s_x - dist; x <= s_x+dist; x++ {
			if 0 <= x && x <= max {
				y_range := [2]int{int(math.Max(float64(s_y-i), 0)), int(math.Min(float64(s_y+i), float64(max)))}
				a := available_squares[x]
				a.insert(y_range)
				available_squares[x] = a
			}

			if x >= s_x {
				i--
			} else {
				i++
			}
		}

	}

	for k, e := range available_squares {
		if len(e.x_values) != 1 {
			fmt.Println("Part 2:", k*4000000+e.x_values[0][1]+1)
		}
	}

}
