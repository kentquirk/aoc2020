package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strings"

	"github.com/kentquirk/stringset/v2"
)

func day06a(forms []string) int {
	total := 0
	for _, group := range forms {
		set := stringset.New()
		for _, c := range group {
			if c >= 'a' && c <= 'z' {
				set.Add(string(c))
			}
		}
		total += set.Length()
	}
	return total
}

func day06b(forms []string) int {
	total := 0
	for _, group := range forms {
		people := strings.Split(group, "\n")
		var intersectionSet *stringset.StringSet
		for _, person := range people {
			set := stringset.New()
			for _, c := range person {
				if c >= 'a' && c <= 'z' {
					set.Add(string(c))
				}
			}
			if intersectionSet == nil {
				intersectionSet = set
			} else {
				intersectionSet = intersectionSet.Intersection(set)
			}
		}

		total += intersectionSet.Length()
	}
	return total
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
	splitpat := regexp.MustCompile("\n[[:blank:]]*\n")
	forms := splitpat.Split(string(b), -1)
	fmt.Printf("%d questions were answered in part A\n", day06a(forms))
	fmt.Printf("%d questions were answered the same by all group participants in part B\n", day06b(forms))
}
