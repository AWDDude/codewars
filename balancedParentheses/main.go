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

func parenPerms(count int) []RowGroups {
	if count == 1 {
		return []RowGroups{
			RowGroups{GroupParens{}},
			RowGroups{GroupParens{&Paren{}}},
		}
	}

	table := parenPerms(count - 1)
	row := RowGroups{NewGroupParens(count)}
	for i := 1; i < count; i++ {
		for j := 0; j < count-i; j++ {
			for k := range table[i] {
				group := NewGroupParens(count - i)
				group[j].Children = table[i][k]
				row = append(row, group)
			}
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

func NewGroupParens(length int) GroupParens {
	gp := make(GroupParens, length)
	for i := range gp {
		gp[i] = &Paren{}
	}
	return gp
}

type GroupParens []*Paren

func (g GroupParens) Render() string {
	s := "\""
	for i := range g {
		s += g[i].Render()
	}
	s += "\""
	return s
}

type RowGroups []GroupParens

func (r RowGroups) Render() string {
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
	// fmt.Println(balancedParens(1))
	// fmt.Println(balancedParens(2))
	// fmt.Println(balancedParens(3))
	fmt.Println(balancedParens(4))
	// fmt.Println(balancedParens(5))
}
