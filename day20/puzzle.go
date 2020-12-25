package main

// Constraint represents a single constraint on a puzzle
// piece. used by matchConstraints.
// A value of zero matches an empty edge.
type Constraint struct {
	edge  Edge
	value int
}

// tries to rotate / flip the tile to match the given
// constraints. If possible, returns the tile, else nil.
func matchConstraints(tile *Tile, cs []Constraint) *Tile {
	for i := 0; i < 2; i++ {
		for j := 0; j < 4; j++ {
			found := true
			for _, c := range cs {
				if c.value == 0 {
					// if constraint wants a 0 value, then
					// we want the tile to be a border on the
					// specified side
					if !tile.isBorder(c.edge) {
						found = false
						break
					}
				} else {
					// if constraint is nonzero, then
					// we need to match the constraint
					if tile.edges[c.edge] != c.value {
						found = false
						break
					}
				}
			}
			if found {
				return tile
			}
			tile = tile.rotateLeft()
		}
		tile = tile.vflip()
	}
	return nil
}

// Puzzle represents a collection of pieces in the right order
type Puzzle struct {
	pieces [][]*Tile
	width  int
	height int
}

func newPuzzle() *Puzzle {
	return &Puzzle{
		pieces: make([][]*Tile, 0),
		width:  0,
		height: 0,
	}
}

func (p *Puzzle) ensureSize(row int, col int) {
	// make sure the puzzle is rectangular and at least
	// as tall as row and as wide as col
	for r := 0; r <= row; r++ {
		if r >= len(p.pieces) {
			p.pieces = append(p.pieces, make([]*Tile, col))
		}
		for c := len(p.pieces[r]); c <= col; c++ {
			p.pieces[r] = append(p.pieces[r], nil)
		}
	}
	p.height = len(p.pieces)
	p.width = len(p.pieces[0])
}

func (p *Puzzle) place(tile *Tile, row int, col int) {
	p.ensureSize(row, col)
	p.pieces[row][col] = tile
}

// getConstraints returns a set of constraints for a given grid
// based on the current state of the puzzle. It assumes we're filling
// out the puzzle top to bottom, left to right.
func (p *Puzzle) getConstraints(row int, col int) []Constraint {
	p.ensureSize(row, col)
	constraints := make([]Constraint, 0)
	value := 0
	if row > 0 && p.pieces[row-1][col] != nil {
		value = p.pieces[row-1][col].edges[bottom]
	}
	constraints = append(constraints, Constraint{top, value})

	value = 0
	if col > 0 && p.pieces[row][col-1] != nil {
		value = p.pieces[row][col-1].edges[right]
	}
	constraints = append(constraints, Constraint{left, value})
	return constraints
}

func (p *Puzzle) asBits() Bitmap {
	result := createBitmap(p.height*8, p.width*8)
	for tilerow := 0; tilerow < p.height; tilerow++ {
		for bitrow := 0; bitrow < 8; bitrow++ {
			row := tilerow*8 + bitrow
			for tilecol := 0; tilecol < p.width; tilecol++ {
				for bitcol := 0; bitcol < 8; bitcol++ {
					col := tilecol*8 + bitcol
					if p.pieces[tilerow][tilecol].GetBit(bitrow, bitcol) {
						result.bits[row].SetBit(uint64(col))
					}
				}
			}
		}
	}
	return Bitmap(result)
}
