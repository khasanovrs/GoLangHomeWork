package main

import (
	"fmt"
	"math"
)

func main() {
	var a float64 = 1
	var b float64 = 1
	var c float64

	fmt.Println("Введите значение катета a:")
	fmt.Scan(&a)
	fmt.Println("Введите значение катета b:")
	fmt.Scan(&b)
	c = math.Sqrt(math.Pow(a, 2) + math.Pow(b, 2))
	fmt.Println("Площадь треугольника:", (a*b)/2)
	fmt.Println("Гипотенуза треугольника:", c)
	fmt.Println("Периметр треугольника:", a+b+c)
}
