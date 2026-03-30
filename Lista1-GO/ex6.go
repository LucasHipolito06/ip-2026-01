package main

import "fmt"

func main() {

	var temperaturas []float64
	var fahrenheit float64

	for {
		fmt.Print("Digite uma temperatura em Fahrenheit (-1 para sair): ")
		fmt.Scan(&fahrenheit)

		if fahrenheit == -1 {
			break
		}

		temperaturas = append(temperaturas, fahrenheit)
	}

	for _, f := range temperaturas {
		celsius := 5 * (f - 32) / 9
		fmt.Printf("%.1f Fahrenheit equivale a %.1f Celsius\n", f, celsius)
	}
}
