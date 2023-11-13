package main

import "fmt"

func main() {
	var a int
	fmt.Println("Введите число:")
	fmt.Scan(&a)

	if a > 0 {
		fmt.Println("Ваша число:", a+1)

	} else if a < 0 {
		fmt.Println("Ваша число:", a-2)
	} else if a == 0 {
		a = 10
		fmt.Println("Ваша число:", a)
	}

}
