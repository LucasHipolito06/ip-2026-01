package main

import "fmt"

func main() {
	preco := 6.0
	maxLucro := -1e18
	precoMax := 0.0
	ingressosMax := 0

	for i := 0; ; i++ {
		p := 6.0 - 0.6*float64(i)
		if p < 1.0 {
			break
		}
		ingressos := 130 + 30*i
		lucro := p*float64(ingressos) - 300.0
		fmt.Printf("%.2f %d %.2f\n", p, ingressos, lucro)
		if lucro > maxLucro {
			maxLucro = lucro
			precoMax = p
			ingressosMax = ingressos
		}
	}
	fmt.Printf("%.2f %.2f %d\n", maxLucro, precoMax, ingressosMax)
}
