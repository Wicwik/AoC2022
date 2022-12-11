package main

import (
	"fmt"
	"sort"

	llq "github.com/emirpasic/gods/queues/linkedlistqueue"
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

type monkey struct {
	items       *llq.Queue
	operation_f func(int, int) int
	operation_n int
	test        [3]int
}

func add(a int, b int) int {
	return a + b
}

func sub(a int, b int) int {
	return a - b
}

func mult(a int, b int) int {
	return a * b
}

func div(a int, b int) int {
	return a / b
}

func main() {
	monkey_counter := []int{0, 0, 0, 0, 0, 0, 0, 0}

	monkeys := [8]monkey{
		{&llq.Queue{}, mult, 3, [3]int{5, 2, 3}},
		{&llq.Queue{}, add, 8, [3]int{11, 4, 7}},
		{&llq.Queue{}, add, 2, [3]int{2, 5, 3}},
		{&llq.Queue{}, add, 4, [3]int{13, 1, 5}},
		{&llq.Queue{}, mult, 19, [3]int{7, 7, 6}},
		{&llq.Queue{}, add, 5, [3]int{3, 4, 1}},
		{&llq.Queue{}, mult, -1, [3]int{17, 0, 2}},
		{&llq.Queue{}, add, 1, [3]int{19, 6, 0}},
	}

	monkeys[0].items = llq.New()
	monkeys[0].items.Enqueue(65)
	monkeys[0].items.Enqueue(78)

	// 54, 78, 86, 79, 73, 64, 85, 88
	monkeys[1].items = llq.New()
	monkeys[1].items.Enqueue(54)
	monkeys[1].items.Enqueue(78)
	monkeys[1].items.Enqueue(86)
	monkeys[1].items.Enqueue(79)
	monkeys[1].items.Enqueue(73)
	monkeys[1].items.Enqueue(64)
	monkeys[1].items.Enqueue(85)
	monkeys[1].items.Enqueue(88)

	// 69, 97, 77, 88, 87
	monkeys[2].items = llq.New()
	monkeys[2].items.Enqueue(69)
	monkeys[2].items.Enqueue(97)
	monkeys[2].items.Enqueue(77)
	monkeys[2].items.Enqueue(88)
	monkeys[2].items.Enqueue(87)

	monkeys[3].items = llq.New()
	monkeys[3].items.Enqueue(99)

	// 60, 57, 52
	monkeys[4].items = llq.New()
	monkeys[4].items.Enqueue(60)
	monkeys[4].items.Enqueue(57)
	monkeys[4].items.Enqueue(52)

	// 91, 82, 85, 73, 84, 53
	monkeys[5].items = llq.New()
	monkeys[5].items.Enqueue(91)
	monkeys[5].items.Enqueue(82)
	monkeys[5].items.Enqueue(85)
	monkeys[5].items.Enqueue(73)
	monkeys[5].items.Enqueue(84)
	monkeys[5].items.Enqueue(53)

	// 88, 74, 68, 56
	monkeys[6].items = llq.New()
	monkeys[6].items.Enqueue(88)
	monkeys[6].items.Enqueue(74)
	monkeys[6].items.Enqueue(68)
	monkeys[6].items.Enqueue(56)

	// 54, 82, 72, 71, 53, 99, 67
	monkeys[7].items = llq.New()
	monkeys[7].items.Enqueue(54)
	monkeys[7].items.Enqueue(82)
	monkeys[7].items.Enqueue(72)
	monkeys[7].items.Enqueue(71)
	monkeys[7].items.Enqueue(53)
	monkeys[7].items.Enqueue(99)
	monkeys[7].items.Enqueue(67)

	rounds := 10000
	prod := 1

	for _, m := range monkeys {
		prod *= m.test[0]
	}

	for i := 0; i < rounds; i++ {
		for j, m := range monkeys {
			for !m.items.Empty() {
				item, _ := m.items.Dequeue()
				monkey_counter[j]++

				// fmt.Println(monkey_counter[j])

				var new_item int
				if m.operation_n == -1 {
					new_item = m.operation_f(item.(int), item.(int))
				} else {
					new_item = m.operation_f(item.(int), m.operation_n)
				}

				// new_item /= 3

				new_item %= prod
				if uint64(new_item)%uint64(m.test[0]) == 0 {
					monkeys[m.test[1]].items.Enqueue(new_item)
				} else {
					monkeys[m.test[2]].items.Enqueue(new_item)
				}
			}
		}
	}

	// for _, m := range monkeys {
	// 	fmt.Println(m.items.Values()...)
	// }

	sort.Ints(monkey_counter)
	fmt.Println(monkey_counter[len(monkey_counter)-1] * monkey_counter[len(monkey_counter)-2])
}
