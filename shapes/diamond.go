package shapes

import (
	"math"
)

func Diamond(tokens []string, ratio float64) string {
	// A diamond is just two triangles stacked atop each other
	// so we divide the area by 2 to get a height for just half the diamond
	area := float64(TotalLength(tokens)) / 2
	base := math.Sqrt(2 * area * ratio)
	height := 2.0 * area / base
	minWidth := 5
	// top and bottom are triangles
	triWidth := parametrizedTriangle(area, base, height, minWidth)
	width := func(h int) int {
		if h <= int(height) {
			return triWidth(h)
		} else {
			// bottom half, subtract h to have a negative slope
			return triWidth(int(2*height - float64(h)))
		}
	}
	return JustifyByWidth(SplitLines(tokens, width), width, true)
}
