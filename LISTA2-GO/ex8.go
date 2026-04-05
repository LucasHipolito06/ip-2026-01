package main 

import "fmt"

func main () {
	var numero int

	fmt.Printf("Digite um número: ")
	fmt.Scan(&numero)

	if numero > 20 && numero < 90 {
		fmt.Println("O número esta dentro da faixa")
	} else {
		fmt.Println("O número não esta dentro da faixa")
	}
}