package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

// IntSet is intended to store a set of integers
type IntSet map[int]struct{}

// Add puts an item in the set
func (i *IntSet) Add(n int) {
	(*i)[n] = struct{}{}
}

// Contains returns whether the value is in the set
func (i IntSet) Contains(n int) bool {
	_, ok := i[n]
	return ok
}

// Intersect creates a new IntSet that is the intersection of i and j
func (i IntSet) Intersect(j IntSet) IntSet {
	result := make(IntSet)
	for k := range i {
		if j.Contains(k) {
			result.Add(k)
		}
	}
	return result
}

const target = 2020

func day1a(data IntSet) {
	diffs := make(IntSet)
	for d := range data {
		diffs.Add(target - d)
	}
	result := diffs.Intersect(data)

	product := 1
	for v := range result {
		fmt.Println(v)
		product *= v
	}
	fmt.Println(product)
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
	data := make(IntSet)
	for _, line := range lines {
		n, err := strconv.Atoi(line)
		if err != nil {
			log.Fatal(err)
		}
		data.Add(n)
	}
	day1a(data)
}
