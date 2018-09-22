// Package math provides structures and functions to handle mathematical functions.
// Example of use: cartesian functions with different types, round numbers, conversion of number types.
package math

import (
	"github.com/veandco/go-sdl2/sdl"
)

// FPoint is the float point
type FPoint struct {
	X float32
	Y float32
}

// SumPoint sums two Points
func SumPoint(a *sdl.Point, b *sdl.Point) *sdl.Point {
	if a == nil {
		return b
	} else if b == nil {
		return a
	}
	x := a.X + b.X
	y := a.Y + b.Y
	return &sdl.Point{x, y}
}

// SumFPoint sums two FPoints
func SumFPoint(a *FPoint, b *FPoint) *FPoint {
	if a == nil {
		return b
	} else if b == nil {
		return a
	}
	return &FPoint{a.X + b.X, a.Y + b.Y}
}

// SumPointWithFPoint sums a Point to a FPoint. It loses precision. The result is rounded.
func SumPointWithFPoint(a *sdl.Point, b *FPoint) *sdl.Point {
	if b == nil {
		return a
	} else if a == nil {
		return &sdl.Point{Round(b.X), Round(b.Y)}
	}
	roundedB := &sdl.Point{Round(b.X), Round(b.Y)}
	return SumPoint(a, roundedB)
}

// ConvertPointToFPoint converts a Point to a FPoint
func ConvertPointToFPoint(a *sdl.Point) *FPoint {
	if a == nil {
		return &FPoint{0, 0}
	}
	return &FPoint{float32(a.X), float32(a.Y)}
}

// MulFPointWithFloat multiply a FPoint with a float
func MulFPointWithFloat(a *FPoint, b float32) *FPoint {
	if a == nil {
		return &FPoint{0, 0}
	}
	return &FPoint{a.X * b, a.Y * b}
}

// MulPointWithInt multiply a point with an int
func MulPointWithInt(a *sdl.Point, b int32) *sdl.Point {
	if a == nil {
		return &sdl.Point{0, 0}
	}
	return &sdl.Point{a.X * b, a.Y * b}
}

// Round rounds the float point to int
func Round(num float32) int32 {
	if num < 0 {
		return int32(num - 0.5)
	}
	return int32(num + 0.5)
}

// Round64 rounds the float point to int
func Round64(num float64) int64 {
	if num < 0 {
		return int64(num - 0.5)
	}
	return int64(num + 0.5)
}
