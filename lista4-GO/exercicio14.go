package main

import "fmt"

func main() {
	var vetor1 [10]int
	var vetor2 [10]int
	var resultado [20]int

	fmt.Println("Digite 10 inteiros para o vetor 1:")
	for i := 0; i < 10; i++ {
		fmt.Printf("Posição %d: ", i)
		fmt.Scan(&vetor1[i])
	}

	fmt.Println("Digite 10 inteiros para o vetor 2:")
	for i := 0; i < 10; i++ {
		fmt.Printf("Posição %d: ", i)
		fmt.Scan(&vetor2[i])
	}

	for i := 0; i < 10; i++ {
		resultado[2*i] = vetor1[i]
		resultado[2*i+1] = vetor2[i]
	}

	fmt.Println("\nVetor resultante da intercalação:")
	fmt.Println(resultado)
}
