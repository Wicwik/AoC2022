package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"

	"github.com/emirpasic/gods/sets/hashset"
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

type Pair struct {
	a, b int
}

type dirPair struct {
	a Pair
	b int
}

func follow_head(head Pair, tail Pair) Pair {
	x, y := head.a-tail.a, head.b-tail.b
	abx, aby := int(math.Abs(float64(x))), int(math.Abs(float64(y)))

	var newx int
	var newy int
	if abx > 1 || aby > 1 {
		if x == 0 {
			newx = tail.a
		} else {
			newx = tail.a + (x / abx)
		}

		if y == 0 {
			newy = tail.b
		} else {
			newy = tail.b + (y / aby)
		}

		return Pair{newx, newy}
	}

	return tail
}

func main() {
	directions := map[string]Pair{
		"L": {-1, 0},
		"R": {1, 0},
		"U": {0, 1},
		"D": {0, -1},
	}

	readFile, err := os.Open("./input.txt")

	knots := 10

	check(err)

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	var moves []dirPair

	for fileScanner.Scan() {
		str := fileScanner.Text()
		str_split := strings.Split(str, " ")
		steps, err := strconv.Atoi(str_split[1])
		check(err)

		moves = append(moves, dirPair{directions[str_split[0]], steps})
	}

	var rope []Pair
	for i := 0; i < knots; i++ {
		rope = append(rope, Pair{0, 0})
	}

	visited := hashset.New()
	visited.Add(rope[len(rope)-1])

	for _, move := range moves {
		for i := 0; i < int(move.b); i++ {
			head := rope[0]
			rope[0] = Pair{head.a + move.a.a, head.b + move.a.b}
			for j := 1; j < knots; j++ {
				rope[j] = follow_head(rope[j-1], rope[j])
			}
			visited.Add(rope[len(rope)-1])
		}
	}

	fmt.Println(visited.Size())
}
