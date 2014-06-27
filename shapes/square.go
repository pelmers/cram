package shapes

import (
	"math"
)

func Square(tokens []string, ratio float64) string {
	// square has constant width, so this is easy
	// the "area" of the code is the total length of all the tokens
	side := int(ratio * math.Sqrt(float64(TotalLength(tokens))))
	width := func(_ int) int {
		return side
	}
	return JustifyByWidth(SplitLines(tokens, width), width, false)
}
