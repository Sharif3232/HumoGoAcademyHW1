package main

import "fmt"

func main() {
	var a, b int

	fmt.Println("Введите число:")
	fmt.Scan(&a, &b)

	if a > b {
		fmt.Println("Наибольшее число:", a)
	}
	if b > a {

		fmt.Println("Наибольшее число:", b)
	}
}
