package main

import "fmt"

func main() {
	var alturas [10]float64

	fmt.Println("Digite a altura de 10 atletas:")
	soma := 0.0
	for i := 0; i < 10; i++ {
		fmt.Printf("Altura do atleta %d: ", i+1)
		fmt.Scan(&alturas[i])
		soma += alturas[i]
	}

	media := soma / 10.0

	fmt.Printf("\nMédia das alturas: %.2f\n", media)
	fmt.Println("Atletas com altura acima da média:")
	for i := 0; i < 10; i++ {
		if alturas[i] > media {
			fmt.Printf("Atleta %d - Altura: %.2f\n", i+1, alturas[i])
		}
	}
}
