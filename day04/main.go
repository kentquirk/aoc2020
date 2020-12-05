package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"regexp"
)

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

// ValidationB validates the rules for partA
func (p Passport) ValidationB(validators ValidationPartB) bool {
	for key, f := range validators {
		value, ok := p[key]
		if !ok {
			return false
		}
		if !f(value) {
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

func day04b(passports []string) int {
	validators := make(ValidationPartB)
	validators["byr"] = yearValidatorBuilder(1920, 2002)
	validators["iyr"] = yearValidatorBuilder(2010, 2020)
	validators["eyr"] = yearValidatorBuilder(2020, 2030)
	validators["hgt"] = heightValidator
	validators["hcl"] = hairValidator
	validators["ecl"] = eyeValidator
	validators["pid"] = passportIDValidator

	validCount := 0
	for _, pp := range passports {
		passport := make(Passport)
		passport.ParseFrom(pp)
		if passport.ValidationB(validators) {
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
	fmt.Printf("%d were valid in part B\n", day04b(passports))
}
