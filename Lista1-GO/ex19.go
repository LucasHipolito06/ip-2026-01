package main

import "fmt"

func main() {

	var numero int

	fmt.Print("Digite um número inteiro: ")
	fmt.Scan(&numero)

	if numero <= 1 {
		fmt.Println("Número inválido")
	} else {

		soma := 0.0

		for i := 1; i <= numero; i++ {
			soma += 1.0 / float64(i)
		}

		fmt.Printf("Soma: %.4f\n", soma)
	}
}
