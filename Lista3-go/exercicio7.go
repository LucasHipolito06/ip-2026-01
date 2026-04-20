package main

import "fmt"

func main() {
	const sentinela = 30000
	var numero int
	var soma, quantidade, pares int
	var somaPares int
	var maior, menor int
	var impares int
	primeiro := true

	for {
		fmt.Scan(&numero)
		if numero == sentinela {
			break
		}
		if primeiro {
			maior = numero
			menor = numero
			primeiro = false
		}
		if numero > maior {
			maior = numero
		}
		if numero < menor {
			menor = numero
		}
		soma += numero
		quantidade++
		if numero%2 == 0 {
			somaPares += numero
			pares++
		} else {
			impares++
		}
	}

	media := 0.0
	if quantidade > 0 {
		media = float64(soma) / float64(quantidade)
	}
	mediaPares := 0.0
	if pares > 0 {
		mediaPares = float64(somaPares) / float64(pares)
	}
	percentualImpares := 0.0
	if quantidade > 0 {
		percentualImpares = float64(impares) * 100 / float64(quantidade)
	}

	fmt.Println(soma)
	fmt.Println(quantidade)
	fmt.Printf("%.2f\n", media)
	fmt.Println(maior)
	fmt.Println(menor)
	fmt.Printf("%.2f\n", mediaPares)
	fmt.Printf("%.2f\n", percentualImpares)
}
