package js

import (
	"github.com/pelmers/cram/tokenize"
	"strings"
	"testing"
)

// Test if tokenizing string gives the expected tokens.
func testTokenization(code string, expected_tokens []string, t *testing.T) {
	tok := NewJSTokenizer()
	tokens := tok.Tokenize(code, []string{""})
	if !tokenize.SlicesEqual(tokens, expected_tokens) {
		t.Error("Expected:", strings.Join(expected_tokens, "|"), "Got", strings.Join(tokens, "|"))
	}
}

func TestFunctionDecl(t *testing.T) {
	code := "function testFunction(param1, param2)"
	expected := []string{"function ", "testFunction", "(", "param1", ",", "param2", ")"}
	testTokenization(code, expected, t)
}

func TestIf(t *testing.T) {
	code := "if (upto < 2) return 1;"
	expected := []string{"if ", "(", "upto", "<", "2", ")", "return 1", ";"}
	testTokenization(code, expected, t)
	code = "if (upto < 2) {return 1;}"
	expected = []string{"if ", "(", "upto", "<", "2", ")", "{", "return 1", ";", "}"}
	testTokenization(code, expected, t)
}

func TestWhile(t *testing.T) {
	code := "while (i    >=  0) i++;"
	expected := []string{"while ", "(", "i", ">=", " 0", ")", "i++", ";"}
	testTokenization(code, expected, t)
}

func TestEmpty(t *testing.T) {
	code := ""
	expected := []string{}
	testTokenization(code, expected, t)
}
