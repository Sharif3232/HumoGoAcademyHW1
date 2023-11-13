// Даны два целых числа: А и В. Проверить истинность высказывания:
// <Каждое из чисел А и В нечетное>
package main

import "fmt"

func main() {
	var a, b int
	var result, result2 bool

	fmt.Println(" Введите ваше число А и В")
	fmt.Scan(&a, &b)
	result = a%2 == 1
	result2 = b%2 == 1
	fmt.Println(result, result2)

}
