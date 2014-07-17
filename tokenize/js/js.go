package js

import (
	"fmt"
	"github.com/pelmers/cram/tokenize"
	"sort"
	"strings"
	"unicode"
)

type JSTokenizer struct {
	symbols  []rune
	keywords []string
	reserved []string
}

func NewJSTokenizer() *JSTokenizer {
	return &JSTokenizer{
		// symbols
		[]rune{'~', '?', ':', '!', '+', '-', '*', '/', '=', '(', ')', '[', ']', '{', '}', '<', '>', ';', ',', '.', '"', '\''},
		// keywords
		[]string{
			"break", "case", "comment", "continue", "default", "delete", "do", "else",
			"export", "for", "function", "if", "import", "in", "instanceof", "label", "new",
			"return", "switch", "this", "typeof", "var", "void", "while", "with",
		},
		// reserved
		[]string{
			"===", "==", "!==", "!=", "console", "document", "window", "length", "push", "hasOwnProperty",
			">=", "<=", "+=", "-=", "reverse", "shift", "Math", "undefined", "null", "true", "false",
		},
	}
}

func (tok *JSTokenizer) firstSymbol(text string) (int, rune) {
	// Return the index of the first symbol in text string
	for i, t := range text {
		if unicode.IsSpace(t) {
			return i, t
		}
		if tokenize.SearchRunes(tok.symbols, t) != -1 {
			return i, t
		}
	}
	// symbol not found, return special end code
	return len(text) - 1, rune(0)
}

// Return true if t is either a space or a symbol
func (tok *JSTokenizer) isSymbol(t string) bool {
	if len(t) != 1 {
		return false
	}
	r := rune(t[0])
	return tokenize.SearchRunes(tok.symbols, r) != -1
}

// Return whether t is white space or empty
func (tok *JSTokenizer) isSpace(t string) bool {
	for _, r := range t {
		if !unicode.IsSpace(r) {
			return false
		}
	}
	return true
}

// Return whether t is a javascript reserved word
func (tok *JSTokenizer) isKw(t string) bool {
	return tokenize.BSearchStrings(tok.keywords, strings.TrimSpace(t)) != -1
}

// Return whether the identifier t is reserved
func (tok *JSTokenizer) isReserved(t string) bool {
	return tokenize.BSearchStrings(tok.reserved, strings.TrimSpace(t)) != -1
}

// Return whether the string is quoted like a JS string
func (tok *JSTokenizer) isQuoted(a string) bool {
	return (strings.HasPrefix(a, "'") && strings.HasSuffix(a, "'")) || (strings.HasPrefix(a, "\"") && strings.HasSuffix(a, "\""))
}

// Arbitrarily rename non-reserved, non-keyword, non-symbol tokens in place
func (tok *JSTokenizer) RenameTokens(tokens []string, length int) {
	// first pass: create definition mapping
	defs := make(map[string]string)
	for i, t := range tokens {
		// rename if it is not a keyword, not a string, not reserved, and not a symbol
		if tok.isKw(t) || tok.isSymbol(t) || tok.isSpace(t) || tok.isReserved(t) || tokenize.IsNum(t) || tok.isQuoted(t) {
			continue
		}
		// if it comes after a dot, also check the part before the dot
		if strings.TrimSpace(tokens[i-1]) == "." {
			t := tokens[i-2]
			if tok.isKw(t) || tok.isSymbol(t) || tok.isSpace(t) || tok.isReserved(t) || tokenize.IsNum(t) || tok.isQuoted(t) {
				continue
			}
		}
		if strings.HasPrefix(t, "return") {
			ident := strings.TrimSpace(strings.TrimPrefix(t, "return"))
			defs[ident] = tokenize.MakeIdentString(length)
			defs[t] = "return " + defs[ident]
		} else if strings.HasSuffix(t, "++") {
			ident := strings.TrimSuffix(t, "++")
			defs[ident] = tokenize.MakeIdentString(length)
			defs[t] = defs[ident] + "++"
		} else if strings.HasSuffix(t, "--") {
			ident := strings.TrimSuffix(t, "--")
			defs[ident] = tokenize.MakeIdentString(length)
			defs[t] = defs[ident] + "--"
		} else {
			defs[t] = tokenize.MakeIdentString(length)
		}
	}
	// second pass: replace tokens with their renames
	for i, _ := range tokens {
		if ren, ok := defs[tokens[i]]; ok && len(tokens[i]) > 0 {
			tokens[i] = ren
		}
	}
}

