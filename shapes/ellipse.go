package shapes

import (
	"math"
)

func parametrizedEllipse(width, height float64, minWidth int) widthFunc {
	return func(y int) int {
		// y ∈ [0, 2*h] and strictly increases, but we need it decreasing
		y64 := math.Floor(2*height - float64(y))
		// Solve the equation for x to get two points:
		// (x/(2*w))^2 = 1 - (y/(2*h))^2
		// -> x = 2*w*sqrt(1 - (y/(2*h))^2)
		x := 2 * width * math.Sqrt(1-math.Pow(y64/(2*height), 2))
		return int(math.Max(float64(minWidth), x+float64(minWidth)))
	}
}

func Ellipse(tokens []string, ratio float64) string {
	// Default parameters try to make it look like a circle
	// Equation of ellipse: (x/(2*w))^2 + (y/(2*h))^2 = 1
	// A_ellipse = π*h*w
	area := float64(TotalLength(tokens)) / 2 // divide by 2 to get both halves
	// A = π*ratio*w^2
	// -> w = sqrt(A/π*ratio^2)
	w := math.Sqrt(area / math.Pi * math.Pow(ratio, 2))
	// magic constant here adjusts for fact that we can't fit perfectly
	// ideally it would be a function of area and minWidth
	h := 0.99 / math.Pow(ratio, 2) * w
	// squash the top and bottom a little bit
	minWidth := 7
	width := parametrizedEllipse(w, h, minWidth)
	return JustifyByWidth(SplitLines(tokens, width), width, true)
}
