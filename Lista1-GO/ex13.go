package main

import "fmt"

func main() {

	var nota float64

	fmt.Print("Qual a nota do aluno? ")
	fmt.Scan(&nota)

	if nota >= 9 {
		fmt.Printf("NOTA = %.1f Conceito A\n", nota)
	} else if nota >= 7.5 {
		fmt.Printf("NOTA = %.1f Conceito B\n", nota)
	} else if nota >= 6 {
		fmt.Printf("NOTA = %.1f Conceito C\n", nota)
	} else {
		fmt.Printf("NOTA = %.1f Conceito D\n", nota)
	}
}
