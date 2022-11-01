package kata

func DigitalRoot(n int) int {
	for n > 9 {
		n = sum(intToDigits(n))
	}

	return n
}

func intToDigits(n int) []int {
	var digits []int

	for n > 0 {
		digits = append([]int{n - ((n / 10) * 10)}, digits)
		n = n / 10
	}

	return digits
}

func sum(x []int) int {
	var r int

	for _, v := range x {
		r += v
	}

	return r
}
