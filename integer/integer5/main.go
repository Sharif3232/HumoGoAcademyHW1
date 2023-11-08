package main

import "fmt"

func main() {

	var a, b int
	fmt.Println("Введите целые положительные числа a и b: ")
	fmt.Scan(&a, &b)

	fmt.Println("Незанятая часть отрезка А: ", a%b)

}
