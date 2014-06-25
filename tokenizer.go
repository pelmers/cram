package reshapify

type Tokenizer interface {
	// Return a slice of tokens as strings between which any amount of whitespace
	// can be inserted.
	Tokenize(string, []string) []string
	// Rename non-reserved tokens in place with given target variable length
	RenameTokens([]string, int)
}
