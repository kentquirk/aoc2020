package main

import (
	"log"
	"math"
	"strconv"
)

type shipA struct {
	x           int
	y           int
	orientation int
}

func (s *shipA) move(instr string) {
	// fmt.Println(instr, s)
	howToMove := instr[0]
	mag, _ := strconv.Atoi(instr[1:])
	switch howToMove {
	case 'N':
		s.y += mag
	case 'S':
		s.y -= mag
	case 'E':
		s.x += mag
	case 'W':
		s.x -= mag
	case 'L':
		s.orientation = (s.orientation + mag) % 360
	case 'R':
		s.orientation = (s.orientation - mag) % 360
	case 'F':
		dx := int(math.Round(math.Cos(float64(s.orientation) * math.Pi / 180.0)))
		dy := int(math.Round(math.Sin(float64(s.orientation) * math.Pi / 180.0)))
		s.x += dx * mag
		s.y += dy * mag
	default:
		log.Fatal("oops")
	}
}

func day12a(lines []string) int {
	s := new(shipA)
	for _, line := range lines {
		s.move(line)
	}
	// fmt.Println(s)
	return int(math.Round(math.Abs(float64(s.x)) + math.Abs(float64(s.y))))
}
