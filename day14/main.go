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

func parseMask(mask string) (int64, int64) {
	var zerosMask, onesMask int64
	zerosMask = 0xffffffff
	bit := int64(1 << 35)
	for _, c := range mask {
		switch c {
		case '0':
			zerosMask &= ^bit
		case '1':
			onesMask |= bit
		case 'X':
		}
		bit >>= 1
	}
	return zerosMask, onesMask
}

func doAll(lines []string) int64 {
	pat := regexp.MustCompile("(mask|mem)([^= ]*)[ =]+(.+)")
	var zerosMask, onesMask int64
	memory := make(map[int]int64)
	for _, line := range lines {
		groups := pat.FindStringSubmatch(line)
		if len(groups) == 0 {
			log.Fatalf("no groups found on %s\n", line)
		}
		switch groups[1] {
		case "mask":
			zerosMask, onesMask = parseMask(groups[3])
		case "mem":
			addr, _ := strconv.Atoi(groups[2][1 : len(groups[2])-1])
			writeVal, _ := strconv.ParseInt(groups[3], 10, 64)
			writeVal &= zerosMask
			writeVal |= onesMask
			memory[addr] = writeVal
		default:
			log.Fatalf("unparseable command %s", line)
		}
	}
	var total int64
	for _, v := range memory {
		total += v
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
	fmt.Println(doAll(lines))
}
