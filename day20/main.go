package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func main() {
	f, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	b, err := ioutil.ReadAll(f)
	if err != nil {
		log.Fatal(err)
	}

	// parse the tiles and create the edgemap,
	// which is a map of the edge values to tiles with them
	textiles := strings.Split(string(b), "\n\n")
	tiles := make(map[ID]*Tile)
	edgemap := make(map[int][]ID)
	for _, t := range textiles {
		tile := parse(t)
		tiles[tile.id] = tile
		for _, e := range tile.edges {
			ids := append(edgemap[e], tile.id)
			edgemap[e] = ids
			r := reverse10(e)
			ids = append(edgemap[r], tile.id)
			edgemap[r] = ids
		}
	}
	// fmt.Println(edgemap)

	// walk the tiles and find all the tiles with 2 shared
	// edges; these are the corners
	// 3 shared edges are puzzle edge pieces, and 4 shared
	// edges are interior pieces
	product := 1
	corners := make([]*Tile, 0)
	for id, tile := range tiles {
		for edge, e := range tile.edges {
			ids := edgemap[e]
			switch len(ids) {
			case 1:
				continue
			case 2:
				other := 1
				if ids[1] == id {
					other = 0
				}
				tile.matches[Edge(edge)] = ids[other]
			case 0:
				ids := edgemap[reverse10(e)]
				if len(ids) == 2 {
					other := 1
					if ids[1] == id {
						other = 0
					}
					tile.matches[Edge(edge)] = ids[other]
				}
			default:
				log.Fatalf("%d tiles in element of edgemap for %d", len(ids), e)
			}
		}
		if tile.edgeCount() == 2 {
			corners = append(corners, tile)
			fmt.Println("corner: ", tile.id)
			product *= int(tile.id)
			// fmt.Println(tile)
		}
	}
	fmt.Println("Part A -- Corner Product: ", product)

	puzzle := newPuzzle()
	// assemble the puzzle; start by constraining the first corner
	piece := matchConstraints(corners[0], []Constraint{{left, 0}, {top, 0}})
	// piece := matchConstraints(tiles[3079], []Constraint{{left, 0}, {top, 0}})
	puzzle.place(piece, 0, 0)

	placed := make(map[ID]struct{})
	placed[piece.id] = struct{}{}

	row := 0
	col := 1
	for {
		if len(placed) == len(tiles) {
			break
		}
		constraints := puzzle.getConstraints(row, col)
		// fmt.Println(constraints)
		foundMatch := false
		for t := range tiles {
			// skip the tiles we've already placed
			if _, ok := placed[t]; ok {
				continue
			}
			piece := matchConstraints(tiles[t], constraints)
			if piece != nil {
				// fmt.Printf("found match for (r%d, c%d) %s\n", row, col, piece)
				puzzle.place(piece, row, col)
				placed[piece.id] = struct{}{}
				if piece.isBorder(right) {
					col = 0
					row++
				} else {
					col++
				}
				foundMatch = true
				break
			}
		}
		if !foundMatch {
			log.Fatalf("no match found at (r%d, c%d)\n", row, col)
		}
	}
	fmt.Println("ANSWER:",
		puzzle.pieces[0][0].id*
			puzzle.pieces[0][puzzle.width-1].id*
			puzzle.pieces[puzzle.height-1][0].id*
			puzzle.pieces[puzzle.height-1][puzzle.width-1].id,
	)

	fmt.Println(puzzle.height, puzzle.width)
	bmp := puzzle.asBits()
	dup := bmp.clone()
	sm := seamonster()
	done := false
outer:
	for flip := 0; flip < 2; flip++ {
		for rotation := 0; rotation < 4; rotation++ {
			dup = bmp.clone()
			for r := 0; r < bmp.height; r++ {
				for c := 0; c < bmp.width; c++ {
					if bmp.matchSeamonsterAt(sm, r, c) {
						fmt.Printf("flip=%d, rotation=%d\n", flip, rotation)
						fmt.Printf("seamonster at (r%d, c%d)\n", r, c)
						dup.eraseSeamonsterAt(sm, r, c)
						done = true
					}
				}
			}
			if done {
				break outer
			}
			bmp = bmp.rotateLeft()
		}
		bmp = bmp.vflip()
	}

	fmt.Println(bmp.countBits(), dup.countBits())
}
