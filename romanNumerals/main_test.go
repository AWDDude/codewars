package main

import (
	"testing"

	. "github.com/onsi/gomega"
)

func Test_romanNumerals(t *testing.T) {
	g := NewGomegaWithT(t)
	g.Expect(romanNumerals()).To(Equal())
}

