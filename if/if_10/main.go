package main

import "fmt"

func main() {
	var a, b, sum int

	fmt.Println("Введите два числа:")
	fmt.Scan(&a, &b)

	if  a != b {
		
		sum = a + b
		fmt.Println("Ваши значение не равны, новые значение A=",sum, "B=", sum) 
	} else if a == b {
		a = 0
		b = 0
		fmt.Println("Значение равны, новые значение:")
		fmt.Println("A=",a, "B=", b)

	}
}