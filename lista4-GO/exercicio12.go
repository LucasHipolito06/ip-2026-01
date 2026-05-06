package main

import "fmt"

func main() {
	var notas [15]int
	var freq [11]int

	fmt.Println("Digite 15 notas dos alunos (0 a 10):")
	for i := 0; i < 15; i++ {
		fmt.Printf("Nota %d: ", i+1)
		fmt.Scan(&notas[i])
		if notas[i] < 0 || notas[i] > 10 {
			fmt.Println("Nota inválida! Digite entre 0 e 10.")
			i--
			continue
		}
		freq[notas[i]]++
	}

	fmt.Println("\n  Nota | Freq. Absoluta | Freq. Relativa")
	fmt.Println("-------|----------------|---------------")
	for i := 0; i <= 10; i++ {
		rel := float64(freq[i]) / 15.0
		fmt.Printf("  %2d   |      %2d        |    %.4f\n", i, freq[i], rel)
	}
}
