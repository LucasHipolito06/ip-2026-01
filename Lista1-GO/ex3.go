package main

import (
	"fmt"
	"math"
)

func main() {
	var n1, n2, n3 int

	fmt.Print("Digite o primeiro número: ")
	fmt.Scan(&n1)
	fmt.Print("Digite o segundo número: ")
	fmt.Scan(&n2)
	fmt.Print("Digite o terceiro número: ")
	fmt.Scan(&n3)

	if n1 <= 0 || n1 > 9 || n2 > 9 || n3 > 9 {
		fmt.Println("DÍGITO INVÁLIDO")
	} else {
		numero := n1*100 + n2*10 + n3
		quadrado := math.Pow(float64(numero), 2)
		fmt.Println(numero, ",", quadrado)
	}
}
