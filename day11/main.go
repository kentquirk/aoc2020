package main

import (
	"crypto/sha256"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

// Cell defines a single square in our grid; it
// may or may not contain a seat, and if it does,
// it may or may not be occupied. A cell without
// a seat can never be occuped.
type Cell struct {
	HasSeat    bool
	IsOccupied bool
}

// Board is a collection of cells, and the width
// and height of the board are padded on all sides
// in the grid so that we can index into the grid
// and calculate neighborhood count without conditionals.
type Board struct {
	width  int
	height int
	grid   [][]Cell
	gens   int
}

func parseLine(s string) []Cell {
	row := make([]Cell, len(s)+2)
	for i, ch := range s {
		switch ch {
		case 'L':
			row[i+1].HasSeat = true
		case '.':
		default:
		}
	}
	return row
}

// NewBoard constructs an empty board.
func NewBoard(width int, height int) *Board {
	grid := make([][]Cell, height+2)
	for i := range grid {
		grid[i] = make([]Cell, width+2)
	}
	return &Board{
		width:  width,
		height: height,
		grid:   grid,
	}
}

// ParseBoard constructs a board from a slice of strings
// that describe it.
func ParseBoard(lines []string) *Board {
	board := NewBoard(len(lines[0]), len(lines))
	for i, line := range lines {
		board.grid[i+1] = parseLine(line)
	}
	return board
}

// ToString does what it says on the tin.
func (b Board) ToString() string {
	s := strings.Builder{}
	for row := 1; row <= b.height; row++ {
		for col := 1; col <= b.width; col++ {
			cell := b.grid[row][col]
			if cell.HasSeat {
				if cell.IsOccupied {
					s.WriteRune('#')
				} else {
					s.WriteRune('L')
				}
			} else {
				s.WriteRune('.')
			}
		}
		s.WriteString("\n")
	}
	return s.String()
}

// Hash generates a unique string per board
func (b Board) Hash() string {
	s := b.ToString()
	sum := sha256.Sum256([]byte(s))
	return fmt.Sprintf("%x", sum)
}

// Clone copies a board to a new entity
func (b Board) Clone() *Board {
	clone := NewBoard(b.width, b.height)
	for row := 1; row <= b.height; row++ {
		copy(clone.grid[row], b.grid[row])
	}
	clone.gens = b.gens + 1
	return clone
}

// NeighborCount returns the number of occupied neighbors that are seats
func (b Board) NeighborCount(row int, col int) int {
	count := 0
	for r := row - 1; r <= row+1; r++ {
		for c := col - 1; c <= col+1; c++ {
			if r == row && c == col {
				continue
			}
			if b.grid[r][c].IsOccupied {
				count++
			}
		}
	}
	return count
}

// Generation runs one cycle of our automaton
// It does not modify the receiver board but creates
// a new one.
func (b *Board) Generation() *Board {
	next := b.Clone()
	for row := 1; row <= b.height; row++ {
		for col := 1; col <= b.width; col++ {
			seat := b.grid[row][col]
			neighbors := b.NeighborCount(row, col)
			if !seat.HasSeat {
				continue
			}
			if !seat.IsOccupied {
				if neighbors == 0 {
					next.grid[row][col].IsOccupied = true
				}
			} else {
				if neighbors >= 4 {
					next.grid[row][col].IsOccupied = false
				}
			}
		}
	}
	return next
}

// OccupiedSeats returns the number of seats that are occupied.
func (b Board) OccupiedSeats() int {
	count := 0
	for row := 1; row <= b.height; row++ {
		for col := 1; col <= b.width; col++ {
			if b.grid[row][col].IsOccupied {
				count++
			}
		}
	}
	return count
}

// Run calculates generations until the board is stable.
func (b *Board) Run() *Board {
	lasthash := ""
	for lasthash != b.Hash() {
		lasthash = b.Hash()
		b = b.Generation()
	}
	return b
}

func day11a(lines []string) (int, int) {
	b := ParseBoard(lines)
	b = b.Run()
	return b.gens, b.OccupiedSeats()
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
	gens, seats := day11a(lines)
	fmt.Printf("%d seats were occupied after %d generations", seats, gens)
}
