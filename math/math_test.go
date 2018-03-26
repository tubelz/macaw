package math

import "testing"

func TestRound(t *testing.T) {
	cases := []struct {
		in   float32
		want int32
	}{
		{3.14, 3},
		{-3.14, -3},
		{2.718, 3},
		{3.141592653589793, 3},
		{-3.141592653589793, -3},
		{-2.718, -3},
		{0, 0},
	}
	for _, c := range cases {
		got := Round(c.in)
		if got != c.want {
			t.Errorf("Round(%f) == %d, want %d", c.in, got, c.want)
		}
	}
}

func TestSumFPoint(t *testing.T) {
	cases := []struct {
		inA  *FPoint
		inB  *FPoint
		want *FPoint
	}{
		{&FPoint{1, 2}, &FPoint{3, 5}, &FPoint{4, 7}},
		{nil, &FPoint{3, 5}, &FPoint{3, 5}},
		{&FPoint{1, 2}, nil, &FPoint{1, 2}},
	}
	for _, c := range cases {
		got := SumFPoint(c.inA, c.inB)
		if *got != *c.want {
			t.Errorf("SumFPoint(%v, %v) == %v, want %v", c.inA, c.inB, got, c.want)
		}
	}
}
