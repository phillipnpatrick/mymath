package datastructures

import (
	"reflect"
	"sort"
	"testing"
)

func TestUnion(t *testing.T) {
	type args struct {
		map1 map[string]int
		map2 map[string]int
	}
	tests := []struct {
		name string
		args args
		want map[string]int
	}{
		{
			name: "Maps_Union_OverlappingMaps",
			args: args{
				map1: map[string]int{"apple": 1, "banana": 2},
				map2: map[string]int{"banana": 3, "cherry": 4},
			},
			want: map[string]int{"apple": 1, "banana": 3, "cherry": 4},
		},
		{
			name: "Maps_Union_DisjointMaps",
			args: args{
				map1: map[string]int{"apple": 1, "banana": 2},
				map2: map[string]int{"cherry": 3, "date": 4},
			},
			want: map[string]int{"apple": 1, "banana": 2, "cherry": 3, "date": 4},
		},
		{
			name: "Maps_Union_EmptyMap1",
			args: args{
				map1: map[string]int{},
				map2: map[string]int{"cherry": 3, "date": 4},
			},
			want: map[string]int{"cherry": 3, "date": 4},
		},
		{
			name: "Maps_Union_EmptyMap2",
			args: args{
				map1: map[string]int{"apple": 1, "banana": 2},
				map2: map[string]int{},
			},
			want: map[string]int{"apple": 1, "banana": 2},
		},
		{
			name: "Maps_Union_EmptyMaps",
			args: args{
				map1: map[string]int{},
				map2: map[string]int{},
			},
			want: map[string]int{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Union(tt.args.map1, tt.args.map2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Union() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIntersection(t *testing.T) {
	type args struct {
		map1 map[string]int
		map2 map[string]int
	}
	tests := []struct {
		name string
		args args
		want map[string]int
	}{
		{
			name: "Maps_Intersection_Test01",
			args: args{
				map1: map[string]int{"apple": 1, "banana": 2, "cherry": 3},
				map2: map[string]int{"banana": 3, "cherry": 4, "date": 5},
			},
			want: map[string]int{"banana": 3, "cherry": 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Intersection(tt.args.map1, tt.args.map2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Intersection() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDifference(t *testing.T) {
	type args struct {
		map1 map[string]int
		map2 map[string]int
	}
	tests := []struct {
		name string
		args args
		want map[string]int
	}{
		{
			name: "Maps_Difference_Test01",
			args: args{
				map1: map[string]int{"apple": 1, "banana": 2, "cherry": 3},
				map2: map[string]int{"banana": 3, "cherry": 4, "date": 5},
			},
			want: map[string]int{"apple": 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Difference(tt.args.map1, tt.args.map2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Difference() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSymmetricDifference(t *testing.T) {
	type args struct {
		map1 map[string]int
		map2 map[string]int
	}
	tests := []struct {
		name string
		args args
		want map[string]int
	}{
		{
			name: "Maps_SymmetricDifference_Test01",
			args: args{
				map1: map[string]int{"apple": 1, "banana": 2},
				map2: map[string]int{"banana": 3, "cherry": 4},
			},
			want: map[string]int{"apple": 1, "cherry": 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SymmetricDifference(tt.args.map1, tt.args.map2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SymmetricDifference() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestKeys(t *testing.T) {
	type args struct {
		m map[string]int
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "Maps_Keys_NonEmptyMap",
			args: args{
				m: map[string]int{"apple": 1, "banana": 2, "cherry": 3},
			},
			want: []string{"apple", "banana", "cherry"},
		},
		{
			name: "Maps_Keys_EmptyMap",
			args: args{
				m: map[string]int{},
			},
			want: []string{},
		},
		{
			name: "Maps_Keys_SingleKeyValuePair",
			args: args{
				m: map[string]int{"apple": 1},
			},
			want: []string{"apple"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Keys(tt.args.m)

			sort.Strings(got)
			sort.Strings(tt.want)

			reflect.DeepEqual(got, tt.want)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Keys() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestValues(t *testing.T) {
	type args struct {
		m map[string]int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "Maps_Values_NonEmptyMap",
			args: args{
				m: map[string]int{"apple": 1, "banana": 2, "cherry": 3},
			},
			want: []int{1,2,3},
		},
		{
			name: "Maps_Values_EmptyMap",
			args: args{
				m: map[string]int{},
			},
			want: []int{},
		},
		{
			name: "Maps_Values_SingleKeyValuePair",
			args: args{
				m: map[string]int{"apple": 42},
			},
			want: []int{42},
		},
		{
			name: "Maps_Values_UnorderedValues",
			args: args{
				m: map[string]int{"apple": 10, "banana": 20, "cherry": 30,},
			},
			want: []int{10,20,30},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Values(tt.args.m)

			sort.Ints(got)
			sort.Ints(tt.want)

			reflect.DeepEqual(got, tt.want)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Values() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFilter(t *testing.T) {
	tests := []struct {
		name      string
		input     map[string]int
		predicate func(string, int) bool
		expected  map[string]int
	}{
		{
			name: "Filter even values",
			input: map[string]int{
				"one": 1, "two": 2, "three": 3, "four": 4,
			},
			predicate: func(key string, value int) bool {
				return value%2 == 0
			},
			expected: map[string]int{
				"two": 2, "four": 4,
			},
		},
		{
			name: "Filter keys starting with 't'",
			input: map[string]int{
				"one": 1, "two": 2, "three": 3, "four": 4,
			},
			predicate: func(key string, value int) bool {
				return key[0] == 't'
			},
			expected: map[string]int{
				"two": 2, "three": 3,
			},
		},
		{
			name: "Empty input map",
			input: map[string]int{},
			predicate: func(key string, value int) bool {
				return true
			},
			expected: map[string]int{},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := Filter(test.input, test.predicate)
			if !reflect.DeepEqual(result, test.expected) {
				t.Errorf("Test %s failed. Expected %v, got %v", test.name, test.expected, result)
			}
		})
	}
}

func TestMapsEqual(t *testing.T) {
	tests := []struct {
		name     string
		map1     map[string]int
		map2     map[string]int
		expected bool
	}{
		{
			name: "Equal maps",
			map1: map[string]int{"a": 1, "b": 2},
			map2: map[string]int{"a": 1, "b": 2},
			expected: true,
		},
		{
			name: "Different maps",
			map1: map[string]int{"a": 1, "b": 2},
			map2: map[string]int{"a": 1, "b": 3},
			expected: false,
		},
		{
			name: "Empty maps",
			map1: map[string]int{},
			map2: map[string]int{},
			expected: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := MapsEqual(test.map1, test.map2)
			if result != test.expected {
				t.Errorf("Test %s failed. Expected %v, got %v", test.name, test.expected, result)
			}
		})
	}
}

func TestMergeWithResolver(t *testing.T) {
	tests := []struct {
		name      string
		map1      map[string]int
		map2      map[string]int
		resolver  func(int, int) int
		expected  map[string]int
	}{
		{
			name: "Non-overlapping maps",
			map1: map[string]int{"a": 1},
			map2: map[string]int{"b": 2},
			resolver: func(v1, v2 int) int {
				return v1 + v2 // Resolver is irrelevant here since no keys overlap.
			},
			expected: map[string]int{"a": 1, "b": 2},
		},
		{
			name: "Overlapping maps with addition resolver",
			map1: map[string]int{"a": 1},
			map2: map[string]int{"a": 2, "b": 3},
			resolver: func(v1, v2 int) int {
				return v1 + v2
			},
			expected: map[string]int{"a": 3, "b": 3},
		},
		{
			name: "Overlapping maps with maximum resolver",
			map1: map[string]int{"a": 5, "b": 2},
			map2: map[string]int{"a": 7, "b": 3, "c": 1},
			resolver: func(v1, v2 int) int {
				if v1 > v2 {
					return v1
				}
				return v2
			},
			expected: map[string]int{"a": 7, "b": 3, "c": 1},
		},
		{
			name: "Empty map1",
			map1: map[string]int{},
			map2: map[string]int{"a": 1, "b": 2},
			resolver: func(v1, v2 int) int {
				return v1 + v2
			},
			expected: map[string]int{"a": 1, "b": 2},
		},
		{
			name: "Empty map2",
			map1: map[string]int{"a": 1, "b": 2},
			map2: map[string]int{},
			resolver: func(v1, v2 int) int {
				return v1 + v2
			},
			expected: map[string]int{"a": 1, "b": 2},
		},
		{
			name: "Empty maps",
			map1: map[string]int{},
			map2: map[string]int{},
			resolver: func(v1, v2 int) int {
				return v1 + v2
			},
			expected: map[string]int{},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := Merge(test.map1, test.map2, test.resolver)
			if !reflect.DeepEqual(result, test.expected) {
				t.Errorf("Test %s failed. Expected %v, got %v", test.name, test.expected, result)
			}
		})
	}
}

func TestInvert(t *testing.T) {
	tests := []struct {
		name     string
		input    map[string]int
		expected map[int]string
	}{
		{
			name: "Invert map",
			input: map[string]int{"a": 1, "b": 2},
			expected: map[int]string{1: "a", 2: "b"},
		},
		{
			name: "Empty map",
			input: map[string]int{},
			expected: map[int]string{},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := Invert(test.input)
			if !reflect.DeepEqual(result, test.expected) {
				t.Errorf("Test %s failed. Expected %v, got %v", test.name, test.expected, result)
			}
		})
	}
}

func TestMapApply(t *testing.T) {
	tests := []struct {
		name     string
		input    map[string]int
		apply    func(string, int) int
		expected map[string]int
	}{
		{
			name: "Double the values",
			input: map[string]int{"a": 1, "b": 2},
			apply: func(key string, value int) int {
				return value * 2
			},
			expected: map[string]int{"a": 2, "b": 4},
		},
		{
			name: "Append key length to value",
			input: map[string]int{"apple": 1, "cat": 2},
			apply: func(key string, value int) int {
				return value + len(key)
			},
			expected: map[string]int{"apple": 6, "cat": 5},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := MapApply(test.input, test.apply)
			if !reflect.DeepEqual(result, test.expected) {
				t.Errorf("Test %s failed. Expected %v, got %v", test.name, test.expected, result)
			}
		})
	}
}

