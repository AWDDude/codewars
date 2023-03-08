package main

import "fmt"

func main() {
	// use an implicit tuple to create all our channels in a single line
	chan1, chan2, chan3, chan4, chan5 := make(chan []int), make(chan []int), make(chan []int), make(chan []int), make(chan []int)
	data := []int{48, 12, 10, 5, 34, 92, 404, 56, 56, 151, 216, 12, 396, 474, 162, 45, 658, 723, 222, 86, 12793, 12898, 286, 39, 59, 95, 328, 82, 47, 50, 102, 34, 992, 746, 33, 99, 238, 330, 366, 61, 856, 952, 774, 129, 892, 996, 309450, 12378, 63, 136, 12067, 9873}

	// start the go routines for each of the math operations, link the output
	// of each goroutine to the input of the next
	go add(chan1, chan2)
	go subtract(chan2, chan3)
	go multiply(chan3, chan4)
	go divide(chan4, chan5)

	// sneakily add a `0` to the beginning of the `data` slice
	data = append([]int{0}, data...)
	// send the data slice through the first channel, starting the loop
	chan1 <- data
	// run this loop in the main thread
	// this loop takes input from the "divide" go routine and sends it to the "add" go routine
	// completing the loop
	watch(chan5, chan1)
}

func add(inChan chan []int, outChan chan []int) {
	var data []int
	for {
		// blocking, wait for data to come through the input channel
		data = <-inChan
		// add the first and second items in the slice and save the result to the second item in the slice
		data[1] = data[0] + data[1]
		// blocking, send the modified slice (omitting the first item) through the output channel
		outChan <- data[1:]
	}
}

func subtract(inChan chan []int, outChan chan []int) {
	var data []int
	for {
		// blocking, wait for data to come through the input channel
		data = <-inChan
		// subtract the second item from the first in the slice and save the result to the second item in the slice
		data[1] = data[0] - data[1]
		// blocking, send the modified slice (omitting the first item) through the output channel
		outChan <- data[1:]
	}
}

func multiply(inChan chan []int, outChan chan []int) {
	var data []int
	for {
		// blocking, wait for data to come through the input channel
		data = <-inChan
		// multiply the first and second items in the slice and save the result to the second item in the slice
		data[1] = data[0] * data[1]
		// blocking, send the modified slice (omitting the first item) through the output channel
		outChan <- data[1:]
	}
}

func divide(inChan chan []int, outChan chan []int) {
	var data []int
	for {
		// blocking, wait for data to come through the input channel
		data = <-inChan
		// divide the second item by the first in the slice and save the result to the second item in the slice
		data[1] = data[0] / data[1]
		// blocking, send the modified slice (omitting the first item) through the output channel
		outChan <- data[1:]
	}
}

func watch(inChan chan []int, outChan chan []int) {
	var data []int
	for {
		// blocking, wait for data to come through the input channel
		data = <-inChan
		// print the first item in the slice as a character (utf-8)
		fmt.Printf("%c", data[0])
		// check if we are done
		if len(data) <= 1 {
			break
		}
		// blocking, send the unmodified slice through the output channel
		outChan <- data
	}
}
