package shapes

import (
	"math"
)

func Trapezoid(tokens []string, ratio float64) string {
	// lol this is just like Triangle but we increase minWidth some
	// (see triangle.go)
	area := float64(TotalLength(tokens))
	base := math.Sqrt(2 * area * ratio)
	height := 2.0 * area / base
	lid := 0.7 * base // what makes it a trapezoid
	// use this to define a widthFunc
	widthFunc := func(h int) int {
		return int(float64(h)/height*base + lid)
	}
	return JustifyByWidth(SplitLines(tokens, widthFunc), widthFunc, true)
}
