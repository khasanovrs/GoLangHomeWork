package main

import "fmt"

type CarType struct {
	CarType         string // тип автомобиля легковой или грузовой
	CarBrand        string // марка автомобиля
	YearOfIssue     int    // год выпуска
	Trunc           int    // объем багажника/кузова
	EngineRunning   bool   // запущен ли двигатель
	WindowsOpened   bool   // открыты ли окна
	TrunkPercentage int    // насколько заполнен объем багажника
}

func main() {
	car1 := CarType{
		CarType:         "car",
		CarBrand:        "KIA",
		YearOfIssue:     2019,
		Trunc:           150,
		EngineRunning:   false,
		WindowsOpened:   false,
		TrunkPercentage: 10,
	}

	fmt.Println(car1)
}
