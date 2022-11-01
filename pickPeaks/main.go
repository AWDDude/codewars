package main

import "fmt"

func main() {
	fmt.Println(PickPeaks([]int{1, 2, 3, 6, 4, 1, 2, 3, 2, 1}))             //{[3 7] [6 3]}
	fmt.Println(PickPeaks([]int{3, 2, 3, 6, 4, 1, 2, 3, 2, 1, 2, 3}))       //{[3 7] [6 3]}
	fmt.Println(PickPeaks([]int{3, 2, 3, 6, 4, 1, 2, 3, 2, 1, 2, 2, 2, 1})) //{[3 7 10] [6 3 2]}
	fmt.Println(PickPeaks([]int{2, 1, 3, 1, 2, 2, 2, 2, 1}))                //{[2 4] [3 2]}
	fmt.Println(PickPeaks([]int{2, 1, 3, 1, 2, 2, 2, 2}))                   //{[2] [3]}
}

type PosPeaks struct {
	Pos   []int
	Peaks []int
}

func PickPeaks(array []int) PosPeaks {
	var r PosPeaks
	var ok bool
	var j int
	for i := 1; i < len(array)-1; i++ {
		if array[i] > array[i-1] && array[i] >= array[i+1] {
			if ok, j = findTrough(array[i:]); ok {
				r.Pos = append(r.Pos, i)
				r.Peaks = append(r.Peaks, array[i])
				i += j
			}
		}
	}
	return r
}

func findTrough(a []int) (bool, int) {
	ok := false
	lowest := 1
	for i := 1; i < len(a); i++ {
		if a[i] < a[i-1] {
			ok = true
		}
		if a[i] < a[lowest] {
			lowest = i
		}
		if a[i] > a[i-1] {
			break
		}
	}
	return ok, lowest
}
