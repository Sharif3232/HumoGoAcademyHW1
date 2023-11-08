package main

import "fmt"

func main() {

	var a, b int

	fmt.Println("Введите целые числа a и b:")
	fmt.Scan(&a, &b)

	fmt.Println("Отрезков b размещенных на отрезке a: ", a/b)

}
