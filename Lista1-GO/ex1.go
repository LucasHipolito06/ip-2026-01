package main

import "fmt"

func main() {
	var nota1, nota2, nota3 float64

	fmt.Print("Digite a primeira nota: ")
	fmt.Scan(&nota1)

	fmt.Print("Digite a segunda nota: ")
	fmt.Scan(&nota2)

	fmt.Print("Digite a terceira nota: ")
	fmt.Scan(&nota3)

	media := (nota1 + nota2 + nota3) / 3

	fmt.Printf("Média: %.2f\n", media)

	if media >= 6 {
		fmt.Println("Aluno aprovado")
	} else {
		fmt.Println("Aluno reprovado")
	}
}
