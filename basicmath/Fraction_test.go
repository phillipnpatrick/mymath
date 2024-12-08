package basicmath

import (
	"testing"
)

func TestFraction_Add(t *testing.T) {
	tests := []struct {
		name       string
		f1         *Fraction
		args       []any
		wantResult *Fraction
	}{
		{
			name:       "Add two fractions",
			f1:         NewFraction(WithNumerator(1), WithDenominator(2)),
			args:       []any{NewFraction(WithNumerator(1), WithDenominator(3))},
			wantResult: NewFraction(WithNumerator(5), WithDenominator(6)), // 1/2 + 1/3 = 5/6
		},
		{
			name:       "Nil first fraction",
			f1:         nil,
			args:       []any{NewFraction(WithNumerator(1), WithDenominator(3))},
			wantResult: nil,
		},
		{
			name:       "Invalid input type",
			f1:         NewFraction(WithNumerator(1), WithDenominator(2)),
			args:       []any{"invalid"},
			wantResult: nil,
		},
		{
			name: "Add multiple fractions (use only the first valid one)",
			f1:   NewFraction(WithNumerator(1), WithDenominator(2)),
			args: []any{
				NewFraction(WithNumerator(1), WithDenominator(3)),
				NewFraction(WithNumerator(1), WithDenominator(4)),
			},
			wantResult: NewFraction(WithNumerator(5), WithDenominator(6)), // Only the first valid fraction is used
		},
		{
			name:       "No arguments",
			f1:         NewFraction(WithNumerator(1), WithDenominator(2)),
			args:       []any{},
			wantResult: nil,
		},
		{
			name:       "test01",
			f1:         NewFraction(WithNumerator(1), WithDenominator(2)),
			args:       []any{NewFraction(WithNumerator(2), WithDenominator(3))},
			wantResult: NewFraction(WithNumerator(7), WithDenominator(6)),
		},
		{
			name:       "test02",
			f1:         NewFraction(WithNumerator(8), WithDenominator(12)),
			args:       []any{NewFraction(WithNumerator(8), WithDenominator(11))},
			wantResult: NewFraction(WithNumerator(46), WithDenominator(33)),
		},
		{
			name:       "test03",
			f1:         NewFraction(WithNumerator(2), WithDenominator(7)),
			args:       []any{NewFraction(WithNumerator(6), WithDenominator(10))},
			wantResult: NewFraction(WithNumerator(31), WithDenominator(35)),
		},
		{
			name:       "test04",
			f1:         NewFraction(WithNumerator(1), WithDenominator(6)),
			args:       []any{NewFraction(WithNumerator(6), WithDenominator(11))},
			wantResult: NewFraction(WithNumerator(47), WithDenominator(66)),
		},
		{
			name:       "test05",
			f1:         NewFraction(WithNumerator(5), WithDenominator(9)),
			args:       []any{NewFraction(WithNumerator(1), WithDenominator(2))},
			wantResult: NewFraction(WithNumerator(19), WithDenominator(18)),
		},
		{
			name:       "test06",
			f1:         NewFraction(WithNumerator(9), WithDenominator(12)),
			args:       []any{NewFraction(WithNumerator(2), WithDenominator(12))},
			wantResult: NewFraction(WithNumerator(11), WithDenominator(12)),
		},
		{
			name:       "test07",
			f1:         NewFraction(WithNumerator(5), WithDenominator(12)),
			args:       []any{NewFraction(WithNumerator(4), WithDenominator(12))},
			wantResult: NewFraction(WithNumerator(3), WithDenominator(4)),
		},
		{
			name:       "test08",
			f1:         NewFraction(WithNumerator(2), WithDenominator(7)),
			args:       []any{NewFraction(WithNumerator(1), WithDenominator(4))},
			wantResult: NewFraction(WithNumerator(15), WithDenominator(28)),
		},
		{
			name:       "test09",
			f1:         NewFraction(WithNumerator(1), WithDenominator(4)),
			args:       []any{NewFraction(WithNumerator(6), WithDenominator(8))},
			wantResult: NewFraction(WithNumerator(1)),
		},
		{
			name:       "test10",
			f1:         NewFraction(WithNumerator(4), WithDenominator(10)),
			args:       []any{NewFraction(WithNumerator(4), WithDenominator(5))},
			wantResult: NewFraction(WithNumerator(6), WithDenominator(5)),
		},
		{
			name:       "test11",
			f1:         NewFraction(WithNumerator(1), WithDenominator(2)),
			args:       []any{NewFraction(WithNumerator(8), WithDenominator(11))},
			wantResult: NewFraction(WithNumerator(27), WithDenominator(22)),
		},
		{
			name:       "test12",
			f1:         NewFraction(WithNumerator(1), WithDenominator(11)),
			args:       []any{NewFraction(WithNumerator(2), WithDenominator(12))},
			wantResult: NewFraction(WithNumerator(17), WithDenominator(66)),
		},
		{
			name:       "test13",
			f1:         NewFraction(WithNumerator(2), WithDenominator(12)),
			args:       []any{NewFraction(WithNumerator(2), WithDenominator(4))},
			wantResult: NewFraction(WithNumerator(2), WithDenominator(3)),
		},
		{
			name:       "test14",
			f1:         NewFraction(WithNumerator(3), WithDenominator(5)),
			args:       []any{NewFraction(WithNumerator(3), WithDenominator(8))},
			wantResult: NewFraction(WithNumerator(39), WithDenominator(40)),
		},
		{
			name:       "test15",
			f1:         NewFraction(WithNumerator(6), WithDenominator(9)),
			args:       []any{NewFraction(WithNumerator(1), WithDenominator(2))},
			wantResult: NewFraction(WithNumerator(7), WithDenominator(6)),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.f1.Add(tt.args...)
			if got == nil && tt.wantResult == nil {
				// Both are nil, test passed
				return
			}
			if got == nil || tt.wantResult == nil || got.(*Fraction).Numerator() != tt.wantResult.Numerator() || got.(*Fraction).Denominator() != tt.wantResult.Denominator() {
				t.Errorf("Fraction.Add() = %v, want %v", got, tt.wantResult)
			}
		})
	}
}

