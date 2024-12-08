package basicmath

import (
	"testing"
)

func TestLcm(t *testing.T) {
	type args struct {
		numbers []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test01",
			args: args{numbers: []int{6, 7, 21}},
			want: 42,
		},
		{
			name: "test02",
			args: args{numbers: []int{12, 30}},
			want: 60,
		},
		{
			name: "test03",
			args: args{numbers: []int{24, 300}},
			want: 600,
		},
		{
			name: "test04",
			args: args{numbers: []int{12, 18, 30}},
			want: 180,
		},
		{
			name: "test05",
			args: args{numbers: []int{12, 18}},
			want: 36,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LCM(tt.args.numbers...); got != tt.want {
				t.Errorf("LCM() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGcf(t *testing.T) {
	type args struct {
		numbers []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test01",
			args: args{numbers: []int{12, 15}},
			want: 3,
		},
		{
			name: "test02",
			args: args{numbers: []int{24, 30}},
			want: 6,
		},
		{
			name: "test03",
			args: args{numbers: []int{30, 40}},
			want: 10,
		},
		{
			name: "test04",
			args: args{numbers: []int{35, 28}},
			want: 7,
		},
		{
			name: "test05",
			args: args{numbers: []int{45, 36}},
			want: 9,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GCF(tt.args.numbers...); got != tt.want {
				t.Errorf("GCF() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMin(t *testing.T) {
	type args struct {
		numbers []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test01",
			args: args{numbers: []int{45, 36, 48, 62, 101, 5}},
			want: 5,
		},
		{
			name: "test02",
			args: args{numbers: []int{45, 36, 48, 62, 101, 5, 229, 331, -45, 10890, 80002}},
			want: -45,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Min(tt.args.numbers...); got != tt.want {
				t.Errorf("Min() = %v, want %v", got, tt.want)
			}
		})
	}
}
