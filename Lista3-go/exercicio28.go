package main

import (
	"fmt"
	"math"
)

func main() {
	soma := 0.0
	sinal := 1.0
	denominador := 1.0
	for i := 0; i < 51; i++ {
		soma += sinal / math.Pow(denominador, 3)
		sinal *= -1
		denominador += 2
	}
	pi := math.Cbrt(32 * soma)
	fmt.Printf("%.10f\n", pi)
}
