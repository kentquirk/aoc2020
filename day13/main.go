package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math/big"
	"os"
	"strconv"
	"strings"
)

// This code swiped from https://rosettacode.org/wiki/Chinese_remainder_theorem#Go
var one = big.NewInt(1)

// crt is "Chinese Remainder Theorem" -- see wikipedia for more
func crt(a, n []*big.Int) (*big.Int, error) {
	p := new(big.Int).Set(n[0])
	for _, n1 := range n[1:] {
		p.Mul(p, n1)
	}
	var x, q, s, z big.Int
	for i, n1 := range n {
		q.Div(p, n1)
		z.GCD(nil, &s, n1, &q)
		if z.Cmp(one) != 0 {
			return nil, fmt.Errorf("%d not coprime", n1)
		}
		x.Add(&x, s.Mul(a[i], s.Mul(&s, &q)))
	}
	return x.Mod(&x, p), nil
}

func day13a(ts int, data string) (int, int) {
	splits := strings.Split(data, ",")
	numbers := make(map[int]struct{})
	for _, s := range splits {
		id, err := strconv.Atoi(s)
		if err == nil {
			numbers[id] = struct{}{}
		}
	}
	bestid := 0
	bestwait := 999999999
	for id := range numbers {
		wait := id - (ts % id)
		if wait < bestwait {
			bestwait = wait
			bestid = id
		}
	}
	return bestid, bestwait
}

func remaindersAt(ts int, numbers map[int]int) map[int]int {
	result := make(map[int]int)
	for id := range numbers {
		result[id] = id - (ts % id)
	}
	return result
}

func day13b(data string) int {
	splits := strings.Split(data, ",")
	n := make([]*big.Int, 0)
	a := make([]*big.Int, 0)
	for ix, s := range splits {
		id, err := strconv.Atoi(s)
		if err == nil {
			n = append(n, big.NewInt(int64(id)))
			remainder := (id - (ix % id)) % id
			a = append(a, big.NewInt(int64(remainder)))
		}
	}
	fmt.Println(n, a)
	result, err := crt(a, n)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(result)
	return 0
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
	ts, _ := strconv.Atoi(lines[0])
	id, wait := day13a(ts, lines[1])
	fmt.Printf("A: best id is %d, best wait is %d, product is %d\n", id, wait, id*wait)
	_ = day13b(lines[1])
}
