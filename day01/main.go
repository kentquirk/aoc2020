package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"

	"golang.org/x/tools/container/intsets"
)

const target = 2020

func copyAndApply(s *intsets.Sparse, fn func(int) (int, bool)) *intsets.Sparse {
	// because set iteration is destructive, we make a copy
	iterset := new(intsets.Sparse)
	result := new(intsets.Sparse)
	iterset.Copy(s)

	var v int
	for iterset.TakeMin(&v) {
		if n, ok := fn(v); ok {
			result.Insert(n)
		}
	}
	return result
}

func day1a(data *intsets.Sparse) {
	diffs := copyAndApply(data, func(v int) (int, bool) {
		return target - v, true
	})

	diffs.IntersectionWith(data)
	product := 1
	v := 0
	for diffs.TakeMin(&v) {
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
	data := new(intsets.Sparse)
	for _, line := range lines {
		n, err := strconv.Atoi(line)
		if err != nil {
			log.Fatal(err)
		}
		data.Insert(n)
	}
	day1a(data)
}
