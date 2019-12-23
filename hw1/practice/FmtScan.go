package main

import "fmt"

func main() {
	var name string
	var age int
	fmt.Println("Как тебя зовут?")
	fmt.Scan(&name)
	fmt.Println("А сколько тебе лет?")
	fmt.Scan(&age)
	if age > 30 {
		fmt.Println(name, "- ты старый :(")
	} else {
		fmt.Println(name, "- ты еще молод :)")
	}

}
