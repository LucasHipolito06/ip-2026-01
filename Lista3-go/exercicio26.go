package main

import "fmt"

func main() {
	resultado := 0.0
	numerador := 100.0
	fatorial := 1.0
	for i := 0; i < 20; i++ {
		if i > 0 {
			fatorial *= float64(i)
		}
		resultado += numerador / fatorial
		numerador--
	}
	fmt.Printf("%.10f\n", resultado)
}
