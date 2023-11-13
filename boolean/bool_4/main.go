package main

import "fmt"

func main() {
	var a, b int
	var otvet bool

	fmt.Println("Введите два числа")

	fmt.Scan(&a, &b)
	otvet = a > 2 && b <= 3
	fmt.Println(otvet)
}
