package main

import "fmt"

func main() {
	var vetor [10]int

	fmt.Println("Digite 10 números inteiros:")
	for i := 0; i < 10; i++ {
		fmt.Printf("Posição %d: ", i)
		fmt.Scan(&vetor[i])
	}

	// Usar map para contar ocorrências
	contagem := make(map[int]int)
	for i := 0; i < 10; i++ {
		contagem[vetor[i]]++
	}

	fmt.Println("\nElementos repetidos:")
	encontrou := false
	for valor, qtd := range contagem {
		if qtd > 1 {
			fmt.Printf("O valor %d aparece %d vezes.\n", valor, qtd)
			encontrou = true
		}
	}

	if !encontrou {
		fmt.Println("Nenhum elemento repetido encontrado.")
	}
}
