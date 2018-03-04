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
	x := a.X + b.X
	y := a.Y + b.Y
	return &sdl.Point{x, y}
}

// SumFPoint sums two FPoints
func SumFPoint(a *FPoint, b *FPoint) *FPoint {
	return &FPoint{a.X + b.X, a.Y + b.Y}
}

// SumPointWithFPoint sums a FPoint to a Point
func SumPointWithFPoint(a *sdl.Point, b *FPoint) *sdl.Point {
	x := float32(a.X) + b.X
	y := float32(a.Y) + b.Y
	return &sdl.Point{Round(x), Round(y)}
}

// ConvertPointToFPoint converts a Point to a FPoint
func ConvertPointToFPoint(a *sdl.Point) *FPoint {
	return &FPoint{float32(a.X), float32(a.Y)}
}

// MulFPointWithFloat multiply a FPoint with a float
func MulFPointWithFloat(a *FPoint, b float32) *FPoint {
	return &FPoint{a.X * b, a.Y * b}
}

// MulPointWithInt multiply a point with an int
func MulPointWithInt(a *sdl.Point, b int32) *sdl.Point {
	return &sdl.Point{a.X * b, a.Y * b}
}

// Round the float point to int
func Round(num float32) int32 {
	if num < 0 {
		return int32(num - 0.5)
	}
	return int32(num + 0.5)
}
