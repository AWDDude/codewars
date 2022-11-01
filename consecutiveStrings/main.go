package main

import (
	"fmt"
	"strings"
)

func main() {
	s := []string{"zone", "abigail", "theta", "form", "libe", "zas"}
	k := 3
	fmt.Printf("input = %v, %v, output = %v", s, k, LongestConsec(s, k))
}

func LongestConsec(strarr []string, k int) string {
	var ls, cs string
	for i := 0; i <= len(strarr)-k; i++ {
		cs = strings.Join(strarr[i:i+k], "")
		if len(cs) > len(ls) {
			ls = cs
		}
	}
	return ls
}
