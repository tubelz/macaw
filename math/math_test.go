package math

import (
	"testing"

	"github.com/veandco/go-sdl2/sdl"
)

func TestSumPoint(t *testing.T) {
	cases := []struct {
		inA  *sdl.Point
		inB  *sdl.Point
		want *sdl.Point
	}{
		{&sdl.Point{1, 2}, &sdl.Point{3, 5}, &sdl.Point{4, 7}},
		{nil, &sdl.Point{3, 5}, &sdl.Point{3, 5}},
		{&sdl.Point{1, 2}, nil, &sdl.Point{1, 2}},
	}
	for _, c := range cases {
		got := SumPoint(c.inA, c.inB)
		if *got != *c.want {
			t.Errorf("SumPoint(%v, %v) == %v, want %v", c.inA, c.inB, got, c.want)
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
		{&FPoint{1.2, 2.1}, &FPoint{2.1, 2.1}, &FPoint{3.3, 4.2}},
	}
	for _, c := range cases {
		got := SumFPoint(c.inA, c.inB)
		if *got != *c.want {
			t.Errorf("SumFPoint(%v, %v) == %v, want %v", c.inA, c.inB, got, c.want)
		}
	}
}

func TestSumPointWithFPoint(t *testing.T) {
	cases := []struct {
		inA  *sdl.Point
		inB  *FPoint
		want *sdl.Point
	}{
		{&sdl.Point{1, 2}, &FPoint{3, 5}, &sdl.Point{4, 7}},
		{nil, &FPoint{3, 5}, &sdl.Point{3, 5}},
		{&sdl.Point{1, 2}, nil, &sdl.Point{1, 2}},
		{&sdl.Point{1, 2}, &FPoint{2.1, 2.1}, &sdl.Point{3, 4}},
		{&sdl.Point{1, 2}, &FPoint{2.8, 2.8}, &sdl.Point{4, 5}},
	}
	for _, c := range cases {
		got := SumPointWithFPoint(c.inA, c.inB)
		if *got != *c.want {
			t.Errorf("SumPointWithFPoint(%v, %v) == %v, want %v", c.inA, c.inB, got, c.want)
		}
	}
}

func TestConvertPointToFPoint(t *testing.T) {
	cases := []struct {
		in   *sdl.Point
		want *FPoint
	}{
		{&sdl.Point{1, 2}, &FPoint{1, 2}},
		{nil, &FPoint{0, 0}},
	}
	for _, c := range cases {
		got := ConvertPointToFPoint(c.in)
		if *got != *c.want {
			t.Errorf("ConvertPointToFPoint(%v) == %v, want %v", c.in, got, c.want)
		}
	}
}

func TestMulFPointWithFloat(t *testing.T) {
	cases := []struct {
		inA  *FPoint
		inB  float32
		want *FPoint
	}{
		{&FPoint{1, 2}, 2, &FPoint{2, 4}},
		{nil, 10, &FPoint{0, 0}},
		{&FPoint{1.1, 2.2}, 2, &FPoint{2.2, 4.4}},
		{&FPoint{1, 2}, 2.1, &FPoint{2.1, 4.2}},
		{&FPoint{1.1, 2.1}, 2.1, &FPoint{1.1 * float32(2.1), 2.1 * float32(2.1)}},
	}
	for _, c := range cases {
		got := MulFPointWithFloat(c.inA, c.inB)
		if *got != *c.want {
			t.Errorf("MulFPointWithFloat(%v, %v) == %v, want %v", c.inA, c.inB, got, c.want)
		}
	}
}

func TestMulPointWithInt(t *testing.T) {

}

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

func TestRound64(t *testing.T) {
	cases := []struct {
		in   float64
		want int64
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
		got := Round64(c.in)
		if got != c.want {
			t.Errorf("Round64(%f) == %d, want %d", c.in, got, c.want)
		}
	}
}

func BenchmarkRound(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Round(10.10)
		Round(10.9)
		Round(-10.10)
		Round(-10.9)
	}
}
