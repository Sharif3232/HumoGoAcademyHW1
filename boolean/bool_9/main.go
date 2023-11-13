//Даны два целых числа: А и В.
// Проверить истинность высказывания.<Хотя бы одно из чисел А и В нечетное>

package main

import "fmt"

func main() {
	var a, b int
	var result bool

	fmt.Println("Введите два целых числа")
	fmt.Scan(&a, &b)
	result = a%2 == 0 || b%2 == 1
	fmt.Println("Ваш ответ:", result)

}
