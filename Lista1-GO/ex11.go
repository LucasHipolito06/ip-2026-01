package main

import "fmt"

func main() {

	var numero int

	fmt.Print("Digite um número inteiro: ")
	fmt.Scan(&numero)

	if numero%3 == 0 && numero%5 == 0 {
		fmt.Println("O número é divisível")
	} else {
		fmt.Println("O número não é divisível")
	}
}
