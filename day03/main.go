package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func bangFactor(lines []string, right int, down int) int {
	rowwidth := len(lines[0])
	collisions := 0
	posx := 0
	posy := 0
	for posy < len(lines) {
		index := posx % rowwidth
		if lines[posy][index] == '#' {
			collisions++
		}
		posx += right
		posy += down
	}
	return collisions
}

func day03a(lines []string) int {
	return bangFactor(lines, 3, 1)
}

func day03b(lines []string) int {
	a := []int{
		bangFactor(lines, 1, 1),
		bangFactor(lines, 3, 1),
		bangFactor(lines, 5, 1),
		bangFactor(lines, 7, 1),
		bangFactor(lines, 1, 2),
	}
	prod := 1
	for _, v := range a {
		prod *= v
	}
	return prod
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
	fmt.Printf("read %d lines, each %d units wide\n", len(lines), len(lines[0]))
	fmt.Printf("There were %d collisions from part A\n", day03a(lines))
	fmt.Printf("There were %d collisions from part B\n", day03b(lines))
}
