package main

import (
	"fmt"
)

func main() {
	fmt.Println(CreatePhoneNumber([10]uint{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}))
}

func CreatePhoneNumber(a [10]uint) string {
	return fmt.Sprintf("(%d%d%d) %d%d%d-%d%d%d%d", a[0], a[1], a[2], a[3], a[4], a[5], a[6], a[7], a[8], a[9])
}
