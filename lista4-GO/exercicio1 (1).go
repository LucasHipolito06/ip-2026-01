package main

import "fmt"

func main() {
	var vetor [10]int

	fmt.Println("Digite 10 números inteiros:")
	for i := 0; i < 10; i++ {
		fmt.Printf("Posição %d: ", i)
		fmt.Scan(&vetor[i])
	}

	encontrou := false
	fmt.Println("\nNúmeros superiores a 30:")
	for i := 0; i < 10; i++ {
		if vetor[i] > 30 {
			fmt.Printf("Valor: %d - Posição: %d\n", vetor[i], i)
			encontrou = true
		}
	}

	if !encontrou {
		fmt.Println("Nenhum número superior a 30 foi encontrado.")
	}
}
