package main

import (
	"fmt"
	"math"
)

func cosAproximado(x float64) float64 {
	resultado := 0.0
	termo := 1.0
	for i := 0; i < 20; i++ {
		if i == 0 {
			resultado += termo
			continue
		}
		termo *= -x * x / float64((2*i-1)*(2*i))
		resultado += termo
	}
	return resultado
}

func main() {
	var x float64
	fmt.Scan(&x)
	aproximado := cosAproximado(x)
	diferenca := aproximado - math.Cos(x)
	fmt.Printf("%.10f\n", aproximado)
	fmt.Printf("%.10f\n", diferenca)
}
