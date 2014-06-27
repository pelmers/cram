package shapes

// alias for function which takes tokens and ratio, returning reshaped text
type Reshaper func([]string, float64) string

// function that takes a line number, returning the desired width
type widthFunc func(int) int
