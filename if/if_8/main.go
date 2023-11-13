package main

import "fmt"

func main() {

	var a, b int

	fmt.Println("Введите два числа")
	fmt.Scan(&a, &b)

	if a >= b {
		fmt.Println("Наибольшее число:","\n", a, b)	
	} else {
	
		fmt.Println("Наибольшее число:","\n", b, a)

	}
}