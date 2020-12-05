package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"regexp"
)

// StringSet is intended to store a set of strings
type StringSet map[string]struct{}

// Add puts an item in the set
func (s *StringSet) Add(ary ...string) {
	for _, a := range ary {
		(*s)[a] = struct{}{}
	}
}

// Contains returns whether the value is in the set
func (s StringSet) Contains(n string) bool {
	_, ok := s[n]
	return ok
}

// Intersect creates a new StringSet that is the intersection of s and t
func (s StringSet) Intersect(t StringSet) StringSet {
	result := make(StringSet)
	for k := range s {
		if t.Contains(k) {
			result.Add(k)
		}
	}
	return result
}

// Passport represents a passport object as described in the problem.
type Passport map[string]string

// ValidationA validates the rules for partA
func (p Passport) ValidationA(requiredKeys StringSet) bool {
	for key := range requiredKeys {
		if _, ok := p[key]; !ok {
			return false
		}
	}
	return true
}

// ParseFrom builds a Passport object from a string that defines it.
// The input string is expected to contain 3-letter keys, with arbitrary
// values, separated by colon (:) and each key-value pair separated by
// whitespace.
func (p *Passport) ParseFrom(s string) {
	kvpat := regexp.MustCompile("([[:lower:]]{3}):([^ \t\n\r]+)")
	matches := kvpat.FindAllStringSubmatch(s, -1)
	for _, ms := range matches {
		key := ms[1]
		value := ms[2]
		(*p)[key] = value
	}
}

func day04a(passports []string) int {
	requiredKeysA := make(StringSet)
	requiredKeysA.Add(
		"byr", // (Birth Year)
		"iyr", // (Issue Year)
		"eyr", // (Expiration Year)
		"hgt", // (Height)
		"hcl", // (Hair Color)
		"ecl", // (Eye Color)
		"pid", // (Passport ID)
	// "cid", // (Country ID)
	)

	validCount := 0
	for _, pp := range passports {
		passport := make(Passport)
		passport.ParseFrom(pp)
		if passport.ValidationA(requiredKeysA) {
			validCount++
		}
	}
	return validCount
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
	passports := splitpat.Split(string(b), -1)
	fmt.Printf("%d were valid in part A\n", day04a(passports))
}
