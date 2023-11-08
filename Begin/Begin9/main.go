package main

import (
    "fmt"
    "math"
)
func main() {
	var a float64
	var b float64

	fmt.Println("Введите значение a и b через запятой:")
	fmt.Scan(&a, &b)
	fmt.Println("Ваша средее геометрическое число:", math.Sqrt(a + b))




}