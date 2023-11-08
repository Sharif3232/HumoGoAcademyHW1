package main

import "fmt"

func main() {

	var a int
	fmt.Println("Дана длина ребра куба а: ")
	fmt.Scan(&a)
	fmt.Println("V =", a*a*a)
	fmt.Println("S = ", 6*a)

}
