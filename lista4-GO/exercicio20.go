package main

import "fmt"

func main() {
	var jogadas [20]int
	var freq [7]int

	fmt.Println("Digite o resultado de 20 jogadas de dado (1 a 6):")
	for i := 0; i < 20; i++ {
		fmt.Printf("Jogada %d: ", i+1)
		fmt.Scan(&jogadas[i])
		if jogadas[i] < 1 || jogadas[i] > 6 {
			fmt.Println("Valor inválido! Digite entre 1 e 6.")
			i--
			continue
		}
		freq[jogadas[i]]++
	}

	fmt.Println("\nNúmeros sorteados:", jogadas)
	fmt.Println("\nFrequência de cada número:")
	for i := 1; i <= 6; i++ {
		fmt.Printf("Número %d: %d vezes\n", i, freq[i])
	}
}
