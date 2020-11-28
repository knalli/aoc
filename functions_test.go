package aoc

import (
	"reflect"
	"testing"
)

func TestMaxInt(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"same negative",
			args{-1, -1},
			-1,
		},
		{
			"same zero",
			args{0, 0},
			0,
		},
		{
			"same positive",
			args{1, 1},
			1,
		},
		{
			"negative and positive",
			args{-1, 1},
			1,
		},
		{
			"positive and negative",
			args{1, -1},
			1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MaxInt(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("MaxInt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParseInts(t *testing.T) {
	type args struct {
		str   string
		delim string
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			"Normal",
			args{"12345", ""},
			[]int{1, 2, 3, 4, 5},
		},
		{
			"CSV",
			args{"5,4,3,2,1", ","},
			[]int{5, 4, 3, 2, 1},
		},
		{
			"Hyphens/doubled",
			args{"10-20-100-4096-0-100", "-"},
			[]int{10, 20, 100, 4096, 0, 100},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ParseInts(tt.args.str, tt.args.delim); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseInts() = %v, want %v", got, tt.want)
			}
		})
	}
}
