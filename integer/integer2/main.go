package main

import "fmt"

func main() {

	var m int

	fmt.Println("Введите массу m в киллограмах: ")
	fmt.Scan(&m)

	fmt.Println("Ваша масса польных тонн: ", m/1000)

}
