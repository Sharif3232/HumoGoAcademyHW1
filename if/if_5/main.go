package main

import "fmt"

func main() {

	var a, b, c, vet1, vet2 int

	fmt.Println("Введите три числа:")
	fmt.Scan(&a, &b, &c)

	vet1, vet2 = 0, 0

	if a > 0 {
		vet1++
	} else {
		vet2++
	}
	if b > 0 {
		vet1++
	} else {
		vet2++
	}
	if c > 0 {
		vet1++
	} else {
		vet2++
	}
	fmt.Println("Количество положительных чисел:", vet1)
	fmt.Println("Количество отрицательных чисел:", vet2)

}
