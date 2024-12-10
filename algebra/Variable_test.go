package algebra

import (
	"mymath/basicmath"
	"testing"
)

func TestVariable_Equals(t *testing.T) {
	type args struct {
		other *Variable
	}
	tests := []struct {
		name string
		v    *Variable
		args args
		want bool
	}{
		{
			name: "Variable_Equals_Test01",
			v:    NewVariable("a"),
			args: args{other: NewVariable("a")},
			want: true,
		},
		{
			name: "Variable_Equals_Test02",
			v:    NewVariable("a"),
			args: args{other: NewVariable("b")},
			want: false,
		},
		{
			name: "Variable_Equals_Test03",
			v:    NewVariableWithDegree("c", basicmath.NewInteger(1)),
			args: args{other: NewVariableWithDegree("c", basicmath.NewInteger(1))},
			want: true,
		},
		{
			name: "Variable_Equals_Test04",
			v:    NewVariableWithDegree("c", basicmath.NewInteger(0)),
			args: args{other: NewVariableWithDegree("c", basicmath.NewInteger(1))},
			want: false,
		},
		{
			name: "Variable_Equals_Test05",
			v:    NewVariable("mnp"),
			args: args{other: NewVariable("m")},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.v.Equals(tt.args.other); got != tt.want {
				t.Errorf("Variable.Equals() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVariable_LaTeX(t *testing.T) {
	tests := []struct {
		name string
		v    Variable
		want string
	}{
		{
			name: "Variable_LaTeX_Test01",
			v:    *NewVariable("a"),
			want: "a",
		},
		{
			name: "Variable_LaTeX_Test02",
			v:    *NewVariableWithDegree("a", basicmath.NewInteger(0)),
			want: "",
		},
		{
			name: "Variable_LaTeX_Test03",
			v:    *NewVariableWithDegree("a", basicmath.NewInteger(8)),
			want: "a^8",
		},
		{
			name: "Variable_LaTeX_Test04",
			v:    *NewVariableWithDegree("a", basicmath.NewFraction(1, 2)),
			want: `a^\left(\dfrac{1}{2}\right)`,
		},
		{
			name: "Variable_LaTeX_Test05",
			v:    *NewVariableWithDegree("a", basicmath.NewFraction(-1, 2)),
			want: `a^\left(-\dfrac{1}{2}\right)`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.v.LaTeX(); got != tt.want {
				t.Errorf("Variable.LaTeX() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVariable_String(t *testing.T) {
	tests := []struct {
		name string
		v    Variable
		want string
	}{
		{
			name: "Variable_String_Test01",
			v:    *NewVariable("a"),
			want: "a",
		},
		{
			name: "Variable_String_Test02",
			v:    *NewVariableWithDegree("a", basicmath.NewInteger(0)),
			want: "",
		},
		{
			name: "Variable_String_Test03",
			v:    *NewVariableWithDegree("a", basicmath.NewInteger(8)),
			want: "a^8",
		},
		{
			name: "Variable_String_Test04",
			v:    *NewVariableWithDegree("a", basicmath.NewFraction(1, 2)),
			want: `a^(1/2)`,
		},
		{
			name: "Variable_String_Test05",
			v:    *NewVariableWithDegree("a", basicmath.NewFraction(-1, 2)),
			want: `a^(-1/2)`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.v.String(); got != tt.want {
				t.Errorf("Variable.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVariable_IsLikeTerm(t *testing.T) {
	type args struct {
		other Variable
	}
	tests := []struct {
		name string
		v    Variable
		args args
		want bool
	}{
		{
			name: "Variable_IsLikeTerm_Test01",
			v:    *NewVariable("a"),
			args: args{other: *NewVariable("a")},
			want: true,
		},
		{
			name: "Variable_IsLikeTerm_Test02",
			v:    *NewVariable("a"),
			args: args{other: *NewVariable("b")},
			want: false,
		},
		{
			name: "Variable_IsLikeTerm_Test03",
			v:    *NewVariableWithDegree("a", basicmath.NewInteger(2)),
			args: args{other: *NewVariableWithDegree("a", basicmath.NewInteger(2))},
			want: true,
		},
		{
			name: "Variable_IsLikeTerm_Test04",
			v:    *NewVariableWithDegree("a", basicmath.NewInteger(2)),
			args: args{other: *NewVariable("a")},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.v.IsLikeTerm(tt.args.other); got != tt.want {
				t.Errorf("Variable.IsLikeTerm() = %v, want %v", got, tt.want)
			}
		})
	}
}
