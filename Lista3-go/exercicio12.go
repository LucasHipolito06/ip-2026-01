package main

import "fmt"

func main() {
	var x float64
	fmt.Scan(&x)
	resultado := 0.0
	potencia := x
	fatorial := 1.0
	sinal := 1.0
	for i := 1; i <= 20; i++ {
		if i == 1 {
			resultado += x
			continue
		}
		fatorial *= float64(i - 1)
		potencia *= x
		resultado += sinal * potencia / fatorial
		sinal *= -1
	}
	fmt.Printf("%.10f\n", resultado)
}
