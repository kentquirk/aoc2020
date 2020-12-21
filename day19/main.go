package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strings"
)

func buildRules(lines []string) (int, map[string]string) {
	rules := make(map[string]string)
	n := 0
	for i, line := range lines {
		if line == "" {
			n = i
			break
		}
		splits := strings.Split(line, ": ")
		ix := splits[0]
		contents := splits[1]
		rules[ix] = "(" + contents + ")"
	}
	return n + 1, rules
}

func substitute(s string, rules map[string]string) string {
	numpat := regexp.MustCompile("[^0-9]+")
	for {
		splits := numpat.Split(s, -1)
		replacements := make(map[string]struct{})
		for _, split := range splits[1 : len(splits)-1] {
			replacements[split] = struct{}{}
		}

		if len(replacements) == 0 {
			break
		}

		for repl := range replacements {
			replpat := regexp.MustCompile(`\b` + repl + `\b`)
			s = replpat.ReplaceAllString(s, rules[repl])
		}
	}

	// clean up a little bit
	replacer := strings.NewReplacer(
		`("a")`, "a",
		`("b")`, "b",
		" ", "")
	s = replacer.Replace(s)

	pat := regexp.MustCompile(`\(([ab]+)\)`)
	s = pat.ReplaceAllString(s, `$1`)

	return s
}

func day19a(lines []string) int {
	_, rules := buildRules(lines)
	// fmt.Println(rules)
	re := substitute(rules["0"], rules)
	// fmt.Println(re)
	pat := regexp.MustCompile("^" + re + "$")
	count := 0
	for _, line := range lines {
		if pat.MatchString(line) {
			count++
		}
	}
	return count
}

func day19b(lines []string) int {
	n, rules := buildRules(lines)
	messages := lines[n:]

	for rule := range rules {
		fmt.Printf("%s: %s\n", rule, substitute(rules[rule], rules))
	}

	rule42 := substitute(rules["42"], rules)
	rule31 := substitute(rules["31"], rules)
	// fmt.Println(rule42, rule31)
	pat42 := regexp.MustCompile("^" + rule42)
	pat31 := regexp.MustCompile(rule31 + "$")

	matches31 := func(s string) int {
		count31 := 0
		for {
			loc := pat31.FindStringIndex(s)
			if loc == nil {
				break
			}
			count31++
			fmt.Printf("31: '%s'  '%s'\n", s[:loc[0]], s[loc[0]:loc[1]])
			s = s[:loc[0]]
		}
		if len(s) != 0 {
			return 0
		}
		return count31
	}

	count := 0
	for _, line := range messages {
		count42 := 0
		count31 := 0
		if !(strings.HasPrefix(line, "a") || strings.HasPrefix(line, "b")) {
			continue
		}
		s := line
		// fmt.Printf("\nstarting: %s\n", s)

		// loop, matching 42s at the beginning; for each one, check
		// if the rest of the string entirely matches a bunch of 31s.
		// If so, and if the number of 31s is less than the number of 42s
		// we found, then we have a match that works for rule 0.
		for {
			loc := pat42.FindStringIndex(s)
			if loc == nil {
				// fmt.Printf("%d: fail %s\n", ix, line)
				break
			}
			count42++
			// fmt.Printf("42: '%s'  '%s'\n", s[loc[0]:loc[1]], s[loc[1]:])
			s = s[loc[1]:]
			count31 = matches31(s)
			if count31 >= 1 && count31 < count42 {
				// fmt.Printf("31: %d\n", count31)
				// fmt.Printf("%d: %s %d %d\n", ix, line, count42, count31)
				count++
				break
			} else {
				// fmt.Printf("31: %d\n", count31)
			}
		}
	}

	return count
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
	fmt.Println(day19a(lines))
	fmt.Println(day19b(lines))
}
