// Loop through each digit (right to left) until you find a digit that has a bigger digit in any
// place to the right
// Of the available digits to the right find the next largest one and swap with our original digit
// arrange all remaining digits in ascending order to fill out the rest of the number

// example:
// input: 59884848459853

// return: 59884848483559

package main

// package main

func main() {
	NextBigger(513)
}

func NextBigger(n int) int {
	d := intToDigits(n)

	switch len(d) {
	case 1:
		return -1
	case 2:
		f := digitsToInt([]int{d[1], d[0]})
		if n < f {
			return f
		}
		return -1
	}

	l := len(d)

	for i := l - 1; i > 0; i-- {
		rSlice := sort(d[i:l])

		if ok, j := findIndexOfLarger(rSlice, d[i-1]); ok {
			d[i-1], rSlice[j] = rSlice[j], d[i-1]
			d = append(d[0:i], rSlice...)
			return digitsToInt(d)
		}
	}
	return -1
}

func intToDigits(n int) []int {
	var digits []int

	for n > 0 {
		digits = append([]int{n - ((n / 10) * 10)}, digits...)
		n = n / 10
	}

	return digits
}

func digitsToInt(d []int) int {
	var n int
	l := len(d) - 1
	p := 1

	for i := l; i >= 0; i-- {
		n += d[i] * p
		p = p * 10
	}
	return n
}

func findIndexOfLarger(a []int, n int) (bool, int) {
	for i, v := range a {
		if v > n {
			return true, i
		}
	}
	return false, 0
}

func sort(a []int) []int {
	t := make(map[int]int)
	for _, v := range a {
		t[v]++
	}
	var b int
	r := make([]int, len(a))
	for i := 0; i < 10; i++ {
		for j := 0; j < t[i]; j++ {
			r[b] = i
			b++
		}
	}
	return r
}
