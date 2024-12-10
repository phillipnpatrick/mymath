package basicmath

import (
	"testing"
)

func TestFraction_Add(t *testing.T) {
	tests := []struct {
		name       string
		initial    *Fraction
		others     []*Fraction
		wantResult *Fraction
	}{
		{
			name:       "Add two fractions",
			initial:    NewFraction(1, 2),
			others:     []*Fraction{NewFraction(1, 3)},
			wantResult: NewFraction(5, 6), // 1/2 + 1/3 = 5/6
		},
		{
			name:       "test01",
			initial:    NewFraction(1, 2),
			others:     []*Fraction{NewFraction(2, 3)},
			wantResult: NewFraction(7, 6),
		},
		{
			name:       "test02",
			initial:    NewFraction(8, 12),
			others:     []*Fraction{NewFraction(8, 11)},
			wantResult: NewFraction(46, 33),
		},
		{
			name:       "test03",
			initial:    NewFraction(2, 7),
			others:     []*Fraction{NewFraction(6, 10)},
			wantResult: NewFraction(31, 35),
		},
		{
			name:       "test04",
			initial:    NewFraction(1, 6),
			others:     []*Fraction{NewFraction(6, 11)},
			wantResult: NewFraction(47, 66),
		},
		{
			name:       "test05",
			initial:    NewFraction(5, 9),
			others:     []*Fraction{NewFraction(1, 2)},
			wantResult: NewFraction(19, 18),
		},
		{
			name:       "test06",
			initial:    NewFraction(9, 12),
			others:     []*Fraction{NewFraction(2, 12)},
			wantResult: NewFraction(11, 12),
		},
		{
			name:       "test07",
			initial:    NewFraction(5, 12),
			others:     []*Fraction{NewFraction(4, 12)},
			wantResult: NewFraction(3, 4),
		},
		{
			name:       "test08",
			initial:    NewFraction(2, 7),
			others:     []*Fraction{NewFraction(1, 4)},
			wantResult: NewFraction(15, 28),
		},
		{
			name:       "test09",
			initial:    NewFraction(1, 4),
			others:     []*Fraction{NewFraction(6, 8)},
			wantResult: NewInteger(1),
		},
		{
			name:       "test10",
			initial:    NewFraction(4, 10),
			others:     []*Fraction{NewFraction(4, 5)},
			wantResult: NewFraction(6, 5),
		},
		{
			name:       "test11",
			initial:    NewFraction(1, 2),
			others:     []*Fraction{NewFraction(8, 11)},
			wantResult: NewFraction(27, 22),
		},
		{
			name:       "test12",
			initial:    NewFraction(1, 11),
			others:     []*Fraction{NewFraction(2, 12)},
			wantResult: NewFraction(17, 66),
		},
		{
			name:       "test13",
			initial:    NewFraction(2, 12),
			others:     []*Fraction{NewFraction(2, 4)},
			wantResult: NewFraction(2, 3),
		},
		{
			name:       "test14",
			initial:    NewFraction(3, 5),
			others:     []*Fraction{NewFraction(3, 8)},
			wantResult: NewFraction(39, 40),
		},
		{
			name:       "test15",
			initial:    NewFraction(6, 9),
			others:     []*Fraction{NewFraction(1, 2)},
			wantResult: NewFraction(7, 6),
		},
		{
			name:       "test16",
			initial:    NewFraction(1, 2),
			others:     []*Fraction{NewFraction(1, 3), NewFraction(1, 4)},
			wantResult: NewFraction(13, 12),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.initial.Add(tt.others...)
			if got == nil && tt.wantResult == nil {
				// Both are nil, test passed
				return
			}
			if got == nil || tt.wantResult == nil || got.Numerator() != tt.wantResult.Numerator() || got.Denominator() != tt.wantResult.Denominator() {
				t.Errorf("Fraction.Add() = %v, want %v", got, tt.wantResult)
			}
		})
	}
}

