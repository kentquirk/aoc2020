package main

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"
)

const (
	tilesize = 10
	nedges   = 4
	lobitix  = 0
	hibitix  = (tilesize - 1)
	lobit    = 1
	hibit    = 1 << hibitix
)

// ID is the tile number
type ID int

// Edge is the type of the constant for edge indices
type Edge int

const (
	top Edge = iota
	right
	bottom
	left
)

// reverses bits 0-9 in an int
func reverse10(n int) int {
	r := 0
	for i := 0; i < tilesize; i++ {
		r = (r << 1) | (n & 1)
		n >>= 1
	}
	return r
}

// Tile represents a single image tile
type Tile struct {
	id        ID
	data      [tilesize]int
	edges     [nedges]int
	matches   map[Edge]ID
	transform string
}

func newTile(id ID, transform string) *Tile {
	return &Tile{
		id:        id,
		transform: transform,
		matches:   make(map[Edge]ID),
	}
}

func (t Tile) edgeCount() int {
	return len(t.matches)
}

func (t Tile) findMatch(edge int) Edge {
	// we only want to look at half the range
	// because we don't need to flip both tiles
	for e := 0; e < len(t.edges)/2; e++ {
		if t.edges[e] == edge {
			return Edge(e)
		}
	}
	return -1
}

func (t Tile) isBorder(edge Edge) bool {
	_, ok := t.matches[edge]
	return !ok
}

func (t *Tile) updateEdges() {
	t.edges[top] = t.data[0]
	t.edges[bottom] = t.data[tilesize-1]
	t.edges[left] = t.vert(hibitix)
	t.edges[right] = t.vert(lobitix)
}

func (t Tile) vert(bitindex int) int {
	mask := 1 << bitindex
	result := 0
	for i := 0; i < tilesize; i++ {
		result = (result << 1) | ((t.data[i] & mask) >> bitindex)
	}
	return result
}

func (t Tile) rotateLeft() *Tile {
	r := newTile(t.id, t.transform+"L")
	for i := 0; i < tilesize; i++ {
		r.data[i] = t.vert(i)
	}
	r.updateEdges()
	// rotate the matches
	if m, ok := t.matches[left]; ok {
		r.matches[bottom] = m
	}
	if m, ok := t.matches[top]; ok {
		r.matches[left] = m
	}
	if m, ok := t.matches[right]; ok {
		r.matches[top] = m
	}
	if m, ok := t.matches[bottom]; ok {
		r.matches[right] = m
	}
	if len(r.matches) != len(t.matches) {
		log.Fatal("Bad match movement!")
	}
	return r
}

func (t Tile) hflip() *Tile {
	r := newTile(t.id, t.transform+"H")
	for i := 0; i < tilesize; i++ {
		r.data[i] = reverse10(t.data[i])
	}
	r.updateEdges()
	// swap left and right
	if m, ok := t.matches[left]; ok {
		r.matches[right] = m
	}
	if m, ok := t.matches[right]; ok {
		r.matches[left] = m
	}
	if m, ok := t.matches[top]; ok {
		r.matches[top] = m
	}
	if m, ok := t.matches[bottom]; ok {
		r.matches[bottom] = m
	}
	return r
}

func (t Tile) vflip() *Tile {
	r := newTile(t.id, t.transform+"V")
	for i := 0; i < tilesize; i++ {
		r.data[i] = t.data[tilesize-i-1]
	}
	r.updateEdges()
	// swap top and bottom
	if m, ok := t.matches[top]; ok {
		r.matches[bottom] = m
	}
	if m, ok := t.matches[bottom]; ok {
		r.matches[top] = m
	}
	if m, ok := t.matches[left]; ok {
		r.matches[left] = m
	}
	if m, ok := t.matches[right]; ok {
		r.matches[right] = m
	}
	return r
}

func parse(d string) *Tile {
	lines := strings.Split(d, "\n")
	pat := regexp.MustCompile("[0-9]+")
	id, _ := strconv.Atoi(pat.FindString(lines[0]))
	t := newTile(ID(id), "")
	rep := strings.NewReplacer(".", "0", "#", "1")
	for i := 0; i < tilesize; i++ {
		bin := rep.Replace(lines[i+1])
		x, err := strconv.ParseInt(bin, 2, 16)
		if err != nil {
			log.Fatal(err)
		}
		t.data[i] = int(x)
	}
	t.updateEdges()
	return t
}

func asText(n int) string {
	s := ""
	for m := hibit; m != 0; m >>= 1 {
		if n&m != 0 {
			s += "#"
		} else {
			s += "."
		}
	}
	return s
}

func (t Tile) String() string {
	sb := &strings.Builder{}
	fmt.Fprintf(sb, "Tile %d: (%s)\n", t.id, t.transform)
	for i := 0; i < tilesize; i++ {
		fmt.Fprintf(sb, "%s\n", asText(t.data[i]))
	}
	fmt.Fprintln(sb, t.edges)
	fmt.Fprintln(sb, t.matches)
	return sb.String()
}

// GetBit returns true if the corresponding bit was set.
// row is a number 0-7 from top to bottom
// col is a number 0-7 from left to right
// both ignore the borders
func (t *Tile) GetBit(row int, col int) bool {
	mask := 1 << (8 - col)
	if (mask & t.data[row+1]) != 0 {
		return true
	}
	return false
}
