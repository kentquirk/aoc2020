package main

import (
	"regexp"
	"strings"

	"github.com/Workiva/go-datastructures/bitarray"
)

// Bitmap represents a way to handle a square array of bits that's
// too big to fit into a simple int64
type Bitmap struct {
	bits   []bitarray.BitArray
	height int
	width  int
}

func createBitmap(height int, width int) Bitmap {
	bm := Bitmap{
		bits:   make([]bitarray.BitArray, height),
		height: height,
		width:  width,
	}

	for r := 0; r < height; r++ {
		bm.bits[r] = bitarray.NewBitArray(uint64(width))
	}
	return bm
}

func (b Bitmap) clone() Bitmap {
	result := createBitmap(b.height, b.width)
	for r := 0; r < b.height; r++ {
		for c := 0; c < b.width; c++ {
			if bit, _ := b.bits[r].GetBit(uint64(c)); bit {
				result.bits[r].SetBit(uint64(c))
			}
		}
	}
	return result
}

func (b Bitmap) rotateLeft() Bitmap {
	result := createBitmap(b.height, b.width)
	for r := 0; r < b.height; r++ {
		for c := 0; c < b.width; c++ {
			if bit, _ := b.bits[r].GetBit(uint64(c)); bit {
				result.bits[b.width-c-1].SetBit(uint64(r))
			}
		}
	}
	return result
}

func (b Bitmap) vflip() Bitmap {
	result := createBitmap(b.height, b.width)
	for r := 0; r < b.height; r++ {
		for c := 0; c < b.width; c++ {
			if bit, _ := b.bits[r].GetBit(uint64(c)); bit {
				result.bits[b.height-r-1].SetBit(uint64(c))
			}
		}
	}
	return result
}

func seamonster() [][]int {
	data := `
                  #
#    ##    ##    ###
 #  #  #  #  #  #
 `
	splits := strings.Split(data, "\n")
	row := 0
	pat := regexp.MustCompile("#")
	monster := make([][]int, 0)
	for _, s := range splits {
		finds := pat.FindAllStringIndex(s, -1)
		if finds != nil {
			for _, f := range finds {
				monster = append(monster, []int{row, f[0]})
			}
			row++
		}
	}
	return monster
}

func (b Bitmap) matchSeamonsterAt(monster [][]int, row int, col int) bool {
	for i := range monster {
		r := monster[i][0] + row
		c := monster[i][1] + col

		if r >= b.height || c >= b.width {
			return false
		}
		if ok, _ := b.bits[r].GetBit(uint64(c)); !ok {
			return false
		}
	}
	return true
}

func (b *Bitmap) eraseSeamonsterAt(monster [][]int, row int, col int) {
	for i := range monster {
		r := monster[i][0] + row
		c := monster[i][1] + col
		b.bits[r].ClearBit(uint64(c))
	}
}

func (b Bitmap) countBits() int {
	total := 0
	for r := 0; r < b.height; r++ {
		for c := 0; c < b.width; c++ {
			if bit, _ := b.bits[r].GetBit(uint64(c)); bit {
				total++
			}
		}
	}
	return total
}