func TestFraction_Subtract(t *testing.T) {
	tests := []struct {
		name       string
		initial    *Fraction
		others     []*Fraction
		wantResult *Fraction
	}{
		{
			name:       "test01",
			initial:    NewFraction(2, 9),
			others:     []*Fraction{NewFraction(1, 6)},
			wantResult: NewFraction(1, 18),
		},
		{
			name:       "test02",
			initial:    NewFraction(16, 25),
			others:     []*Fraction{NewFraction(2, 4)},
			wantResult: NewFraction(7, 50),
		},
		{
			name:       "test03",
			initial:    NewFraction(5, 7),
			others:     []*Fraction{NewFraction(1, 3)},
			wantResult: NewFraction(8, 21),
		},
		{
			name:       "test04",
			initial:    NewFraction(5, 7),
			others:     []*Fraction{NewFraction(2, 5)},
			wantResult: NewFraction(11, 35),
		},
		{
			name:       "test05",
			initial:    NewFraction(5, 9),
			others:     []*Fraction{NewFraction(1, 3)},
			wantResult: NewFraction(2, 9),
		},
		{
			name:       "test06",
			initial:    NewFraction(6, 10),
			others:     []*Fraction{NewFraction(1, 8)},
			wantResult: NewFraction(19, 40),
		},
		{
			name:       "test07",
			initial:    NewFraction(9, 16),
			others:     []*Fraction{NewFraction(1, 4)},
			wantResult: NewFraction(5, 16),
		},
		{
			name:       "test08",
			initial:    NewFraction(3, 4),
			others:     []*Fraction{NewFraction(2, 6)},
			wantResult: NewFraction(5, 12),
		},
		{
			name:       "test09",
			initial:    NewFraction(2, 3),
			others:     []*Fraction{NewFraction(1, 8)},
			wantResult: NewFraction(13, 24),
		},
		{
			name:       "test10",
			initial:    NewFraction(8, 20),
			others:     []*Fraction{NewFraction(1, 4)},
			wantResult: NewFraction(3, 20),
		},
		{
			name:       "test11",
			initial:    NewFraction(4, 7),
			others:     []*Fraction{NewFraction(4, 8)},
			wantResult: NewFraction(1, 14),
		},
		{
			name:       "test12",
			initial:    NewFraction(5, 8),
			others:     []*Fraction{NewFraction(2, 6)},
			wantResult: NewFraction(7, 24),
		},
		{
			name:       "test13",
			initial:    NewFraction(1, 2),
			others:     []*Fraction{NewFraction(1, 3), NewFraction(1, 4)},
			wantResult: NewFraction(-1, 12),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.initial.Subtract(tt.others...)
			if got == nil && tt.wantResult == nil {
				// Both are nil, test passed
				return
			}
			if got == nil || tt.wantResult == nil || got.Numerator() != tt.wantResult.Numerator() || got.Denominator() != tt.wantResult.Denominator() {
				t.Errorf("Fraction.Subtract() = %v, want %v", got, tt.wantResult)
			}
		})
	}
}

