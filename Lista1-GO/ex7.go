package main

import "fmt"

func main() {

	var fahrenheit float64
	var polegadas float64

	fmt.Print("Temperatura em Fahrenheit: ")
	fmt.Scan(&fahrenheit)

	fmt.Print("Quantidade de chuva em polegadas: ")
	fmt.Scan(&polegadas)

	celsius := (5*fahrenheit - 160) / 9
	chuva_mm := polegadas * 25.4

	fmt.Printf("Temperatura em Celsius: %.1f\n", celsius)
	fmt.Printf("Quantidade de chuva: %.1f mm\n", chuva_mm)
}
