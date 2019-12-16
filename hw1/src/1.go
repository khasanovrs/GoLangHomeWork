package main

import (
	"fmt"
	"math"
)

func main() {
	var rub float64
	var dollars float64
	var courseRubDollars float64 = 65.5
	fmt.Println("Укажите сумму в рублях:")
	fmt.Scan(&rub)
	fmt.Println("Вы указали:", rub, "рублей")
	dollars = math.Round(rub / courseRubDollars)
	fmt.Printf("Рассчитанные доллары: %.2f$", dollars)
}
