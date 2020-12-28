package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"os"
	"regexp"
	"strings"
)

type hextile struct {
	x int16
	y int16
	z int16
}

func (h hextile) Hash() int64 {
	const mask16 = 0xFFFF
	return (int64(h.x)&mask16)<<32 |
		(int64(h.y)&mask16)<<16 |
		(int64(h.z) & mask16)
}

func (h *hextile) add(rhs hextile) {
	h.x += rhs.x
	h.y += rhs.y
	h.z += rhs.z
}

var coords = map[string]hextile{
	"nw": {0, +1, -1},
	"se": {0, -1, +1},
	"ne": {+1, 0, -1},
	"sw": {-1, 0, +1},
	"e":  {+1, -1, 0},
	"w":  {-1, +1, 0},
}

// Floor represents our hex-tiled floor
type Floor map[int64]hextile

func makeFloor() Floor {
	return make(map[int64]hextile)
}

func (f Floor) xExtents() (int16, int16) {
	var xmin int16 = math.MaxInt16
	var xmax int16 = math.MinInt16
	for _, tile := range f {
		if xmin > tile.x {
			xmin = tile.x
		}
		if xmax > tile.x {
			xmax = tile.x
		}
	}
	return xmin, xmax
}

func (f Floor) yExtents() (int16, int16) {
	var ymin int16 = math.MaxInt16
	var ymax int16 = math.MinInt16
	for _, tile := range f {
		if ymin > tile.y {
			ymin = tile.y
		}
		if ymax > tile.y {
			ymax = tile.y
		}
	}
	return ymin, ymax
}

func (f Floor) zExtents() (int16, int16) {
	var zmin int16 = math.MaxInt16
	var zmax int16 = math.MinInt16
	for _, tile := range f {
		if zmin > tile.z {
			zmin = tile.z
		}
		if zmax > tile.z {
			zmax = tile.z
		}
	}
	return zmin, zmax
}

func (h hextile) neighbors() []hextile {
	neighbors := make([]hextile, 0)
	for _, offset := range coords {
		t := h
		t.add(offset)
		neighbors = append(neighbors, t)
	}
	return neighbors
}

func (f Floor) neighborCount(h hextile) int {
	count := 0
	for _, tile := range h.neighbors() {
		if _, ok := f[tile.Hash()]; ok {
			count++
		}
	}
	return count
}

func (f Floor) generation() Floor {
	newfloor := makeFloor()
	for _, tile := range f {
		// look at a known-black tile
		n := f.neighborCount(tile)
		if n == 1 || n == 2 {
			newfloor[tile.Hash()] = tile
		}
		for _, neighbor := range tile.neighbors() {
			if _, ok := f[neighbor.Hash()]; !ok {
				if f.neighborCount(neighbor) == 2 {
					newfloor[neighbor.Hash()] = neighbor
				}
			}
		}
	}
	return newfloor
}

func parseFloor(lines []string) Floor {
	floor := makeFloor()
	// input looks like: sesenwnenenewseeswwswswwnenewsewsw
	pat := regexp.MustCompile("ne|se|nw|sw|e|w")
	for _, line := range lines {
		moves := pat.FindAllString(line, -1)
		// fmt.Println(moves)
		pos := hextile{}
		for _, move := range moves {
			// fmt.Printf("pos: %v move: %s add: %v ", pos, move, coords[move])
			pos.add(coords[move])
			// fmt.Printf("newpos: %v\n", pos)
		}
		if _, ok := floor[pos.Hash()]; ok {
			// fmt.Printf("Flipping %v to white\n", pos)
			delete(floor, pos.Hash())
		} else {
			// fmt.Printf("Flipping %v to black\n", pos)
			floor[pos.Hash()] = pos
		}
	}
	return floor
}

func day24a(lines []string) int {
	return len(parseFloor(lines))
}

func day24b(lines []string, generations int) int {
	floor := parseFloor(lines)
	for i := 0; i < generations; i++ {
		floor = floor.generation()
	}
	return len(floor)
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
	fmt.Println(day24a(lines))
	fmt.Println(day24b(lines, 100))
}
