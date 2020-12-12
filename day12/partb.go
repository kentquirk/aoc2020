package main

import (
	"fmt"
	"log"
	"math"
	"strconv"
)

type shipB struct {
	x      int
	y      int
	wayptX int
	wayptY int
}

func (s *shipB) move(instr string) {
	fmt.Println(instr, s)
	howToMove := instr[0]
	mag, _ := strconv.Atoi(instr[1:])
	switch howToMove {
	case 'N':
		s.wayptY += mag
	case 'S':
		s.wayptY -= mag
	case 'E':
		s.wayptX += mag
	case 'W':
		s.wayptX -= mag
	case 'R':
		wx := s.wayptX
		wy := s.wayptY
		switch mag {
		case 0:
		case 90:
			wx = s.wayptY
			wy = -s.wayptX
		case 180:
			wx = -s.wayptX
			wy = -s.wayptY
		case 270:
			wx = -s.wayptY
			wy = s.wayptX
		}
		s.wayptX = wx
		s.wayptY = wy
	case 'L':
		wx := s.wayptX
		wy := s.wayptY
		switch mag {
		case 0:
		case 90:
			wx = -s.wayptY
			wy = s.wayptX
		case 180:
			wx = -s.wayptX
			wy = -s.wayptY
		case 270:
			wx = s.wayptY
			wy = -s.wayptX
		}
		s.wayptX = wx
		s.wayptY = wy
	case 'F':
		s.x += s.wayptX * mag
		s.y += s.wayptY * mag
	default:
		log.Fatal("oops")
	}
}

func day12b(lines []string) int {
	s := new(shipB)
	s.wayptX = 10
	s.wayptY = 1
	for _, line := range lines {
		s.move(line)
	}
	fmt.Println(s)
	return int(math.Round(math.Abs(float64(s.x)) + math.Abs(float64(s.y))))
}
