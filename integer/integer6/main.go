package main

import "fmt"

func main() {

	var a int

	fmt.Println("Введите двузначное число:")
	fmt.Scan(&a)
	fmt.Println("Левая сторона вашего числа:", a/10, "Правая цифра вашего числа:", a%10)

}
