package main

import (
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

// NewBoard constructs a board.
func NewBoard(lines []string) *Board {
	width := len(lines[0])
	height := len(lines)
	grid := make([][]Cell, height+2)
	for i, line := range lines {
		grid[i+1] = parseLine(line)
	}
	grid[0] = make([]Cell, width+2)
	grid[height+1] = make([]Cell, width+2)
	return &Board{
		width:  width,
		height: height,
		grid:   grid,
	}
}

// Print does what it says on the tin.
func (b Board) Print() {
	for row := 1; row <= b.height; row++ {
		s := ""
		for col := 1; col <= b.width; col++ {
			cell := b.grid[row][col]
			if cell.HasSeat {
				if cell.IsOccupied {
					s += "#"
				} else {
					s += "L"
				}
			} else {
				s += "."
			}
		}
		fmt.Println(s)
	}
}

func main() {
	f, err := os.Open("./inputsample.txt")
	if err != nil {
		log.Fatal(err)
	}
	b, err := ioutil.ReadAll(f)
	if err != nil {
		log.Fatal(err)
	}
	lines := strings.Split(string(b), "\n")
	board := NewBoard(lines)
	board.Print()
}
