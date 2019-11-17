package main

import "fmt"

func main() {
	x := 2.0
	fmt.Printf("%f %f", x, Sqrt(x))
}

func Sqrt(number float64) float64 {
	left := 1.0
	right := number
	average := 0.0

	for (right - left) > 1e-8 {
		average = left + (right - left)/2

		if average * average > number {
			right = average
		} else if average * average < number {
			left = average
		} else {
			left = average
			break;
		}
	}

	return left
}
