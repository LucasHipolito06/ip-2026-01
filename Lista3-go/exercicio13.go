package main

import "fmt"

func main() {
	resultado := 0.0
	for i := 1; i <= 50; i++ {
		numerador := float64(2*i - 1)
		resultado += numerador / float64(i)
	}
	fmt.Printf("%.10f\n", resultado)
}
