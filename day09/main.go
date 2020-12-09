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

func day09a(preambleLength int, data []int) int {
outer:
	for start := 0; start < len(data)-preambleLength; start++ {
		preamble := data[start : start+preambleLength]
		for i := 0; i < preambleLength-1; i++ {
			for j := i + 1; j < preambleLength; j++ {
				sum := preamble[i] + preamble[j]
				if data[start+preambleLength] == sum {
					continue outer
				}
			}
		}
		return data[start+preambleLength]
	}
	return -1
}

func sum(a []int) int {
	total := 0
	for _, i := range a {
		total += i
	}
	return total
}

func day09b(magic int, data []int) int {
	for windowSize := 2; windowSize < len(data); windowSize++ {
		for start := 0; start < len(data)-windowSize; start++ {
			if sum(data[start:start+windowSize]) == magic {
				result := make([]int, windowSize)
				copy(result, data[start:start+windowSize])
				sort.Ints(result)
				return result[0] + result[windowSize-1]
			}
		}
	}
	return -1
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
	data := make([]int, len(lines))
	for i, line := range lines {
		n, err := strconv.Atoi(line)
		if err != nil {
			log.Fatal(err)
		}
		data[i] = n
	}

	magic := day09a(25, data)
	fmt.Printf("Magic number was %d\n", magic)
	fmt.Printf("Weakness is %d\n", day09b(magic, data))
}
