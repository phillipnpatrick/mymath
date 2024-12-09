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
			initial:    NewFraction(WithNumerator(1), WithDenominator(2)),
			others:     []*Fraction{NewFraction(WithNumerator(1), WithDenominator(3))},
			wantResult: NewFraction(WithNumerator(5), WithDenominator(6)), // 1/2 + 1/3 = 5/6
		},
		{
			name:       "test01",
			initial:    NewFraction(WithNumerator(1), WithDenominator(2)),
			others:     []*Fraction{NewFraction(WithNumerator(2), WithDenominator(3))},
			wantResult: NewFraction(WithNumerator(7), WithDenominator(6)),
		},
		{
			name:       "test02",
			initial:    NewFraction(WithNumerator(8), WithDenominator(12)),
			others:     []*Fraction{NewFraction(WithNumerator(8), WithDenominator(11))},
			wantResult: NewFraction(WithNumerator(46), WithDenominator(33)),
		},
		{
			name:       "test03",
			initial:    NewFraction(WithNumerator(2), WithDenominator(7)),
			others:     []*Fraction{NewFraction(WithNumerator(6), WithDenominator(10))},
			wantResult: NewFraction(WithNumerator(31), WithDenominator(35)),
		},
		{
			name:       "test04",
			initial:    NewFraction(WithNumerator(1), WithDenominator(6)),
			others:     []*Fraction{NewFraction(WithNumerator(6), WithDenominator(11))},
			wantResult: NewFraction(WithNumerator(47), WithDenominator(66)),
		},
		{
			name:       "test05",
			initial:    NewFraction(WithNumerator(5), WithDenominator(9)),
			others:     []*Fraction{NewFraction(WithNumerator(1), WithDenominator(2))},
			wantResult: NewFraction(WithNumerator(19), WithDenominator(18)),
		},
		{
			name:       "test06",
			initial:    NewFraction(WithNumerator(9), WithDenominator(12)),
			others:     []*Fraction{NewFraction(WithNumerator(2), WithDenominator(12))},
			wantResult: NewFraction(WithNumerator(11), WithDenominator(12)),
		},
		{
			name:       "test07",
			initial:    NewFraction(WithNumerator(5), WithDenominator(12)),
			others:     []*Fraction{NewFraction(WithNumerator(4), WithDenominator(12))},
			wantResult: NewFraction(WithNumerator(3), WithDenominator(4)),
		},
		{
			name:       "test08",
			initial:    NewFraction(WithNumerator(2), WithDenominator(7)),
			others:     []*Fraction{NewFraction(WithNumerator(1), WithDenominator(4))},
			wantResult: NewFraction(WithNumerator(15), WithDenominator(28)),
		},
		{
			name:       "test09",
			initial:    NewFraction(WithNumerator(1), WithDenominator(4)),
			others:     []*Fraction{NewFraction(WithNumerator(6), WithDenominator(8))},
			wantResult: NewFraction(WithNumerator(1)),
		},
		{
			name:       "test10",
			initial:    NewFraction(WithNumerator(4), WithDenominator(10)),
			others:     []*Fraction{NewFraction(WithNumerator(4), WithDenominator(5))},
			wantResult: NewFraction(WithNumerator(6), WithDenominator(5)),
		},
		{
			name:       "test11",
			initial:    NewFraction(WithNumerator(1), WithDenominator(2)),
			others:     []*Fraction{NewFraction(WithNumerator(8), WithDenominator(11))},
			wantResult: NewFraction(WithNumerator(27), WithDenominator(22)),
		},
		{
			name:       "test12",
			initial:    NewFraction(WithNumerator(1), WithDenominator(11)),
			others:     []*Fraction{NewFraction(WithNumerator(2), WithDenominator(12))},
			wantResult: NewFraction(WithNumerator(17), WithDenominator(66)),
		},
		{
			name:       "test13",
			initial:    NewFraction(WithNumerator(2), WithDenominator(12)),
			others:     []*Fraction{NewFraction(WithNumerator(2), WithDenominator(4))},
			wantResult: NewFraction(WithNumerator(2), WithDenominator(3)),
		},
		{
			name:       "test14",
			initial:    NewFraction(WithNumerator(3), WithDenominator(5)),
			others:     []*Fraction{NewFraction(WithNumerator(3), WithDenominator(8))},
			wantResult: NewFraction(WithNumerator(39), WithDenominator(40)),
		},
		{
			name:       "test15",
			initial:    NewFraction(WithNumerator(6), WithDenominator(9)),
			others:     []*Fraction{NewFraction(WithNumerator(1), WithDenominator(2))},
			wantResult: NewFraction(WithNumerator(7), WithDenominator(6)),
		},
		{
			name:       "test16",
			initial:    NewFraction(WithNumerator(1), WithDenominator(2)),
			others:     []*Fraction{NewFraction(WithNumerator(1), WithDenominator(3)), NewFraction(WithNumerator(1), WithDenominator(4))},
			wantResult: NewFraction(WithNumerator(13), WithDenominator(12)),
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
			initial:         NewFraction(WithNumerator(2), WithDenominator(9)),
			others:       []*Fraction{NewFraction(WithNumerator(1), WithDenominator(6))},
			wantResult: NewFraction(WithNumerator(1), WithDenominator(18)),
		},
		{
			name:       "test02",
			initial:         NewFraction(WithNumerator(16), WithDenominator(25)),
			others:       []*Fraction{NewFraction(WithNumerator(2), WithDenominator(4))},
			wantResult: NewFraction(WithNumerator(7), WithDenominator(50)),
		},
		{
			name:       "test03",
			initial:         NewFraction(WithNumerator(5), WithDenominator(7)),
			others:       []*Fraction{NewFraction(WithNumerator(1), WithDenominator(3))},
			wantResult: NewFraction(WithNumerator(8), WithDenominator(21)),
		},
		{
			name:       "test04",
			initial:         NewFraction(WithNumerator(5), WithDenominator(7)),
			others:       []*Fraction{NewFraction(WithNumerator(2), WithDenominator(5))},
			wantResult: NewFraction(WithNumerator(11), WithDenominator(35)),
		},
		{
			name:       "test05",
			initial:         NewFraction(WithNumerator(5), WithDenominator(9)),
			others:       []*Fraction{NewFraction(WithNumerator(1), WithDenominator(3))},
			wantResult: NewFraction(WithNumerator(2), WithDenominator(9)),
		},
		{
			name:       "test06",
			initial:         NewFraction(WithNumerator(6), WithDenominator(10)),
			others:       []*Fraction{NewFraction(WithNumerator(1), WithDenominator(8))},
			wantResult: NewFraction(WithNumerator(19), WithDenominator(40)),
		},
		{
			name:       "test07",
			initial:         NewFraction(WithNumerator(9), WithDenominator(16)),
			others:       []*Fraction{NewFraction(WithNumerator(1), WithDenominator(4))},
			wantResult: NewFraction(WithNumerator(5), WithDenominator(16)),
		},
		{
			name:       "test08",
			initial:         NewFraction(WithNumerator(3), WithDenominator(4)),
			others:       []*Fraction{NewFraction(WithNumerator(2), WithDenominator(6))},
			wantResult: NewFraction(WithNumerator(5), WithDenominator(12)),
		},
		{
			name:       "test09",
			initial:         NewFraction(WithNumerator(2), WithDenominator(3)),
			others:       []*Fraction{NewFraction(WithNumerator(1), WithDenominator(8))},
			wantResult: NewFraction(WithNumerator(13), WithDenominator(24)),
		},
		{
			name:       "test10",
			initial:         NewFraction(WithNumerator(8), WithDenominator(20)),
			others:       []*Fraction{NewFraction(WithNumerator(1), WithDenominator(4))},
			wantResult: NewFraction(WithNumerator(3), WithDenominator(20)),
		},
		{
			name:       "test11",
			initial:         NewFraction(WithNumerator(4), WithDenominator(7)),
			others:       []*Fraction{NewFraction(WithNumerator(4), WithDenominator(8))},
			wantResult: NewFraction(WithNumerator(1), WithDenominator(14)),
		},
		{
			name:       "test12",
			initial:         NewFraction(WithNumerator(5), WithDenominator(8)),
			others:       []*Fraction{NewFraction(WithNumerator(2), WithDenominator(6))},
			wantResult: NewFraction(WithNumerator(7), WithDenominator(24)),
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
		initial         *Fraction
		others       []*Fraction
		wantResult *Fraction
	}{
		{
			name:       "test01",
			initial:         NewFraction(WithNumerator(6), WithDenominator(12)),
			others:       []*Fraction{NewFraction(WithNumerator(2), WithDenominator(10))},
			wantResult: NewFraction(WithNumerator(1), WithDenominator(10)),
		},
		{
			name:       "test02",
			initial:         NewFraction(WithNumerator(1), WithDenominator(16)),
			others:       []*Fraction{NewFraction(WithNumerator(7), WithDenominator(21))},
			wantResult: NewFraction(WithNumerator(1), WithDenominator(48)),
		},
		{
			name:       "test03",
			initial:         NewFraction(WithNumerator(8), WithDenominator(9)),
			others:       []*Fraction{NewFraction(WithNumerator(1), WithDenominator(2))},
			wantResult: NewFraction(WithNumerator(4), WithDenominator(9)),
		},
		{
			name:       "test04",
			initial:         NewFraction(WithNumerator(11), WithDenominator(20)),
			others:       []*Fraction{NewFraction(WithNumerator(4), WithDenominator(14))},
			wantResult: NewFraction(WithNumerator(11), WithDenominator(70)),
		},
		{
			name:       "test05",
			initial:         NewFraction(WithNumerator(4), WithDenominator(10)),
			others:       []*Fraction{NewFraction(WithNumerator(1), WithDenominator(6))},
			wantResult: NewFraction(WithNumerator(1), WithDenominator(15)),
		},
		{
			name:       "test06",
			initial:         NewFraction(WithNumerator(2), WithDenominator(5)),
			others:       []*Fraction{NewFraction(WithNumerator(1), WithDenominator(4))},
			wantResult: NewFraction(WithNumerator(1), WithDenominator(10)),
		},
		{
			name:       "test07",
			initial:         NewFraction(WithNumerator(2), WithDenominator(3)),
			others:       []*Fraction{NewFraction(WithNumerator(2), WithDenominator(10))},
			wantResult: NewFraction(WithNumerator(2), WithDenominator(15)),
		},
		{
			name:       "test08",
			initial:         NewFraction(WithNumerator(8), WithDenominator(10)),
			others:       []*Fraction{NewFraction(WithNumerator(4), WithDenominator(7))},
			wantResult: NewFraction(WithNumerator(16), WithDenominator(35)),
		},
		{
			name:       "test09",
			initial:         NewFraction(WithNumerator(8), WithDenominator(12)),
			others:       []*Fraction{NewFraction(WithNumerator(1), WithDenominator(5))},
			wantResult: NewFraction(WithNumerator(2), WithDenominator(15)),
		},
		{
			name:       "test10",
			initial:         NewFraction(WithNumerator(5), WithDenominator(7)),
			others:       []*Fraction{NewFraction(WithNumerator(1), WithDenominator(3))},
			wantResult: NewFraction(WithNumerator(5), WithDenominator(21)),
		},
		{
			name:       "test11",
			initial:         NewFraction(WithNumerator(4), WithDenominator(5)),
			others:       []*Fraction{NewFraction(WithNumerator(3), WithDenominator(8))},
			wantResult: NewFraction(WithNumerator(3), WithDenominator(10)),
		},
		{
			name:       "test12",
			initial:         NewFraction(WithNumerator(1), WithDenominator(3)),
			others:       []*Fraction{NewFraction(WithNumerator(2), WithDenominator(6))},
			wantResult: NewFraction(WithNumerator(1), WithDenominator(9)),
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
		initial         *Fraction
		others       []*Fraction
		wantResult *Fraction
	}{
		{
			name:       "test01",
			initial:         NewFraction(WithNumerator(1), WithDenominator(4)),
			others:       []*Fraction{NewFraction(WithNumerator(9), WithDenominator(10))},
			wantResult: NewFraction(WithNumerator(5), WithDenominator(18)),
		},
		{
			name:       "test02",
			initial:         NewFraction(WithNumerator(5), WithDenominator(9)),
			others:       []*Fraction{NewFraction(WithNumerator(1), WithDenominator(2))},
			wantResult: NewFraction(WithNumerator(10), WithDenominator(9)),
		},
		{
			name:       "test03",
			initial:         NewFraction(WithNumerator(1), WithDenominator(3)),
			others:       []*Fraction{NewFraction(WithNumerator(6), WithDenominator(9))},
			wantResult: NewFraction(WithNumerator(1), WithDenominator(2)),
		},
		{
			name:       "test04",
			initial:         NewFraction(WithNumerator(8), WithDenominator(10)),
			others:       []*Fraction{NewFraction(WithNumerator(2), WithDenominator(5))},
			wantResult: NewFraction(WithNumerator(2)),
		},
		{
			name:       "test05",
			initial:         NewFraction(WithNumerator(3), WithDenominator(8)),
			others:       []*Fraction{NewFraction(WithNumerator(7), WithDenominator(8))},
			wantResult: NewFraction(WithNumerator(3), WithDenominator(7)),
		},
		{
			name:       "test06",
			initial:         NewFraction(WithNumerator(2), WithDenominator(5)),
			others:       []*Fraction{NewFraction(WithNumerator(1), WithDenominator(2))},
			wantResult: NewFraction(WithNumerator(4), WithDenominator(5)),
		},
		{
			name:       "test07",
			initial:         NewFraction(WithNumerator(5), WithDenominator(10)),
			others:       []*Fraction{NewFraction(WithNumerator(6), WithDenominator(12))},
			wantResult: NewFraction(WithNumerator(1)),
		},
		{
			name:       "test08",
			initial:         NewFraction(WithNumerator(7), WithDenominator(11)),
			others:       []*Fraction{NewFraction(WithNumerator(1), WithDenominator(6))},
			wantResult: NewFraction(WithNumerator(42), WithDenominator(11)),
		},
		{
			name:       "test09",
			initial:         NewFraction(WithNumerator(5), WithDenominator(6)),
			others:       []*Fraction{NewFraction(WithNumerator(5), WithDenominator(10))},
			wantResult: NewFraction(WithNumerator(5), WithDenominator(3)),
		},
		{
			name:       "test10",
			initial:         NewFraction(WithNumerator(1), WithDenominator(4)),
			others:       []*Fraction{NewFraction(WithNumerator(1), WithDenominator(7))},
			wantResult: NewFraction(WithNumerator(7), WithDenominator(4)),
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
