package main

import "fmt"

type phones []int

func (p phones) ViewList() {
	for i, phone := range p {
		fmt.Printf("\t %v) %v \n", i, phone)
	}
}

type contact struct {
	name        string
	phoneMain   int
	phoneSecond int
}

func main() {
	addressBook := make(map[string]phones)

	addressBook["Миша"] = phones{78293467382}
	addressBook["Никита"] = phones{89167253764, 89635437382}

	for name, ph := range addressBook {
		fmt.Println(name)
		ph.ViewList()
	}

	phoneBook := make(map[string]contact)
	phoneBook["Миша"] = contact{
		name:        "Миша",
		phoneMain:   78293467382,
		phoneSecond: 89167253764,
	}

	fmt.Println(phoneBook)
}
