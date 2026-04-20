package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	resultado := 0.0
	numero := 1000.0
	sinal := 1.0
	for i := 1; i <= n; i++ {
		resultado += sinal * numero / float64(i)
		numero -= 3
		sinal *= -1
	}
	fmt.Printf("%.10f\n", resultado)
}
