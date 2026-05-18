package main

import (
	"fmt" ; "math"
)

func main() {
	var numero float64

	fmt.Print("Digite um número: ")
	fmt.Scan(&numero)

	if numero >= 0 {
		fmt.Printf("Raiz quadrada de %.2f: %.2f\n", numero, math.Sqrt(numero))
	} else {
		fmt.Printf("Quadrado de %.2f: %.2f\n", numero, numero*numero)
	}
}