package math

import "testing"

func TestRound(t *testing.T) {
	cases := []struct {
		in float32
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
			t.Errorf("Round(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}
