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
	return parenPerms(count)[count].Render()
}

func parenPerms(count int) []rowGroups {
	if count == 1 {
		return []rowGroups{
			rowGroups{groupParens{}},
			rowGroups{groupParens{&Paren{}}},
		}
	}

	table := parenPerms(count - 1)
	row := rowGroups{}
	for i := 0; i < count; i++ {
		for k := range table[i] {
			var group groupParens
			for j := 0; j < count-i; j++ {
				group = append(group, &Paren{Children: table[i][k]})
			}
			row = append(row, group)
		}
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

type groupParens []*Paren

func (g groupParens) Render() string {
	s := "\""
	for i := range g {
		s += g[i].Render()
	}
	s += "\""
	return s
}

type rowGroups []groupParens

func (r rowGroups) Render() string {
	s := "["
	for i := range r {
		s += r[i].Render()
		if i < len(r)-1 {
			s += ","
		}
	}
	s += "]"
	return s
}

func main() {
	// perms := Paren{}
	// perms.Children = append(perms.Children, Paren{}, Paren{})
	// fmt.Println(perms.Render())
	// fmt.Println(balancedParens(0))
	fmt.Println(balancedParens(1))
	fmt.Println(balancedParens(2))
	fmt.Println(balancedParens(3))
	// fmt.Println(balancedParens(4))
	// fmt.Println(balancedParens(5))
}
