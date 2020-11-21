package aoc

import "testing"

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
