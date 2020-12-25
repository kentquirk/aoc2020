package main

import (
	"testing"
)

func Test_reverse10(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want int
	}{
		{"a", 1, 512},
		{"b", 512, 1},
		{"c", 1023, 1023},
		{"d", 0, 0},
		{"e", 0x155, 0x2AA},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverse10(tt.n); got != tt.want {
				t.Errorf("reverse10() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTile_vert(t *testing.T) {
	tests := []struct {
		name     string
		data     [tilesize]int
		bitindex int
		want     int
	}{
		{"a", [10]int{0, 0x3FF, 0, 0, 0, 0, 0, 0, 0, 0}, 0, 256},
		{"b", [10]int{0, 0x3FF, 0, 0, 0, 0, 0, 0, 0, 0}, 5, 256},
		{"c", [10]int{0, 0x3FF, 0, 0, 0, 0, 0, 0, 0, 0}, 9, 256},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tile := newTile(123, "")
			tile.data = tt.data
			if got := tile.vert(tt.bitindex); got != tt.want {
				t.Errorf("Tile.vert() = %v, want %v", got, tt.want)
			}
		})
	}
}
