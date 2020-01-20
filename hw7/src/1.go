package main

import (
	"fmt"
	"time"
)

func main() {
	go spinner(50 * time.Millisecond)
	const n = 42
	time.Sleep(10000 * time.Millisecond)
	fmt.Println("Sleep 10 seconds")
}

func spinner(delay time.Duration) {
	for {
		for _, r := range "-\\|/" {
			fmt.Printf("%c\r", r)
			time.Sleep(delay)
		}
	}
}
