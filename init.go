package cram

import (
	"math/rand"
	"time"
)

var MakeIdentString func(int) string

func init() {
	// initialize some things, make sure keywords and symbols are sorted
	rand.Seed(time.Now().UnixNano())
	usedIdents := make(map[string]bool)
	MakeIdentString = func(n int) string {
		// make random string of n letters
		// if already used, try again
		ident := make([]byte, n)
		for i := 0; i < len(ident); i++ {
			c := rand.Intn(52)
			if c < 26 {
				c += 97
			} else {
				c += 39
			}
			ident[i] = byte(c)
		}
		str := string(ident)
		if !usedIdents[str] {
			usedIdents[str] = true
			return str
		} else {
			return MakeIdentString(n + 1)
		}
	}
}
