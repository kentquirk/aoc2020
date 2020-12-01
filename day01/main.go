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
	// make the difference set from target-each element
	diffs := make(IntSet)
	for d := range data {
		diffs.Add(target - d)
	}
	// intersect it with the data
	result := diffs.Intersect(data)

	// there will be two results; multiply them
	product := 1
	for v := range result {
		fmt.Println(v)
		product *= v
	}
	fmt.Println(product)
}

func day1b(data IntSet) {
	// This time, we have to find three elements that add to 2020, so we make the difference set from
	// all combinations of two elements (that aren't the same element)
	diffs := make(IntSet)
	data2 := make(IntSet)
	for d1 := range data {
		for d2 := range data {
			if d1 != d2 {
				data2.Add(d1 + d2)
			}
		}
	}
	for d := range data2 {
		diffs.Add(target - d)
	}

	result := diffs.Intersect(data)

	// There will be 3 elements that add correctly so we just multiply them
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
	day1b(data)
}
