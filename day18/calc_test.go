package main

import (
	"testing"
)

func TestCalculator(t *testing.T) {
	tests := []struct {
		name    string
		s       string
		want    int
		wantErr bool
	}{
		{"a", "3*5", 15, false},
		{"b", "12*25", 300, false},
		{"c", "12+25", 37, false},
		{"d", "12*25* 7 ", 2100, false},
		{"e", "12 25* 7 ", 12, true},
		{"f", "1 + (2*3) ", 7, false},
		{"g", "5 + (8 * 3 + 9 + 3 * 4 * 3)", 437, false},
		{"h", "5 * 9 * (7 * 3 * 3 + 9 * 3 + (8 + 6 * 4))", 12240, false},
		{"i", "(1*2) + (2*3)", 8, false},
		{"i", "((1*2) + (2*3))", 8, false},
		{"j", "((2 + 4 * 9) * (6 + 9 * 8 + 6) + 6) + 2 + 4 * 2", 13632, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			calc := &AoCCalc{Buffer: tt.s}
			calc.Init()
			calc.Expression.Init(tt.s)
			if err := calc.Parse(); (err != nil) != tt.wantErr {
				t.Errorf("Parse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			calc.Execute()
			got := calc.Evaluate()
			if got != tt.want {
				t.Errorf("evaluate() = %v, want %v", got, tt.want)
			}
		})
	}
}
