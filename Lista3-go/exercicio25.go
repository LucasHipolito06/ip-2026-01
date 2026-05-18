package main

import "fmt"

func main() {
	resultado := 0.0
	numero := 1.0
	denominador := 15.0
	sinal := 1.0
	for i := 0; i < 15; i++ {
		resultado += sinal * numero / (denominador * denominador)
		numero *= 2
		denominador--
		sinal *= -1
	}
	fmt.Printf("%.10f\n", resultado)
}
