package main

import (
	"testing"

	. "github.com/onsi/gomega"
)

func Test_IntToString(t *testing.T) {
	g := NewGomegaWithT(t)
	g.Expect(intToString(0)).To(Equal("0"))
	g.Expect(intToString(1)).To(Equal("1"))
	g.Expect(intToString(100)).To(Equal("100"))
	g.Expect(intToString(5)).To(Equal("5"))
	g.Expect(intToString(50)).To(Equal("50"))
	g.Expect(intToString(1234)).To(Equal("1234"))
	g.Expect(intToString(-1234)).To(Equal("-1234"))
}

func TestHumanReadableTime(t *testing.T) {
	g := NewGomegaWithT(t)
	g.Expect(HumanReadableTime(0)).To(Equal("00:00:00"))
	g.Expect(HumanReadableTime(59)).To(Equal("00:00:59"))
	g.Expect(HumanReadableTime(60)).To(Equal("00:01:00"))
	g.Expect(HumanReadableTime(90)).To(Equal("00:01:30"))
	g.Expect(HumanReadableTime(3599)).To(Equal("00:59:59"))
	g.Expect(HumanReadableTime(3600)).To(Equal("01:00:00"))
	g.Expect(HumanReadableTime(45296)).To(Equal("12:34:56"))
	g.Expect(HumanReadableTime(86399)).To(Equal("23:59:59"))
	g.Expect(HumanReadableTime(86400)).To(Equal("24:00:00"))
	g.Expect(HumanReadableTime(359999)).To(Equal("99:59:59"))
}
