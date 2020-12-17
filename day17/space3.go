package main

import "fmt"

// Coord represents an integral coordinate
type Coord struct {
	x int
	y int
	z int
}

func doDeltas() []Coord {
	coords := make([]Coord, 0)
	for x := -1; x <= 1; x++ {
		for y := -1; y <= 1; y++ {
			for z := -1; z <= 1; z++ {
				if x == 0 && y == 0 && z == 0 {
					continue
				}
				coords = append(coords, Coord{x, y, z})
			}
		}
	}
	return coords
}

var neighborDeltas = doDeltas()

const bits20 = 0xFFFFF
const bit21 = 0x100000

func bits21(n int) int64 {
	if n < 0 {
		return bit21 | int64(bits20&(-n))
	}
	return int64(bits20 & n)
}

func toInt(n int64) int {
	if n&bit21 != 0 {
		return int(-(n & bits20))
	}
	return int(n & bits20)
}

// Pack calculates a unique value (provided that no coordinate
// goes past 20 bits)
func (c Coord) Pack() int64 {
	return bits21(c.z)<<42 | bits21(c.y)<<21 | bits21(c.x)
}

// Unpack takes apart a packed value and puts it back into the
// coordinate
func Unpack(n int64) Coord {
	return Coord{
		x: toInt(n),
		y: toInt(n >> 21),
		z: toInt(n >> 42),
	}
}

// Board represents our "energy system"
type Board struct {
	grid map[int64]bool
}

// Active returns the state of a single cell
func (b *Board) Active(c Coord) bool {
	if active, ok := b.grid[c.Pack()]; ok {
		return active
	}
	return false
}

// NeighborActive returns the state of a single cell
func (b *Board) NeighborActive(c Coord, d Coord) bool {
	c.x += d.x
	c.y += d.y
	c.z += d.z
	if active, ok := b.grid[c.Pack()]; ok {
		return active
	}
	return false
}

// CountActive gets the total number of active cells on the board
func (b *Board) CountActive() int {
	total := 0
	for _, state := range b.grid {
		if state {
			total++
		}
	}
	return total
}

// CountNeighbors returns the number of active neighbors for a grid coordinate.
func (b *Board) CountNeighbors(coord Coord) int {
	total := 0
	for _, neighbor := range neighborDeltas {
		if b.NeighborActive(coord, neighbor) {
			total++
		}
	}
	return total
}

// Neighborhood returns a collection of all of the possible
// cells that need to be inspected for a given generation.
func (b *Board) Neighborhood() map[int64]Coord {
	possibles := make(map[int64]Coord)
	for pcoord := range b.grid {
		for _, n := range neighborDeltas {
			coord := Unpack(pcoord)
			coord.x += n.x
			coord.y += n.y
			coord.z += n.z
			possibles[coord.Pack()] = coord
		}
	}
	return possibles
}

// Generation iterates a single generation into a new board.
func (b *Board) Generation() *Board {
	nextBoard := NewBoard()
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

func (b *Board) printActives() {
	for pcoord := range b.grid {
		fmt.Println(Unpack(pcoord))
	}
}

// NewBoard builds a board
func NewBoard() *Board {
	return &Board{
		grid: make(map[int64]bool),
	}
}

// ParseBoard builds a new board from an input text
func ParseBoard(lines []string) *Board {
	board := NewBoard()
	for row, line := range lines {
		for col, ch := range line {
			if ch == '#' {
				coord := Coord{col, row, 0}
				board.grid[coord.Pack()] = true
			}
		}
	}
	return board
}
