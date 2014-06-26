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
	widthFunc := func(h int) int {
		if h <= int(height) {
			return int(float64(h)/height*base) + minWidth
		} else {
			return int((2*height-float64(h))/height*base) + minWidth
		}
	}
	return JustifyByWidth(SplitLines(tokens, widthFunc), widthFunc, true)
}
