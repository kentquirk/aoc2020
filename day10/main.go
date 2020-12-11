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

func day10a(orig []int) int {
	data := []int{0}
	data = append(data, orig...)
	sort.Ints(data)
	data = append(data, data[len(data)-1]+3)
	diffs := make(map[int]int)
	for ix := 1; ix < len(data); ix++ {
		diff := data[ix] - data[ix-1]
		diffs[diff]++
	}
	return diffs[1] * diffs[3]
}

// This tests if a group is a "valid" group, meaning
// that all its elements are no more than 3 units
// apart.
func isValidGroup(group []int) bool {
	for ix := 1; ix < len(group); ix++ {
		if group[ix]-group[ix-1] > 3 {
			return false
		}
	}
	return true
}

// This is a basic hash routine that takes a slice
// of ints and returns an integer that characterizes it
// so we can use it as a hash key.
// We did it this way instead of using Go's hash
// functions because those work on and generate
// arrays of bytes, and it's tedious doing all the conversions.
// This is certainly not robust but it's effective for
// this purpose.
func groupHash(group []int) int {
	x := 0x45d9f3b
	for _, g := range group {
		x ^= ((g >> 16) ^ g) * 0x45d9f3b
		x = ((x >> 16) ^ x) * 0x45d9f3b
		x = (x >> 16) ^ x
	}
	return x
}

// This is a recursive routine that
// returns a collection of valid subgroups within
// a given group. Because the recursion generates
// multiple copies of some subgroups, we store them
// in a map that is keyed by a hash of the members
// of the group.
func countValids(group []int) map[int][]int {
	all := make(map[int][]int)
	if isValidGroup(group) {
		// add the hash of the group itself
		all[groupHash(group)] = group
		for i := 1; i < len(group)-1; i++ {
			subgroup := append([]int{}, group[0:i]...)
			subgroup = append(subgroup, group[i+1:]...)
			// now add all the subhashes
			for _, v := range countValids(subgroup) {
				all[groupHash(v)] = v
			}
		}
	}
	return all
}

func divideGroups(data []int) [][]int {
	// we're going to split data up into a slice of
	// slices -- each slice is a group where both ends
	// of the group are separated by 3 'jolts' from the next
	result := make([][]int, 0)
	firstIx := 0
	for ix := 0; ix < len(data)-1; ix++ {
		diff := data[ix+1] - data[ix]
		if diff == 3 {
			result = append(result, data[firstIx:ix+1])
			firstIx = ix + 1
		}
	}
	return result
}

func day10b(orig []int) int {
	data := []int{0}
	data = append(data, orig...)
	sort.Ints(data)
	data = append(data, data[len(data)-1]+3)

	groups := divideGroups(data)
	product := 1
	for _, group := range groups {
		valids := countValids(group)
		product *= len(valids)
	}
	return product
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
	data := []int{}
	for _, line := range lines {
		n, _ := strconv.Atoi(line)
		data = append(data, n)
	}
	fmt.Printf("product is %d\n", day10a(data))
	fmt.Printf("number of arrangements is %d\n", day10b(data))
}
