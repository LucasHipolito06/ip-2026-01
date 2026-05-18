package main

import (
	"fmt"
	"math"
)

func senoAproximado(a float64) float64 {
	return a - math.Pow(a, 3)/6 + math.Pow(a, 5)/120 - math.Pow(a, 7)/5040
}

func main() {
	for i := 0; i <= 63; i++ {
		a := float64(i) / 10.0
		fmt.Printf("%.1f %.10f\n", a, senoAproximado(a))
	}
}
