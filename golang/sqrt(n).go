package main

import "fmt"

var right float64
var left = 1.0
var average = 0.0

func main() {
	x := 9.0
	fmt.Printf("%f %f", x, Sqrt(x))
}

func Sqrt(number float64) float64 {
	right = number

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
