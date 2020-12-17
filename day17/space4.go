package main

import "fmt"

// Coord4 represents an integral coordinate
type Coord4 struct {
	x int
	y int
	z int
	w int
}

func doDeltas4() []Coord4 {
	coords := make([]Coord4, 0)
	for x := -1; x <= 1; x++ {
		for y := -1; y <= 1; y++ {
			for z := -1; z <= 1; z++ {
				for w := -1; w <= 1; w++ {
					if x == 0 && y == 0 && z == 0 && w == 0 {
						continue
					}
					coords = append(coords, Coord4{x, y, z, w})
				}
			}
		}
	}
	return coords
}

var neighborDeltas4 = doDeltas4()

const bits15 = 0x7FFF
const bit16 = 0x8000

// Packed4 defines the type of a packed coordinate
type Packed4 uint64

func bits16(n int) Packed4 {
	if n < 0 {
		return bit16 | Packed4(bits15&(-n))
	}
	return Packed4(bits15 & n)
}

func toInt4(n Packed4) int {
	if n&bit16 != 0 {
		return int(-(n & bits15))
	}
	return int(n & bits15)
}

// Pack calculates a unique value (provided that no coordinate
// goes past 15 bits)
func (c Coord4) Pack() Packed4 {
	return bits16(c.w)<<48 | bits16(c.z)<<32 | bits16(c.y)<<16 | bits16(c.x)
}

// Unpack4 takes apart a packed value and puts it back into the
// coordinate
func (p Packed4) Unpack4() Coord4 {
	return Coord4{
		x: toInt4(p),
		y: toInt4(p >> 16),
		z: toInt4(p >> 32),
		w: toInt4(p >> 48),
	}
}

// Board4 represents our "energy system"
type Board4 struct {
	grid map[Packed4]bool
}

// Active returns the state of a single cell
func (b *Board4) Active(c Coord4) bool {
	if active, ok := b.grid[c.Pack()]; ok {
		return active
	}
	return false
}

// NeighborActive returns the state of a single cell
func (b *Board4) NeighborActive(c Coord4, d Coord4) bool {
	c.x += d.x
	c.y += d.y
	c.z += d.z
	c.w += d.w
	if active, ok := b.grid[c.Pack()]; ok {
		return active
	}
	return false
}

// CountActive gets the total number of active cells on the board
func (b *Board4) CountActive() int {
	total := 0
	for _, state := range b.grid {
		if state {
			total++
		}
	}
	return total
}

// CountNeighbors returns the number of active neighbors for a grid coordinate.
func (b *Board4) CountNeighbors(coord Coord4) int {
	total := 0
	for _, neighbor := range neighborDeltas4 {
		if b.NeighborActive(coord, neighbor) {
			total++
		}
	}
	return total
}

// Neighborhood returns a collection of all of the possible
// cells that need to be inspected for a given generation.
func (b *Board4) Neighborhood() map[Packed4]Coord4 {
	possibles := make(map[Packed4]Coord4)
	for pcoord := range b.grid {
		for _, n := range neighborDeltas4 {
			coord := pcoord.Unpack4()
			coord.x += n.x
			coord.y += n.y
			coord.z += n.z
			coord.w += n.w
			possibles[coord.Pack()] = coord
		}
	}
	return possibles
}

// Generation iterates a single generation into a new board.
func (b *Board4) Generation() *Board4 {
	nextBoard := NewBoard4()
	for pcoord, coord := range b.Neighborhood() {
		neighbors := b.CountNeighbors(coord)
		if b.Active(coord) {
			if neighbors == 2 || neighbors == 3 {
				nextBoard.grid[pcoord] = true
			}
		} else {
			if neighbors == 3 {
				nextBoard.grid[pcoord] = true
			}
		}
	}
	return nextBoard
}

func (b *Board4) printActives() {
	for pcoord := range b.grid {
		fmt.Println(pcoord.Unpack4())
	}
}

// NewBoard4 builds a board
func NewBoard4() *Board4 {
	return &Board4{
		grid: make(map[Packed4]bool),
	}
}

// ParseBoard4 builds a new board from an input text
func ParseBoard4(lines []string) *Board4 {
	board := NewBoard4()
	for row, line := range lines {
		for col, ch := range line {
			if ch == '#' {
				coord := Coord4{col, row, 0, 0}
				board.grid[coord.Pack()] = true
			}
		}
	}
	return board
}
