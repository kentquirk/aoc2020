package main

import (
	"fmt"
)

func day15a(endAt int, data []int) int {
	visited := make(map[int]int)
	for i, n := range data[:len(data)-1] {
		visited[n] = i + 1
	}
	prev := data[len(data)-1]
	// for i, n := range data {
	// 	fmt.Println(i+1, n)
	// }
	for turn := len(data) + 1; ; turn++ {
		var next int
		if prevTurn, ok := visited[prev]; ok {
			next = turn - prevTurn - 1
		} else {
			next = 0
		}
		visited[prev] = turn - 1
		// fmt.Println(turn, next)
		prev = next
		if turn == endAt {
			break
		}
	}
	return prev
}

func main() {
	const endAt = 30000000
	data := []int{0, 3, 6}
	fmt.Println(day15a(endAt, data))
}
