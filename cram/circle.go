package main

import (
	"math"
)

func Circle(tokens []string) string {
	// A bit misleading: we call it a circle, but define it to be an ellipse
	// This is because characters in a font are not squares.
	// Equation of ellipse: (x/(2*w))^2 + (y/(2*h))^2 = 1
	// A_ellipse = π*h*w
	area := float64(TotalLength(tokens)) / 2 // divide by 2 to get both halves
	// define ratio height/width
	ratio := 0.2
	// A = π*ratio*w^2
	// -> w = sqrt(A/ratio/π)
	w := math.Sqrt(area / ratio / math.Pi)
	h := ratio * w
	// squash the top and bottom a little bit
	minWidth := 5
	widthFunc := func(y int) int {
		// y ∈ [0, 2*h] and strictly increases, but we need it decreasing
		y64 := 2*h - float64(y)
		// Solve the equation for x to get two points:
		// (x/(2*w))^2 = 1 - (y/(2*h))^2
		// -> x = 2*w*sqrt(1 - (y/(2*h))^2)
		x := 2.0 * w * math.Sqrt(1-math.Pow(y64/(2*h), 2.0))
		return minWidth + int(x)
	}
	return JustifyByWidth(SplitLines(tokens, widthFunc), widthFunc, true)
}