func (tok *JSTokenizer) Tokenize(code string, _reserved []string) []string {
	tok.reserved = append(tok.reserved, _reserved...)
	sort.Strings(tok.reserved)
	sort.Strings(tok.keywords)
	tokens := make([]string, 0)
	for {
		nsIndex, nsRune := tok.firstSymbol(code)
		if nsRune == rune(0) {
			break
		}
		tokens = append(tokens, code[:nsIndex])
		if !unicode.IsSpace(nsRune) {
			tokens = append(tokens, string(nsRune))
		}
		code = code[nsIndex+1:]
		// check for the speshul cases
		switch nsRune {
		// add the rest of the string as a token
		// TODO: support escaped quotes
		case '"', '\'':
			nsIndex = strings.IndexRune(code, nsRune)
			if nsIndex >= 0 {
				tokens[len(tokens)-1] = fmt.Sprintf("%c%s%c",
					nsRune, code[:nsIndex], nsRune)
				// advance the code to the close quote character
				code = code[nsIndex+1:]
			}
		case '+', '-':
			// check for unary -- or ++
			if rune(code[0]) == nsRune {
				// move code forward 1 rune
				code = code[1:]
				// delete the - or + we just added
				tokens = tokens[:len(tokens)-1]
				// add -- or ++ to either the end of the previous token
				// TODO: or the start of the next one
				// maybe I should preprocess prefix and turn them into postfix?
				tokens[len(tokens)-1] += string([]rune{nsRune, nsRune})
			}
			// check for += and -=
			if code[0] == '=' {
				tokens[len(tokens)-1] = string(nsRune) + "="
				code = code[1:]
			}
		case '/':
			// check for single-line comment
			if code[0] == '/' {
				// ignore the rest of the line and remove the / token we added
				tokens = tokens[:len(tokens)-1]
				code = code[strings.IndexRune(code, '\n')+1:]
			}
			if code[0] == '*' {
				// skip the codepoint forward until the "*/" end comment
				tokens = tokens[:len(tokens)-1]
				code = code[strings.Index(code, "*/")+2:]
			}
			// if the previous token is a symbol, this is a regexp (I think)
			if tok.isSymbol(tokens[len(tokens)-1]) {
				// find the next symbol after the next / restart there
				nextSlash := strings.Index(code, "/")
				nextLine := strings.Index(code, "\n")
				next, r := tok.firstSymbol(code[nextSlash+1:])
				next += nextSlash + 1
				if r != rune(0) && nextSlash < nextLine {
					tokens[len(tokens)-1] = "/" + code[:next]
					code = code[next:]
				}
			}
		case '.':
			// check for floating point number: if previous token was a number and next one is too
			beforeDot := tokens[len(tokens)-2]
			if tokenize.IsNum(beforeDot) {
				pos, _ := tok.firstSymbol(code)
				nextTok := code[:pos]
				if tokenize.IsNum(nextTok) {
					tokens[len(tokens)-2] = fmt.Sprintf("%s.%s", beforeDot, nextTok)
					tokens[len(tokens)-1] = ""
					code = code[pos:]
				}
			}
		case '=':
			if code[0] == '=' && code[1] == '=' {
				tokens[len(tokens)-1] = "==="
				code = code[2:]
			} else if code[0] == '=' {
				tokens[len(tokens)-1] = "=="
				code = code[1:]
			}
		case '!':
			if code[0] == '=' && code[1] == '=' {
				tokens[len(tokens)-1] = "!=="
				code = code[2:]
			} else if code[0] == '=' {
				tokens[len(tokens)-1] = "!="
				code = code[1:]
			}
		case '<', '>':
			if code[0] == '=' {
				tokens[len(tokens)-1] = string(nsRune) + "="
				code = code[1:]
			}
		}
		code = strings.TrimSpace(code)
	}
	// go through tokens and hardcode a space before and after any keywords
	for i, _ := range tokens[:len(tokens)-1] {
		if tok.isKw(tokens[i]) || tokenize.IsNum(tokens[i]) {
			if !tok.isSymbol(tokens[i+1]) {
				tokens[i] = tokens[i] + " "
				if tokens[i] == "return " {
					tokens[i] += tokens[i+1]
					if i+2 < len(tokens) {
						if tok.isSymbol(tokens[i+2]) {
							tokens[i+1] = ""
						} else {
							tokens[i+1] = " "
						}
					}
				}
			}
			if i > 0 && !tok.isSymbol(tokens[i-1]) {
				tokens[i] = " " + tokens[i]
			}
		}
	}
	return tokens
}
