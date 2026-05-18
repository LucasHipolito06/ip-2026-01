package main

import (
	"fmt"
	"math"
)

func main() {
	for i := 0; i <= 40; i++ {
		raio := float64(i) * 0.5
		volume := 4.0 / 3.0 * math.Pi * math.Pow(raio, 3)
		fmt.Printf("%.1f %.10f\n", raio, volume)
	}
}
