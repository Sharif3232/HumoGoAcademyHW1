package main

import (
	"fmt"
	"math"
)

func main() {

	var a, b float64

	fmt.Println("Введите значение a и b через запятую:")
	fmt.Scan(&a, &b)

	fmt.Println("Сумма квадратов a и b равно:", math.Sqrt(a+b))
	fmt.Println("Разность квадратов a и b:", math.Sqrt(a-b))
	fmt.Println("Произведение квадратов a и b:", math.Sqrt(a*b))
	fmt.Println("Частное квадратов a и b:", math.Sqrt(a/b))

}
