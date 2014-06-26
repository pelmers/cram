package main

type unimplementedTokenizer struct{}

func (_ unimplementedTokenizer) Tokenize(_ string, in []string) []string {
	return in
}

func (_ unimplementedTokenizer) RenameTokens(_ []string, _ int) {
	return
}
