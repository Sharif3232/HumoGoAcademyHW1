package main

import "fmt"

func main() {

	var l int

	fmt.Println("Введите растояние L в сантимертрах: ")
	fmt.Scan(&l)

	fmt.Println("Количество полных метров: ", l/100)

}