func TestFraction_Subtract(t *testing.T) {
	tests := []struct {
		name       string
		f1         *Fraction
		args       []any
		wantResult *Fraction
	}{
		{
			name:       "test01",
			f1:         NewFraction(WithNumerator(2), WithDenominator(9)),
			args:       []any{NewFraction(WithNumerator(1), WithDenominator(6))},
			wantResult: NewFraction(WithNumerator(1), WithDenominator(18)),
		},
		{
			name:       "test02",
			f1:         NewFraction(WithNumerator(16), WithDenominator(25)),
			args:       []any{NewFraction(WithNumerator(2), WithDenominator(4))},
			wantResult: NewFraction(WithNumerator(7), WithDenominator(50)),
		},
		{
			name:       "test03",
			f1:         NewFraction(WithNumerator(5), WithDenominator(7)),
			args:       []any{NewFraction(WithNumerator(1), WithDenominator(3))},
			wantResult: NewFraction(WithNumerator(8), WithDenominator(21)),
		},
		{
			name:       "test04",
			f1:         NewFraction(WithNumerator(5), WithDenominator(7)),
			args:       []any{NewFraction(WithNumerator(2), WithDenominator(5))},
			wantResult: NewFraction(WithNumerator(11), WithDenominator(35)),
		},
		{
			name:       "test05",
			f1:         NewFraction(WithNumerator(5), WithDenominator(9)),
			args:       []any{NewFraction(WithNumerator(1), WithDenominator(3))},
			wantResult: NewFraction(WithNumerator(2), WithDenominator(9)),
		},
		{
			name:       "test06",
			f1:         NewFraction(WithNumerator(6), WithDenominator(10)),
			args:       []any{NewFraction(WithNumerator(1), WithDenominator(8))},
			wantResult: NewFraction(WithNumerator(19), WithDenominator(40)),
		},
		{
			name:       "test07",
			f1:         NewFraction(WithNumerator(9), WithDenominator(16)),
			args:       []any{NewFraction(WithNumerator(1), WithDenominator(4))},
			wantResult: NewFraction(WithNumerator(5), WithDenominator(16)),
		},
		{
			name:       "test08",
			f1:         NewFraction(WithNumerator(3), WithDenominator(4)),
			args:       []any{NewFraction(WithNumerator(2), WithDenominator(6))},
			wantResult: NewFraction(WithNumerator(5), WithDenominator(12)),
		},
		{
			name:       "test09",
			f1:         NewFraction(WithNumerator(2), WithDenominator(3)),
			args:       []any{NewFraction(WithNumerator(1), WithDenominator(8))},
			wantResult: NewFraction(WithNumerator(13), WithDenominator(24)),
		},
		{
			name:       "test10",
			f1:         NewFraction(WithNumerator(8), WithDenominator(20)),
			args:       []any{NewFraction(WithNumerator(1), WithDenominator(4))},
			wantResult: NewFraction(WithNumerator(3), WithDenominator(20)),
		},
		{
			name:       "test11",
			f1:         NewFraction(WithNumerator(4), WithDenominator(7)),
			args:       []any{NewFraction(WithNumerator(4), WithDenominator(8))},
			wantResult: NewFraction(WithNumerator(1), WithDenominator(14)),
		},
		{
			name:       "test12",
			f1:         NewFraction(WithNumerator(5), WithDenominator(8)),
			args:       []any{NewFraction(WithNumerator(2), WithDenominator(6))},
			wantResult: NewFraction(WithNumerator(7), WithDenominator(24)),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.f1.Subtract(tt.args...)
			if got == nil && tt.wantResult == nil {
				// Both are nil, test passed
				return
			}
			if got == nil || tt.wantResult == nil || got.(*Fraction).Numerator() != tt.wantResult.Numerator() || got.(*Fraction).Denominator() != tt.wantResult.Denominator() {
				t.Errorf("Fraction.Subtract() = %v, want %v", got, tt.wantResult)
			}
		})
	}
}

