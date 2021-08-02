package main

import "fmt"

func main() {
	var (
		a         int
		b         int
		operation string
	)
	a = 2
	b = 0
	operation = "/"

	switch operation {
	case "*":
		fmt.Println(a * b)
	case "/":
		if b == 0 {
			fmt.Println("You tried to divide by zero you naughty person.")
		} else {
			fmt.Println(a / b)
		}
	case "+":
		fmt.Println(a + b)
	case "-":
		fmt.Println(a - b)
	default:
		fmt.Println("I don't know what " + operation + " means.")
	}
}
