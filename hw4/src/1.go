package main

import "fmt"

type MachineInterface interface {
	getColor() string
	getBrand() string
	setColor(string)
}

type Car struct {
	color string
	brand string
}

func (c *Car) setColor(color string) {
	c.color = color
}

func (c Car) getColor() string {
	return c.color
}
func (c Car) getBrand() string {
	return c.brand
}

type Track struct {
	color   string
	brand   string
	tonnage int64
}

func (t *Track) setColor(color string) {
	t.color = color
}

func (t Track) getColor() string {
	return t.color
}
func (t Track) getBrand() string {
	return t.brand
}

func getColors(machines ...MachineInterface) string {
	colors := ""
	for _, machine := range machines {
		fmt.Println("машина", machine.getBrand(), "имеет цвет", machine.getColor())
		colors += machine.getColor() + "; "
	}
	return colors
}

func setColors(color string, machines ...MachineInterface) {
	for _, machine := range machines {
		machine.setColor(color)
	}
}

func main() {
	car := Car{
		"red",
		"kia",
	}

	track := Track{
		"black",
		"bmv",
		100,
	}

	fmt.Println(car)
	fmt.Println(track)

	fmt.Println(getColors(&car, &track))
	setColors("white", &car, &track)
	fmt.Println(getColors(&car, &track))
}
