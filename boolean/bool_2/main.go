package main

import "fmt"

func main() {
	var a int
	var otvet bool

	fmt.Println("Введите число")
	fmt.Scan(&a)
	otvet = a%2 == 1
	fmt.Println(otvet)

}
