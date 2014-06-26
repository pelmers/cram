package tokenize

type Unimplemented struct{}

func (_ Unimplemented) Tokenize(_ string, in []string) []string {
	return in
}

func (_ Unimplemented) RenameTokens(_ []string, _ int) {
	return
}
