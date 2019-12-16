package main

import "fmt"

func main() {
	var number int
	fmt.Println("Укажите целое число:")
	fmt.Scan(&number)
	if divideBy3(number) {
		fmt.Println("Данное число делится на 3!")
	} else {
		fmt.Println("Данное число не делится на 3!")
	}
}

func divideBy3(number int) bool {
	return number%3 == 0
}
