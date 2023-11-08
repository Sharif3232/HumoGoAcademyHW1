package main

import "fmt"

func main() {
	var a int
	var b int
	fmt.Scan(&a)
	fmt.Scan(&b)
	fmt.Println("S =", a*b)
	fmt.Println("P = ", 2*(a+b))
}
