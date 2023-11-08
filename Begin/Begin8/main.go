package main

import "fmt"

func main() {
	var a int
	var b int

	fmt.Println("Введите два числа a и b")

	fmt.Scan(&a, &b)
	fmt.Println("Их средное арифметическое число равен: ", (a+b)/2)

}
