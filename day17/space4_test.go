package main

import (
	"testing"
)

func Test_toInt4(t *testing.T) {
	type args struct {
		n Packed4
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"a", args{Packed4(4)}, 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := toInt4(tt.args.n); got != tt.want {
				t.Errorf("toInt4() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCoord4_Pack(t *testing.T) {
	type fields struct {
		x int
		y int
		z int
		w int
	}
	tests := []struct {
		name   string
		fields fields
		want   Packed4
	}{
		{"a", fields{1, 2, 3, 4}, Packed4(0x0001000200030004)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := Coord4{
				x: tt.fields.x,
				y: tt.fields.y,
				z: tt.fields.z,
				w: tt.fields.w,
			}
			if got := c.Pack(); got != tt.want {
				t.Errorf("Coord4.Pack() = %v, want %v", got, tt.want)
			}
		})
	}
}
