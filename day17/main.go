package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func day17a(lines []string) int {
	board := ParseBoard(lines)
	// board.printActives()
	for i := 0; i < 6; i++ {
		board = board.Generation()
		// board.printActives()
	}
	return board.CountActive()
}

func day17b(lines []string) int {
	board := ParseBoard4(lines)
	// board.printActives()
	for i := 0; i < 6; i++ {
		board = board.Generation()
		// board.printActives()
	}
	return board.CountActive()
}

func main() {
	f, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	b, err := ioutil.ReadAll(f)
	if err != nil {
		log.Fatal(err)
	}
	lines := strings.Split(string(b), "\n")
	fmt.Println(day17a(lines))
	fmt.Println(day17b(lines))
}
