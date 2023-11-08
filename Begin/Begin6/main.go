package main

import "fmt"

func main() {

	var a int
	var b int
	var c int

	fmt.Println("Введите стороны a, b, c используя пробел")
	fmt.Scan(&a)
	fmt.Scan(&b)
	fmt.Scan(&c)
	fmt.Println("Обьем равен V = ", a*b*c)
	fmt.Println("Площад поверхности S = ", 2*(a*b+b*c+a*c))

}
