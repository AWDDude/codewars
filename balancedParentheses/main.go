/*
balancedParens 0 -> [""]
balancedParens 1 -> ["()"](1)
balancedParens 2 -> ["()()","(())"] (2 + 1)-1
balancedParens 3 -> ["()()()","(())()","()(())","(()())","((()))"] (3+2+1) -1
balancedParens 3 -> ["()()()()","(())()()","()(())()","()()(())","(()())()", "((()))()","()(()())","()((()))","(()()())","((())())","(()(()))" ,"((()()))","(((())))"] (4+3+2+1) -1
*/

package main

import (
	"fmt"
)

func balancedParens(count int) string {
	switch count {
	case 0:
		return "[\"\"]"
	case 1:
		return "[\"()\"]"
	}

	table := parenPerms(count)

	s := "["
	for i := range perms[count] {
		s += "\"" + perms[count][i].Render() + "\""
	}
	return s + "]"
}

func parenPerms(count int) [][][]*Paren {
	if count == 1 {
		return [][][]*Paren{{{}, {&Paren{}}}}
	}

	table := parenPerms(count - 1)
	row := [][]*Paren{}
	for i := 0; i <= count; i++ {
		group := make([]*Paren, count-i)
		for j := range group {
			for k := range table[i] {
				group[j] = &Paren{Children: table[i][k]}
			}
		}
		row = append(row, group)
	}
	table = append(table, row)
	return table
}

type Paren struct {
	Children []*Paren
}

func (p *Paren) Render() string {
	s := "("
	for i := range p.Children {
		s += p.Children[i].Render()
	}
	s += ")"
	return s
}

type group []*Paren

func (g group) Render() string {
	s := "\""
	for i := range g {
		s += g[i].Render()
	}
	s += "\""
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
