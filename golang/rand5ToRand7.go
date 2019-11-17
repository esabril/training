package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())

	for i := 0; i < 10; i++ {
		fmt.Printf("%d ", Rand7())
	}
}

func Rand5() int {
	return rand.Intn(5)
}

func Rand7() int {
	result := 0

	for true {
		result = Rand5() + 5*Rand5()

		if result <= 20 {
			break;
		}
	}

	return 1 + result%7;
}
