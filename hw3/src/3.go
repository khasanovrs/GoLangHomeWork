package main

import (
	"fmt"
	"hw3/stack"
)

func main() {
	stack.Push("Этот текст")
	stack.Push("Будет находиться в стеке")
	stack.Push("До первого обращения к pop")

	fmt.Println(stack.Pop())
	fmt.Println(stack.Pop())

	stack.Push("Добавим еще текста")

	fmt.Println(stack.Pop())
	fmt.Println(stack.Pop())

}
