package main

import "fmt"

func main() {

	var a, b, c int

	fmt.Print("Coeficiente A: ")
	fmt.Scan(&a)

	fmt.Print("Coeficiente B: ")
	fmt.Scan(&b)

	fmt.Print("Coeficiente C: ")
	fmt.Scan(&c)

	delta := b*b - 4*a*c

	fmt.Println("Delta:", delta)
}
