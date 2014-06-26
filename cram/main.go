package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"github.com/pelmers/cram"
	"github.com/pelmers/cram/js"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

type Reshaper func([]string) string

// Return the right tokenizer for the filename.
// If error != nil, then we could not pick a tokenizer for it.
func pickTokenizer(filename string) (cram.Tokenizer, error) {
	if strings.HasSuffix(filename, ".js") {
		return js.NewJSTokenizer(), nil
	}
	if filename == "stdin" {
		// default stdin to javascript
		return js.NewJSTokenizer(), nil
	}
	return unimplementedTokenizer{}, errors.New(filename + " is not a supported filetype")
}

// Return a reshaper function for the option selection.
// If option is not matched to a reshaper, return the default: concatenation
func pickReshaper(option string) Reshaper {
	switch option {
	case "square":
		return Square
	case "triangle":
		return Triangle
	case "circle":
		return Circle
	}
	// default choice is to just concatenate everything
	return func(tok []string) string {
		return strings.Join(tok, "")
	}
}

func main() {
	filename := flag.String("f", "stdin", "Input file name, default to stdin")
	allow_rename := flag.Bool("ar", false, "Allow identifier renaming")
	reserved := flag.String("r", "", "Comma separated list of reserved identifier names (won't be renamed)")
	shape := flag.String("s", "none", "Shape to transform code into")
	length := flag.Int("l", 3, "Target length of renamed identifiers")
	flag.Parse()
	var file *os.File
	if *filename == "stdin" {
		file = os.Stdin
	} else {
		var err error
		file, err = os.Open(*filename)
		if err != nil {
			log.Fatal(err)
		}
	}
	reader := bufio.NewReader(file)
	code, err := ioutil.ReadAll(reader)
	if err != nil {
		log.Fatal(err)
	}
	tok, err := pickTokenizer(*filename)
	if err != nil {
		log.Fatal(err)
	}
	tokens := tok.Tokenize(string(code), strings.Split(*reserved, ","))
	if *allow_rename {
		tok.RenameTokens(tokens, *length)
	}
	reshape := pickReshaper(*shape)
	fmt.Println(reshape(tokens))
	file.Close()
}
