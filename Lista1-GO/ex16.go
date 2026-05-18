package main

import "fmt"

func main() {

	var salario float64

	fmt.Print("Qual o salário inicial? ")
	fmt.Scan(&salario)

	if salario <= 300 {
		fmt.Printf("O novo salário é: R$ %.2f\n", salario*1.5)
	} else {
		fmt.Printf("O novo salário é: R$ %.2f\n", salario*1.3)
	}
}
