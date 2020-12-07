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

// Containment indicates how many children of which index are used
type Containment struct {
	qty  int
	name string // name of the child color
}

// Bag defines the color and contents of a bag
type Bag struct {
	adjective string
	color     string
	children  []Containment
}

// Name returns a canonical name for the bag
func (b Bag) Name() string {
	return b.adjective + " " + b.color
}

// NewBag constructs an empty bag
func NewBag(adj string, col string) Bag {
	return Bag{
		adjective: adj,
		color:     col,
		children:  make([]Containment, 0),
	}
}

func parseBag(s string) (int, Bag) {
	pat := regexp.MustCompile("([[:digit:]]*) ?([[:lower:]]+) ([[:lower:]]+)")
	match := pat.FindStringSubmatch(s)
	bag := NewBag(match[2], match[3])
	n, _ := strconv.Atoi(match[1])
	return n, bag
}

// line looks like <adj> <color> bags contain [<n> <adj> <color> bag[s][,]]...
func parseLine(line string) Bag {
	pat := regexp.MustCompile(
		` bags\.| bags, | bag\.| bag, | bags contain no other bags\.| bags contain `)
	matches := pat.Split(line, -1)
	if matches == nil {
		log.Fatalf("line '%s' didn't parse!", line)
	}
	_, bag := parseBag(matches[0])
	for _, m := range matches[1:] {
		if m == "" {
			continue
		}
		n, child := parseBag(m)
		bag.children = append(bag.children, Containment{
			qty:  n,
			name: child.Name(),
		})
	}
	return bag
}

func parseLines(lines []string) map[string]Bag {
	var bags = make(map[string]Bag)

	for _, line := range lines {
		bag := parseLine(line)
		bags[bag.Name()] = bag
	}
	return bags
}

func day07a(target string, bags map[string]Bag) int {
	found := map[string]struct{}{}
	for name, bag := range bags {
		if name == target {
			continue
		}
		if searchA(bag, target, bags, found) {
			found[name] = struct{}{}
		}
	}
	return len(found)
}

func searchA(bag Bag, target string, bags map[string]Bag, found map[string]struct{}) bool {
	for _, child := range bag.children {
		if child.name == target {
			return true
		}
		if _, ok := found[child.name]; ok {
			return true
		}
		if searchA(bags[child.name], target, bags, found) {
			return true
		}
	}
	return false
}

func countBagsInChildren(target string, bags map[string]Bag) int {
	bag := bags[target]
	total := 0
	for _, child := range bag.children {
		total += (1 + countBagsInChildren(child.name, bags)) * child.qty
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
	lines := strings.Split(string(b), "\n")
	bags := parseLines(lines)

	target := "shiny gold"
	fmt.Printf("%d bags were found that could contain %s\n", day07a(target, bags), target)

	fmt.Printf("%s bags contain %d other bags\n", target, countBagsInChildren(target, bags))
}
