package main

import "fmt"

func main () {
	var numero int

	fmt.Printf("Digite um número: ")
	fmt.Scan(&numero)

	if numero%5 == 0 {
		fmt.Println("O número é divisível por 5")
	} else {
		fmt.Println("O número não é divsível por 5")
	}
}