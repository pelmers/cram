package main

import (
	"math"
)

func Square(tokens []string) string {
	// square has constant width, so this is easy
	// the "area" of the code is the total length of all the tokens
	width := int(2 * math.Sqrt(float64(TotalLength(tokens))))
	widthFunc := func(_ int) int {
		return width
	}
	return JustifyByWidth(SplitLines(tokens, widthFunc), widthFunc, false)
}
