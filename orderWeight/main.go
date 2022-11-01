package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	fmt.Println(OrderWeight("103 123 4444 99 2000"))
	fmt.Println(OrderWeight("2000 10003 1234000 44444444 9999 11 11 22 123"))
	fmt.Println(OrderWeight(""))
}

// an object to store weights and their precalculated sorting values
type weightHash struct {
	hash int
	str  string
}

func newWeightHash(s string) (r weightHash) {
	for i := range s {
		r.hash += int(s[i] - '0')
	}
	r.str = s
	return
}

func OrderWeight(in string) (out string) {
	// split input string into slice of word strings
	words := strings.Fields(in)
	// short circuit empty input
	if len(words) == 0 {
		return ""
	}
	// initialize array of weights
	weights := make([]weightHash, len(words))
	// loop through all words
	for i := range words {
		buf := newWeightHash(words[i])
		// sort weights
		for j := range weights {
			// if the weight object is empty fill it with the new one
			if weights[j].str == "" {
				weights[j] = buf
				break
			}
			// try sorting by hash values first, then alphabetical order
			if (buf.hash < weights[j].hash) || (buf.hash == weights[j].hash && sort.StringsAreSorted([]string{buf.str, weights[j].str})) {
				buf, weights[j] = weights[j], buf
			}
		}
	}

	// concatenate weights into the output string
	for i := range weights {
		out += weights[i].str
		if i < len(weights)-1 {
			out += " "
		}
	}
	return
}
