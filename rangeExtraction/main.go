package main

import "fmt"

func main() {
	i := []int{-6, -3, -2, -1, 0, 1, 3, 4, 5, 7, 8, 9, 10, 11, 14, 15, 17, 18, 19, 20, 21, 22}
	fmt.Printf("input = %v, output = %v\n", i, Solution(i))
}

func Solution(list []int) string {
	var s, r, rr string
	i := 0
	for i <= len(list)-1 {
		r = fmt.Sprint(list[i])
		if i+2 <= len(list)-1 && list[i+2]-list[i] == 2 {
			for i+1 <= len(list)-1 && list[i+1]-list[i] == 1 {
				i++
				rr = fmt.Sprintf("-%v", list[i])
			}
			r += rr
		}
		if i <= len(list)-2 {
			r += ","
		}
		s += r
		i++
	}
	return s
}
