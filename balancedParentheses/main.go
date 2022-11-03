/*
balancedParens 0 -> [""]
balancedParens 1 -> ["()"]
balancedParens 2 -> ["()()","(())"]
balancedParens 3 -> ["()()()","(())()","()(())","(()())","((()))"]
*/

package main

import "fmt"

func balancedParens(count int) string {
	if count == 0 {
		return "[\"\"]"
	}

	aaap := make([][][]Paren, count)
	aap := [][]Paren{}
	ap := []Paren{}


	for i := 1; i <= count; i++ {
		aap = [][]Paren{}
		for j := 0; j <= i; j++ {
			ap = []Paren{}
			for k := 0; k < i-j; k++ {
				ap[k] = aaap[][][]
			}
			aap[j] = ap
		}
		aaap[i] = aap
	}
	s := "["
	for i := range perms[count] {
		s += "\"" + perms[count][i].Render() + "\""
	}
	return s + "]"
}

type Paren struct {
	Children []Paren
}

func (p Paren) Render() string {
	s := "("
	for i := range p.Children {
		s += p.Children[i].Render()
	}
	s += ")"
	return s
}

func main() {
	// perms := Paren{}
	// perms.Children = append(perms.Children, Paren{}, Paren{})
	// fmt.Println(perms.Render())
	// fmt.Println(balancedParens(0))
	// fmt.Println(balancedParens(1))
	// fmt.Println(balancedParens(2))
	fmt.Println(balancedParens(3))
}
