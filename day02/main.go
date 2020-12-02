package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

// PWData holds the password data as specified in the problem set
type PWData struct {
	Letter   string
	MinCount int
	MaxCount int
	Password string
}

// Validate checks if the pwdata meets the validation rules
func (p PWData) Validate() bool {
	if p.Letter == "" {
		return false
	}
	n := strings.Count(p.Password, p.Letter)
	return n >= p.MinCount && n <= p.MaxCount
}

func parsePassword(s string) PWData {
	pat := regexp.MustCompile("([0-9]+)-([0-9]+) ([a-z]): ([a-z]+)")
	matches := pat.FindStringSubmatch(s)
	if len(matches) != 5 {
		log.Fatalf("line: %s does not match pattern! %s", s, matches)
	}
	mincount, _ := strconv.Atoi(matches[1])
	maxcount, _ := strconv.Atoi(matches[2])
	return PWData{
		Letter:   matches[3],
		MinCount: mincount,
		MaxCount: maxcount,
		Password: matches[4],
	}
}

func day02a(lines []string) {
	nvalid := 0
	for _, line := range lines {
		pwd := parsePassword(line)
		if pwd.Validate() {
			nvalid++
		}
	}
	fmt.Printf("%d passwords are valid\n", nvalid)
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

	day02a(lines)
}