func TestFraction_Multiply(t *testing.T) {
	tests := []struct {
		name       string
		f1         *Fraction
		args       []any
		wantResult *Fraction
	}{
		{
			name:       "test01",
			f1:         NewFraction(WithNumerator(6), WithDenominator(12)),
			args:       []any{NewFraction(WithNumerator(2), WithDenominator(10))},
			wantResult: NewFraction(WithNumerator(1), WithDenominator(10)),
		},
		{
			name:       "test02",
			f1:         NewFraction(WithNumerator(1), WithDenominator(16)),
			args:       []any{NewFraction(WithNumerator(7), WithDenominator(21))},
			wantResult: NewFraction(WithNumerator(1), WithDenominator(48)),
		},
		{
			name:       "test03",
			f1:         NewFraction(WithNumerator(8), WithDenominator(9)),
			args:       []any{NewFraction(WithNumerator(1), WithDenominator(2))},
			wantResult: NewFraction(WithNumerator(4), WithDenominator(9)),
		},
		{
			name:       "test04",
			f1:         NewFraction(WithNumerator(11), WithDenominator(20)),
			args:       []any{NewFraction(WithNumerator(4), WithDenominator(14))},
			wantResult: NewFraction(WithNumerator(11), WithDenominator(70)),
		},
		{
			name:       "test05",
			f1:         NewFraction(WithNumerator(4), WithDenominator(10)),
			args:       []any{NewFraction(WithNumerator(1), WithDenominator(6))},
			wantResult: NewFraction(WithNumerator(1), WithDenominator(15)),
		},
		{
			name:       "test06",
			f1:         NewFraction(WithNumerator(2), WithDenominator(5)),
			args:       []any{NewFraction(WithNumerator(1), WithDenominator(4))},
			wantResult: NewFraction(WithNumerator(1), WithDenominator(10)),
		},
		{
			name:       "test07",
			f1:         NewFraction(WithNumerator(2), WithDenominator(3)),
			args:       []any{NewFraction(WithNumerator(2), WithDenominator(10))},
			wantResult: NewFraction(WithNumerator(2), WithDenominator(15)),
		},
		{
			name:       "test08",
			f1:         NewFraction(WithNumerator(8), WithDenominator(10)),
			args:       []any{NewFraction(WithNumerator(4), WithDenominator(7))},
			wantResult: NewFraction(WithNumerator(16), WithDenominator(35)),
		},
		{
			name:       "test09",
			f1:         NewFraction(WithNumerator(8), WithDenominator(12)),
			args:       []any{NewFraction(WithNumerator(1), WithDenominator(5))},
			wantResult: NewFraction(WithNumerator(2), WithDenominator(15)),
		},
		{
			name:       "test10",
			f1:         NewFraction(WithNumerator(5), WithDenominator(7)),
			args:       []any{NewFraction(WithNumerator(1), WithDenominator(3))},
			wantResult: NewFraction(WithNumerator(5), WithDenominator(21)),
		},
		{
			name:       "test11",
			f1:         NewFraction(WithNumerator(4), WithDenominator(5)),
			args:       []any{NewFraction(WithNumerator(3), WithDenominator(8))},
			wantResult: NewFraction(WithNumerator(3), WithDenominator(10)),
		},
		{
			name:       "test12",
			f1:         NewFraction(WithNumerator(1), WithDenominator(3)),
			args:       []any{NewFraction(WithNumerator(2), WithDenominator(6))},
			wantResult: NewFraction(WithNumerator(1), WithDenominator(9)),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.f1.Multiply(tt.args...)
			if got == nil && tt.wantResult == nil {
				// Both are nil, test passed
				return
			}
			if got == nil || tt.wantResult == nil || got.(*Fraction).Numerator() != tt.wantResult.Numerator() || got.(*Fraction).Denominator() != tt.wantResult.Denominator() {
				t.Errorf("Fraction.Multiply() = %v, want %v", got, tt.wantResult)
			}
		})
	}
}

