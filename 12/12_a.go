package main

import (
	"bufio"
	"fmt"
	"os"

	pq "github.com/emirpasic/gods/queues/priorityqueue"
	"github.com/emirpasic/gods/utils"
)

type Element struct {
	cords    Pair
	priority int //number of steps
}

type Pair struct {
	a, b int
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

func byPriority(a, b interface{}) int {
	priorityA := a.(Element).priority
	priorityB := b.(Element).priority
	return utils.IntComparator(priorityA, priorityB) // "-" descending order
}

func dijsktra(start Pair, exit Pair, map_grid [][]int) int {
	visited := make([][]bool, len(map_grid))
	for i := range visited {
		visited[i] = make([]bool, len(map_grid[len(map_grid)-1]))
	}

	print_pos := make([][]string, len(map_grid))
	for i := range print_pos {
		print_pos[i] = make([]string, len(map_grid[len(map_grid)-1]))
		for j := range print_pos[i] {
			print_pos[i][j] = "."
		}
	}

	queue := pq.NewWith(byPriority)
	moves := [4]Pair{{-1, 0}, {1, 0}, {0, 1}, {0, -1}} // L R U D

	queue.Enqueue(Element{start, 0})
	// fmt.Println(queue.Values(), moves)
	for !queue.Empty() {
		// for i := 0; i < 5; i++ {
		// fmt.Println("queue:", queue.Values())
		pos, _ := queue.Dequeue()
		// fmt.Println("current pos:", pos)

		if visited[pos.(Element).cords.b][pos.(Element).cords.a] {
			continue
		}

		visited[pos.(Element).cords.b][pos.(Element).cords.a] = true

		// print_pos[pos.(Element).cords.b][pos.(Element).cords.a] = string(rune(map_grid[pos.(Element).cords.b][pos.(Element).cords.a]) + 97)
		// for _, v := range print_pos {
		// 	fmt.Println(v)
		// }

		for _, m := range moves {
			new_cord := Pair{pos.(Element).cords.a + m.a, pos.(Element).cords.b + m.b}
			// fmt.Println(new_cord)

			if new_cord.a < 0 || new_cord.b < 0 || new_cord.b >= len(map_grid) || new_cord.a >= len(map_grid[len(map_grid)-1]) {
				continue
			}

			if visited[new_cord.b][new_cord.a] {
				// fmt.Println("visited")
				continue
			}

			old_cord_height := map_grid[pos.(Element).cords.b][pos.(Element).cords.a]
			new_cord_height := map_grid[new_cord.b][new_cord.a]
			height_diff := new_cord_height - old_cord_height
			new_prio := pos.(Element).priority + 1

			// fmt.Println("info:", height_diff, new_cord, pos.(Element).cords, map_grid[pos.(Element).cords.a][pos.(Element).cords.b], map_grid[new_cord.a][new_cord.b])
			if height_diff <= 1 {
				queue.Enqueue(Element{new_cord, new_prio})

				if new_cord.a == exit.a && new_cord.b == exit.b {
					return new_prio
				}
			}
		}
	}

	return 0
}

func main() {
	readFile, err := os.Open("./input.txt")

	check(err)

	map_grid := [][]int{}

	var start Pair
	var exit Pair

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		str := fileScanner.Text()

		map_grid = append(map_grid, []int{})

		for i, v := range str {
			if v == 'S' {
				start = Pair{i, len(map_grid) - 1}
				map_grid[len(map_grid)-1] = append(map_grid[len(map_grid)-1], 0)
			} else if v == 'E' {
				exit = Pair{i, len(map_grid) - 1}
				map_grid[len(map_grid)-1] = append(map_grid[len(map_grid)-1], 25)
			} else {
				map_grid[len(map_grid)-1] = append(map_grid[len(map_grid)-1], int(v)-97)
			}
		}
	}

	for _, v := range map_grid {
		fmt.Println(v)
	}
	fmt.Println(start, exit)

	fmt.Println(dijsktra(start, exit, map_grid))
}
