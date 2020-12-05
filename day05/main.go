package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func convertToBinary(s string) (int, error) {
	rep := strings.NewReplacer(
		"F", "0",
		"B", "1",
		"L", "0",
		"R", "1",
	)
	bs := rep.Replace(s)
	n, err := strconv.ParseInt(bs, 2, 32)
	return int(n), err
}

func day05a(lines []string) int {
	max := 0
	for _, line := range lines {
		n, err := convertToBinary(line)
		if err != nil {
			log.Fatal(err)
		}
		if n > max {
			max = n
		}
	}
	return max
}

func day05b(lines []string) int {
	seats := make([]int, len(lines))
	for i, line := range lines {
		n, err := convertToBinary(line)
		if err != nil {
			log.Fatal(err)
		}
		seats[i] = n
	}
	sort.Ints(seats)
	prev := seats[0] - 1
	for _, seat := range seats {
		if prev+1 != seat {
			return prev + 1
		}
		prev = seat
	}
	return 0
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
	fmt.Printf("Part A max value was %d\n", day05a(lines))
	fmt.Printf("Part B missing seat value was %d\n", day05b(lines))
}
