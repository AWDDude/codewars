package main

import "fmt"

func main() {
	fmt.Printf("input = %v, output = %v\n", 0, CountBits(0))
	fmt.Printf("input = %v, output = %v\n", 4, CountBits(4))
	fmt.Printf("input = %v, output = %v\n", 7, CountBits(7))
	fmt.Printf("input = %v, output = %v\n", 9, CountBits(9))
	fmt.Printf("input = %v, output = %v\n", 10, CountBits(10))

}

func CountBits(n uint) int {
	var c int
	for _, v := range fmt.Sprintf("%b", n) {
		if v == '1' {
			c++
		}
	}
	return c
}
