package main

import "fmt"

func main() {
	var number int
	fmt.Println("Укажите целое число:")
	fmt.Scan(&number)
	if isNumberEven(number) {
		fmt.Println("Данное число четное!")
	} else {
		fmt.Println("Данное число нечетное!")
	}
}

func isNumberEven(number int) bool {
	return number%2 == 0
}
