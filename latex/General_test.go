package latex

import (
	"testing"
)

func TestConnectWithMinusSign(t *testing.T) {
	type args struct {
		a string
		b string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "LaTeX_ConnectWithMinusSign_Test01",
			args: args{"4x", "8"},
			want: "4x - 8",
		},
		{
			name: "LaTeX_ConnectWithMinusSign_Test02",
			args: args{"4x", "-8"},
			want: "4x + 8",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ConnectWithMinusSign(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("ConnectWithMinusSign() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConnectWithPlusSign(t *testing.T) {
	type args struct {
		a string
		b string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "LaTeX_ConnectWithPlusSign_Test01",
			args: args{"4x", "8"},
			want: "4x + 8",
		},
		{
			name: "LaTeX_ConnectWithPlusSign_Test02",
			args: args{"4x", "-8"},
			want: "4x - 8",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ConnectWithPlusSign(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("ConnectWithPlusSign() = %v, want %v", got, tt.want)
			}
		})
	}
}
