package main

import "fmt"

func main () {
	var x float64

	fmt.Printf("Digite o valor de x: ")
	fmt.Scan(&x)

	if x == 2 {
		fmt.Println("Não existe divisão por zero")
	} else {
		funcao := 8 / (2 - x)
		fmt.Println("f(x) =", funcao)
	}	
}