func TestFraction_Multiply(t *testing.T) {
	tests := []struct {
		name       string
		initial    *Fraction
		others     []*Fraction
		wantResult *Fraction
	}{
		{
			name:       "test01",
			initial:    NewFraction(6, 12),
			others:     []*Fraction{NewFraction(2, 10)},
			wantResult: NewFraction(1, 10),
		},
		{
			name:       "test02",
			initial:    NewFraction(1, 16),
			others:     []*Fraction{NewFraction(7, 21)},
			wantResult: NewFraction(1, 48),
		},
		{
			name:       "test03",
			initial:    NewFraction(8, 9),
			others:     []*Fraction{NewFraction(1, 2)},
			wantResult: NewFraction(4, 9),
		},
		{
			name:       "test04",
			initial:    NewFraction(11, 20),
			others:     []*Fraction{NewFraction(4, 14)},
			wantResult: NewFraction(11, 70),
		},
		{
			name:       "test05",
			initial:    NewFraction(4, 10),
			others:     []*Fraction{NewFraction(1, 6)},
			wantResult: NewFraction(1, 15),
		},
		{
			name:       "test06",
			initial:    NewFraction(2, 5),
			others:     []*Fraction{NewFraction(1, 4)},
			wantResult: NewFraction(1, 10),
		},
		{
			name:       "test07",
			initial:    NewFraction(2, 3),
			others:     []*Fraction{NewFraction(2, 10)},
			wantResult: NewFraction(2, 15),
		},
		{
			name:       "test08",
			initial:    NewFraction(8, 10),
			others:     []*Fraction{NewFraction(4, 7)},
			wantResult: NewFraction(16, 35),
		},
		{
			name:       "test09",
			initial:    NewFraction(8, 12),
			others:     []*Fraction{NewFraction(1, 5)},
			wantResult: NewFraction(2, 15),
		},
		{
			name:       "test10",
			initial:    NewFraction(5, 7),
			others:     []*Fraction{NewFraction(1, 3)},
			wantResult: NewFraction(5, 21),
		},
		{
			name:       "test11",
			initial:    NewFraction(4, 5),
			others:     []*Fraction{NewFraction(3, 8)},
			wantResult: NewFraction(3, 10),
		},
		{
			name:       "test12",
			initial:    NewFraction(1, 3),
			others:     []*Fraction{NewFraction(2, 6)},
			wantResult: NewFraction(1, 9),
		},
		{
			name:       "test13",
			initial:    NewFraction(1, 2),
			others:     []*Fraction{NewFraction(1, 3), NewFraction(1, 4)},
			wantResult: NewFraction(1, 24),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.initial.Multiply(tt.others...)
			if got == nil && tt.wantResult == nil {
				// Both are nil, test passed
				return
			}
			if got == nil || tt.wantResult == nil || got.Numerator() != tt.wantResult.Numerator() || got.Denominator() != tt.wantResult.Denominator() {
				t.Errorf("Fraction.Multiply() = %v, want %v", got, tt.wantResult)
			}
		})
	}
}

func TestFraction_Divide(t *testing.T) {
	tests := []struct {
		name       string
		initial    *Fraction
		others     []*Fraction
		wantResult *Fraction
	}{
		{
			name:       "test01",
			initial:    NewFraction(1, 4),
			others:     []*Fraction{NewFraction(9, 10)},
			wantResult: NewFraction(5, 18),
		},
		{
			name:       "test02",
			initial:    NewFraction(5, 9),
			others:     []*Fraction{NewFraction(1, 2)},
			wantResult: NewFraction(10, 9),
		},
		{
			name:       "test03",
			initial:    NewFraction(1, 3),
			others:     []*Fraction{NewFraction(6, 9)},
			wantResult: NewFraction(1, 2),
		},
		{
			name:       "test04",
			initial:    NewFraction(8, 10),
			others:     []*Fraction{NewFraction(2, 5)},
			wantResult: NewInteger(2),
		},
		{
			name:       "test05",
			initial:    NewFraction(3, 8),
			others:     []*Fraction{NewFraction(7, 8)},
			wantResult: NewFraction(3, 7),
		},
		{
			name:       "test06",
			initial:    NewFraction(2, 5),
			others:     []*Fraction{NewFraction(1, 2)},
			wantResult: NewFraction(4, 5),
		},
		{
			name:       "test07",
			initial:    NewFraction(5, 10),
			others:     []*Fraction{NewFraction(6, 12)},
			wantResult: NewInteger(1),
		},
		{
			name:       "test08",
			initial:    NewFraction(7, 11),
			others:     []*Fraction{NewFraction(1, 6)},
			wantResult: NewFraction(42, 11),
		},
		{
			name:       "test09",
			initial:    NewFraction(5, 6),
			others:     []*Fraction{NewFraction(5, 10)},
			wantResult: NewFraction(5, 3),
		},
		{
			name:       "test10",
			initial:    NewFraction(1, 4),
			others:     []*Fraction{NewFraction(1, 7)},
			wantResult: NewFraction(7, 4),
		},
		{
			name:       "test11",
			initial:    NewFraction(1, 2),
			others:     []*Fraction{NewFraction(1, 3), NewFraction(1, 4)},
			wantResult: NewInteger(6),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.initial.Divide(tt.others...)
			if got == nil && tt.wantResult == nil {
				// Both are nil, test passed
				return
			}
			if got == nil || tt.wantResult == nil || got.Numerator() != tt.wantResult.Numerator() || got.Denominator() != tt.wantResult.Denominator() {
				t.Errorf("Fraction.Divide() = %v, want %v", got, tt.wantResult)
			}
		})
	}
}

