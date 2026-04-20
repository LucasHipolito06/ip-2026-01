package main

import "fmt"

func main() {
	resultado := 0.0
	for i := 1; i <= 37; i++ {
		numerador := float64(38-i) * float64(39-i)
		resultado += numerador / float64(i)
	}
	fmt.Printf("%.10f\n", resultado)
}
