package main

import "fmt"

func main() {
	var vetor [100]int

	for i := 0; i < 100; i++ {
		vetor[i] = 2*i + 1
	}

	fmt.Println("Os 100 primeiros números ímpares:")
	for i := 0; i < 100; i++ {
		fmt.Printf("%d ", vetor[i])
	}
	fmt.Println()
}
