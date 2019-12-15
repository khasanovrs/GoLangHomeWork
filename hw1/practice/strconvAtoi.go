package main

import (
	"fmt"
	"log"
	"strconv"
)

func main() {
	str := "12"
	num, err := strconv.Atoi(str)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("В переменной num храниться число", num)
}
