package main

import "fmt"

func main() {
	var p float64 = 1
	var i float64 = 1
	var t float64 = 5 * 12
	var sum float64

	fmt.Println("Укажите сумму вклада:")
	fmt.Scan(&p)

	fmt.Println("Укажите годовой процент:")
	fmt.Scan(&i)
	sum = (p * i * t) / (365 * 100)
	fmt.Printf("Через 5 лет годовой доход: %.1f\n", sum)
	fmt.Printf("Сумма вклада через 5 лет: %.1f", p+sum)

}
