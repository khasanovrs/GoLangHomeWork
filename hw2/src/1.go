package main

import "fmt"

func main() {
	var number int
	fmt.Println("Укажите целое число:")
	fmt.Scan(&number)
	if canDivide(number, 2) {
		fmt.Println("Данное число четное!")
	} else {
		fmt.Println("Данное число нечетное!")
	}
}

func canDivide(number int, divider int) bool {
	return number%divider == 0
}
