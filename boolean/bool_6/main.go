package main

import "fmt"

func main() {

	var a int
	var b int
	var c int
	var result bool

	fmt.Scan(&a, &b, &c)

	result = a < b && b < c

	fmt.Println(result)
}
