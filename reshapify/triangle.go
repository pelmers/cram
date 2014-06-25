package main

import (
	"math"
)

func Triangle(tokens []string) string {
	// area of triangle = 0.5 * base * height
	area := float64(TotalLength(tokens))
	// define conversion ratio height:base (since letters are taller than wide)
	heightBaseRatio := 0.4
	// area = 0.5 * base * heightBaseRatio * base
	// -> base = sqrt(2*area / heightBaseRatio)
	base := math.Sqrt(2 * area / heightBaseRatio)
	// height = 2*area / base
	height := 2.0 * area / base
	// minimum width of a line (flattens the top of the triangle)
	minWidth := 10
	// use this to define a widthFunc
	widthFunc := func(h int) int {
		return int(float64(h)/height*base) + minWidth
	}
	return JustifyByWidth(SplitLines(tokens, widthFunc), widthFunc, true)
}