func TestFraction_Divide(t *testing.T) {
	tests := []struct {
		name       string
		f1         *Fraction
		args       []any
		wantResult *Fraction
	}{
		{
			name:       "test01",
			f1:         NewFraction(WithNumerator(1), WithDenominator(4)),
			args:       []any{NewFraction(WithNumerator(9), WithDenominator(10))},
			wantResult: NewFraction(WithNumerator(5), WithDenominator(18)),
		},
		{
			name:       "test02",
			f1:         NewFraction(WithNumerator(5), WithDenominator(9)),
			args:       []any{NewFraction(WithNumerator(1), WithDenominator(2))},
			wantResult: NewFraction(WithNumerator(10), WithDenominator(9)),
		},
		{
			name:       "test03",
			f1:         NewFraction(WithNumerator(1), WithDenominator(3)),
			args:       []any{NewFraction(WithNumerator(6), WithDenominator(9))},
			wantResult: NewFraction(WithNumerator(1), WithDenominator(2)),
		},
		{
			name:       "test04",
			f1:         NewFraction(WithNumerator(8), WithDenominator(10)),
			args:       []any{NewFraction(WithNumerator(2), WithDenominator(5))},
			wantResult: NewFraction(WithNumerator(2)),
		},
		{
			name:       "test05",
			f1:         NewFraction(WithNumerator(3), WithDenominator(8)),
			args:       []any{NewFraction(WithNumerator(7), WithDenominator(8))},
			wantResult: NewFraction(WithNumerator(3), WithDenominator(7)),
		},
		{
			name:       "test06",
			f1:         NewFraction(WithNumerator(2), WithDenominator(5)),
			args:       []any{NewFraction(WithNumerator(1), WithDenominator(2))},
			wantResult: NewFraction(WithNumerator(4), WithDenominator(5)),
		},
		{
			name:       "test07",
			f1:         NewFraction(WithNumerator(5), WithDenominator(10)),
			args:       []any{NewFraction(WithNumerator(6), WithDenominator(12))},
			wantResult: NewFraction(WithNumerator(1)),
		},
		{
			name:       "test08",
			f1:         NewFraction(WithNumerator(7), WithDenominator(11)),
			args:       []any{NewFraction(WithNumerator(1), WithDenominator(6))},
			wantResult: NewFraction(WithNumerator(42), WithDenominator(11)),
		},
		{
			name:       "test09",
			f1:         NewFraction(WithNumerator(5), WithDenominator(6)),
			args:       []any{NewFraction(WithNumerator(5), WithDenominator(10))},
			wantResult: NewFraction(WithNumerator(5), WithDenominator(3)),
		},
		{
			name:       "test10",
			f1:         NewFraction(WithNumerator(1), WithDenominator(4)),
			args:       []any{NewFraction(WithNumerator(1), WithDenominator(7))},
			wantResult: NewFraction(WithNumerator(7), WithDenominator(4)),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.f1.Divide(tt.args...)
			if got == nil && tt.wantResult == nil {
				// Both are nil, test passed
				return
			}
			if got == nil || tt.wantResult == nil || got.(*Fraction).Numerator() != tt.wantResult.Numerator() || got.(*Fraction).Denominator() != tt.wantResult.Denominator() {
				t.Errorf("Fraction.Divide() = %v, want %v", got, tt.wantResult)
			}
		})
	}
}
