package main

import "fmt"

func main() {
	var vetor [10]float64
	var codigo int

	fmt.Println("Digite 10 números reais para o vetor:")
	for i := 0; i < 10; i++ {
		fmt.Printf("Posição %d: ", i)
		fmt.Scan(&vetor[i])
	}

	for {
		fmt.Print("\nDigite o código (0=sair, 1=ordem direta, 2=ordem inversa): ")
		fmt.Scan(&codigo)

		switch codigo {
		case 0:
			fmt.Println("Programa encerrado.")
			return
		case 1:
			fmt.Println("Vetor na ordem direta:")
			for i := 0; i < 10; i++ {
				fmt.Printf("%.2f ", vetor[i])
			}
			fmt.Println()
		case 2:
			fmt.Println("Vetor na ordem inversa:")
			for i := 9; i >= 0; i-- {
				fmt.Printf("%.2f ", vetor[i])
			}
			fmt.Println()
		default:
			fmt.Println("Código inválido!")
		}
	}
}
