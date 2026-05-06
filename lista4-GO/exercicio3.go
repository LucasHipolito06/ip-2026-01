package main

import "fmt"

func main() {
	var vetor [10]int

	fmt.Println("Digite 10 números inteiros:")
	for i := 0; i < 10; i++ {
		fmt.Printf("Número %d: ", i+1)
		fmt.Scan(&vetor[i])
	}

	somaPares := 0
	qtdImpares := 0

	fmt.Println("\na) Números pares digitados:")
	for i := 0; i < 10; i++ {
		if vetor[i]%2 == 0 {
			fmt.Printf("  %d\n", vetor[i])
			somaPares += vetor[i]
		}
	}

	fmt.Printf("\nb) Soma dos números pares: %d\n", somaPares)

	fmt.Println("\nc) Números ímpares digitados:")
	for i := 0; i < 10; i++ {
		if vetor[i]%2 != 0 {
			fmt.Printf("  %d\n", vetor[i])
			qtdImpares++
		}
	}

	fmt.Printf("\nd) Quantidade de números ímpares: %d\n", qtdImpares)
}
