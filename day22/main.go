package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

func score(q *IntQueue) int {
	total := 0
	for multiplier := q.Size(); q.Size() != 0; multiplier-- {
		total += multiplier * q.Dequeue()
	}
	return total
}

func day22a(p1, p2 *IntQueue) int {
	for p1.Size() != 0 && p2.Size() != 0 {
		c1 := p1.Dequeue()
		c2 := p2.Dequeue()
		if c1 > c2 {
			p1.Enqueue(c1)
			p1.Enqueue(c2)
		} else {
			p2.Enqueue(c2)
			p2.Enqueue(c1)
		}
	}

	if p1.Size() != 0 {
		return score(p1)
	}
	return score(p2)
}

// returns 1 or 2, the number of the winner of this round
func recursiveRound(p1, p2 *IntQueue, previous map[int]struct{}, depth int) int {
	// fmt.Printf("%d -- Player %d's deck: %s\n", depth, 1, p1)
	// fmt.Printf("%d -- Player %d's deck: %s\n", depth, 2, p2)
	hash := (p1.Hash() << 16) ^ p2.Hash()
	if _, found := previous[hash]; found {
		// fmt.Println("We've seen this before:", hash)
		return -1
	}
	previous[hash] = struct{}{}
	c1 := p1.Dequeue()
	c2 := p2.Dequeue()
	// fmt.Printf("%d -- Player %d plays %d\n", depth, 1, c1)
	// fmt.Printf("%d -- Player %d plays %d\n", depth, 2, c2)
	// when both players have enough cards for a recursive game
	// we play one
	winner := 0
	if c1 <= p1.Size() && c2 <= p2.Size() {
		winner = recursiveGame(p1.Subset(c1), p2.Subset(c2), depth+1)
	} else if c1 > c2 { // just compare the cards
		winner = 1
	} else {
		winner = 2
	}

	if winner == 1 {
		p1.Enqueue(c1)
		p1.Enqueue(c2)
	} else {
		p2.Enqueue(c2)
		p2.Enqueue(c1)
	}
	// fmt.Printf("%d -- Player %d wins\n\n", depth, winner)
	return winner
}

// returns 1 or 2, the number of the winner
func recursiveGame(p1, p2 *IntQueue, depth int) int {
	previous := make(map[int]struct{})

	for p1.Size() != 0 && p2.Size() != 0 {
		if recursiveRound(p1, p2, previous, depth) == -1 {
			return 1
		}
	}
	if p2.Size() == 0 {
		return 1
	}
	return 2
}

func day22b(p1, p2 *IntQueue) int {
	winner := recursiveGame(p1, p2, 1)
	fmt.Println(p1, p2)
	if winner == 1 {
		return score(p1)
	}
	return score(p2)
}

func parse1(lines []string) (int, *IntQueue) {
	count := 0
	q := newIntQueue(100)
	for i, l := range lines[1:] {
		if l == "" {
			count = i + 2
			break
		}
		n, err := strconv.Atoi(l)
		if err != nil {
			log.Fatal(err)
		}
		q.Enqueue(n)
	}
	return count, q
}

func parse(lines []string) (*IntQueue, *IntQueue) {
	n, p1 := parse1(lines)
	_, p2 := parse1(lines[n:])
	return p1, p2
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
	p1, p2 := parse(lines)
	fmt.Println(day22a(p1.Clone(), p2.Clone()))

	fmt.Println(day22b(p1, p2))
}