func TestFraction_LaTeX(t *testing.T) {
	tests := []struct {
		name string
		f    *Fraction
		want string
	}{
		{
			name: "Fraction_LaTeX_Test01",
			f:    NewFraction(7, 8),
			want: `\dfrac{7}{8}`,
		},
		{
			name: "Fraction_LaTeX_Test02",
			f:    NewFraction(7, -8),
			want: `-\dfrac{7}{8}`,
		},
		{
			name: "Fraction_LaTeX_Test03",
			f:    NewFraction(7, 1),
			want: "7",
		},
		{
			name: "Fraction_LaTeX_Test04",
			f:    NewFraction(-7, 1),
			want: "-7",
		},
		{
			name: "Fraction_LaTeX_Test05",
			f:    NewInteger(88),
			want: "88",
		},
		{
			name: "Fraction_LaTeX_Test06",
			f:    NewInteger(0),
			want: "0",
		},
		{
			name: "Fraction_LaTeX_Test07",
			f:    NewFraction(-7, -1),
			want: "7",
		},
		{
			name: "Fraction_LaTeX_Test08",
			f:    NewFraction(7, -1),
			want: "-7",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.f.LaTeX(); got != tt.want {
				t.Errorf("Fraction.LaTeX() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFraction_IsInteger(t *testing.T) {
	tests := []struct {
		name string
		f    *Fraction
		want bool
	}{
		{
			name: "Fraction_IsInteger_Test01",
			f:    NewFraction(7, 8),
			want: false,
		},
		{
			name: "Fraction_IsInteger_Test02",
			f:    NewFraction(7, 1),
			want: true,
		},
		{
			name: "Fraction_IsInteger_Test03",
			f:    NewFraction(7, -1),
			want: true,
		},
		{
			name: "Fraction_IsInteger_Test04",
			f:    NewInteger(88),
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.f.IsInteger(); got != tt.want {
				t.Errorf("Fraction.IsInteger() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFraction_String(t *testing.T) {
	tests := []struct {
		name string
		f    *Fraction
		want string
	}{
		{
			name: "Fraction_String_Test01",
			f:    NewFraction(7, 8),
			want: `7/8`,
		},
		{
			name: "Fraction_String_Test02",
			f:    NewFraction(7, -8),
			want: `-7/8`,
		},
		{
			name: "Fraction_String_Test03",
			f:    NewFraction(7, 1),
			want: "7",
		},
		{
			name: "Fraction_String_Test04",
			f:    NewFraction(-7, 1),
			want: "-7",
		},
		{
			name: "Fraction_String_Test05",
			f:    NewInteger(88),
			want: "88",
		},
		{
			name: "Fraction_String_Test06",
			f:    NewInteger(0),
			want: "0",
		},
		{
			name: "Fraction_String_Test07",
			f:    NewFraction(-7, -1),
			want: "7",
		},
		{
			name: "Fraction_String_Test08",
			f:    NewFraction(7, -1),
			want: "-7",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.f.String(); got != tt.want {
				t.Errorf("Fraction.String() = %v, want %v", got, tt.want)
			}
		})
	}
}
