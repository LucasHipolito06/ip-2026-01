package main

import "fmt"

func main() {
	var idades [50]int

	fmt.Println("Digite 50 idades:")
	for i := 0; i < 50; i++ {
		fmt.Printf("Idade %d: ", i+1)
		fmt.Scan(&idades[i])
	}

	contagem := make(map[int]int)
	for i := 0; i < 50; i++ {
		contagem[idades[i]]++
	}

	moda := idades[0]
	maiorFreq := 0
	for valor, freq := range contagem {
		if freq > maiorFreq {
			maiorFreq = freq
			moda = valor
		}
	}

	fmt.Printf("\nA moda das idades é: %d (aparece %d vezes)\n", moda, maiorFreq)
}
