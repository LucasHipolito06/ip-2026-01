package main

import (
	"fmt"
	"math"
)

func main() {

	var raio float64
	var altura float64

	fmt.Print("Raio da lata em metros: ")
	fmt.Scan(&raio)

	fmt.Print("Altura da lata em metros: ")
	fmt.Scan(&altura)

	area_circulo := math.Pi * raio * raio
	area_lateral := 2 * math.Pi * raio * altura
	area_total := 2*area_circulo + area_lateral
	custo := area_total * 100

	fmt.Printf("Custo: R$ %.2f\n", custo)
}
