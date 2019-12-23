package main

import (
	"fmt"
)

func main() {
	var previousPrevious int64
	var previous int64
	var tec int64

	previousPrevious = 0
	previous = 1

	fmt.Println("Значение", 1, "=", 0)
	fmt.Println("Значение", 2, "=", 1)

	for i := 3; i <= 100; i++ {
		tec = previousPrevious + previous
		fmt.Println("Значение", i, "=", tec)
		previousPrevious = previous
		previous = tec
	}

}
