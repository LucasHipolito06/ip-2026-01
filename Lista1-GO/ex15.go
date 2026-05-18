package main

import "fmt"

func main() {

	var numero int

	fmt.Print("Digite um número inteiro: ")
	fmt.Scan(&numero)

	for i := 2; i <= numero; i += 2 {
		fmt.Printf("%d^2 = %d\n", i, i*i)
	}
}
