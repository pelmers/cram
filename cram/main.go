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
	"math"
	"os"
	"strings"
)

type Reshaper func([]string, float64) string

// Return the right tokenizer for the filename.
// If error != nil, then we could not pick a tokenizer for it.
func pickTokenizer(filename string) (cram.Tokenizer, error) {
	if strings.HasSuffix(filename, ".js") {
		return js.NewJSTokenizer(), nil
	}
	// default stdin to javascript
	if filename == "stdin" {
		return js.NewJSTokenizer(), nil
	}
	return unimplementedTokenizer{}, errors.New(filename + " is not a supported filetype")
}

// Return a reshaper function for the option selection.
// If option is not matched to a reshaper, return the default: concatenation
func pickReshaper(option string) Reshaper {
	switch option {
	case "square", "box":
		return Square
	case "triangle", "pyramid":
		return Triangle
	case "trapezoid", "volcano":
		return Trapezoid
	case "circle", "ellipse":
		return Ellipse
	}
	// default choice is to just concatenate everything
	return func(tok []string, _ float64) string {
		return strings.Join(tok, "")
	}
}

func main() {
	filename := flag.String("f", "stdin", "Input file name, default to stdin")
	allow_rename := flag.Bool("rename", false, "Allow identifier renaming")
	reserved := flag.String("reserved", "", "Comma separated list of unrenameable identifiers")
	shape := flag.String("s", "none", "Shape to cram code into")
	length := flag.Int("l", 3, "Target length of renamed identifiers")
	ratio := flag.Float64("r", 2.25, "Height:Width ratio (bigger for taller, shorter for wider)")
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
	fmt.Println(reshape(tokens, math.Abs(*ratio)))
	file.Close()
}
