package main

import "fmt"

func main() {

	var bt int

	fmt.Println("Введите количество байтов: ")
	fmt.Scan(&bt)

	fmt.Println("Количество польных килобайтах: ", bt/1024)

}
