package main

import "fmt"

func main() {
	var num [10]int
	var divs [5]int

	fmt.Println("Digite 10 números inteiros para o vetor Num:")
	for i := 0; i < 10; i++ {
		fmt.Printf("Posição %d: ", i)
		fmt.Scan(&num[i])
	}

	fmt.Println("Digite 5 números inteiros para o vetor Divs:")
	for i := 0; i < 5; i++ {
		fmt.Printf("Posição %d: ", i)
		fmt.Scan(&divs[i])
	}

	fmt.Println("\nResultado:")
	for i := 0; i < 10; i++ {
		fmt.Printf("Número %d:\n", num[i])
		temDivisor := false
		for j := 0; j < 5; j++ {
			if divs[j] != 0 && num[i]%divs[j] == 0 {
				fmt.Printf("    Divisível por %d na posição %d\n", divs[j], j)
				temDivisor = true
			}
		}
		if !temDivisor {
			fmt.Println("    Nenhum divisor encontrado")
		}
	}
}
