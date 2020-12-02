package main

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
)

// PWData2 holds the password data as specified in the problem set
type PWData2 struct {
	Letter   string
	Index1   int
	Index2   int
	Password string
}

// Validate checks if the pwdata meets the validation rules
func (p PWData2) Validate() bool {
	if p.Letter == "" {
		return false
	}
	// requirement is that exactly one of these is true
	at1 := false
	at2 := false
	if p.Index1 <= len(p.Password) && p.Password[p.Index1-1:p.Index1] == p.Letter {
		at1 = true
	}
	if p.Index2 <= len(p.Password) && p.Password[p.Index2-1:p.Index2] == p.Letter {
		at2 = true
	}

	return at1 != at2
}

func parseNewPassword(s string) PWData2 {
	pat := regexp.MustCompile("([0-9]+)-([0-9]+) ([a-z]): ([a-z]+)")
	matches := pat.FindStringSubmatch(s)
	if len(matches) != 5 {
		log.Fatalf("line: %s does not match pattern! %s", s, matches)
	}
	mincount, _ := strconv.Atoi(matches[1])
	maxcount, _ := strconv.Atoi(matches[2])
	return PWData2{
		Letter:   matches[3],
		Index1:   mincount,
		Index2:   maxcount,
		Password: matches[4],
	}
}
func day02b(lines []string) {
	nvalid := 0
	for _, line := range lines {
		pwd := parseNewPassword(line)
		if pwd.Validate() {
			nvalid++
		}
	}
	fmt.Printf("%d passwords are valid\n", nvalid)
}
