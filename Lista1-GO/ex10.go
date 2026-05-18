package main

import "fmt"

func main() {
	var a, b, c, d int

	fmt.Println("Qual elemento A? ")
	fmt.Scan(&a)

	fmt.Println("Qual o elemento B? ")
	fmt.Scan(&b)

	fmt.Println("Qual o elemento C? ")
	fmt.Scan(&c)

	fmt.Println("Qual o elemento D? ")
	fmt.Scan(&d)

	det := (a * d) - (b * c)

	fmt.Println("O determinante da matriz é: ", det)

}
