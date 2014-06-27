package shapes

import (
	"math"
)

func Trapezoid(tokens []string, ratio float64) string {
	// Trapezoid represented as triangle with flattened top
	// (see triangle.go)
	area := float64(TotalLength(tokens))
	base := math.Sqrt(2 * area * ratio)
	height := 2.0 * area / base
	lid := 0.7 * base // what makes it a trapezoid
	width := parametrizedTriangle(area, base, height, int(lid))
	return JustifyByWidth(SplitLines(tokens, width), width, true)
}
