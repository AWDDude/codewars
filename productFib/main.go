package main

import "fmt"

func main() {
	fmt.Printf("input = %v, output = %v\n", 4895, ProductFib(4895))
	fmt.Printf("input = %v, output = %v\n", 5895, ProductFib(5895))
	fmt.Printf("input = %v, output = %v\n", 74049690, ProductFib(74049690))
	fmt.Printf("input = %v, output = %v\n", 84049690, ProductFib(84049690))
}

func ProductFib(prod uint64) [3]uint64 {
	fibs := []uint64{0, 1}
	var f, n1, n2, p uint64
	i := 2
	for {
		n1 = fibs[i-2]
		n2 = fibs[i-1]
		f = n1 + n2
		fibs = append(fibs, f)
		p = n1 * n2
		if p == prod {
			return [3]uint64{n1, n2, 1}
		}
		if p > prod {
			return [3]uint64{n1, n2, 0}
		}
		i++
	}
}
