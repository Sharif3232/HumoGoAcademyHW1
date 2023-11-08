package main

import "fmt"

const pi = 3.14

func main() {

	var R float32

	fmt.Println("Задайте значени R:")
	fmt.Scan(&R)

	fmt.Println("Ваша длина окружности L: ", 2*pi*R)
	fmt.Println("Ваша площад круга S: ", pi*R*R)

}
