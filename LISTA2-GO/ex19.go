package main

import (
	"fmt"
	"math"
)

func main() {
	var opcao int
	var raio, altura float64

	const pi = math.Pi

	fmt.Println("1 - Cone")
	fmt.Println("2 - Cilindro")
	fmt.Println("3 - Esfera")
	fmt.Print("Digite a opção: ")
	fmt.Scan(&opcao)

	 if opcao == 1 {
		fmt.Print("Digite o raio do cone: ")
		fmt.Scan(&raio)
		fmt.Print("Digite a altura do cone: ")
		fmt.Scan(&altura)

		volume := (pi * raio * raio * altura) / 3
		area := pi * raio * math.Sqrt(raio*raio + altura*altura)

		fmt.Printf("Volume:          %.2f\n", volume)
		fmt.Printf("Área superfície: %.2f\n", area)

	} else if opcao == 2 {
		fmt.Print("Digite o raio do cilindro: ")
		fmt.Scan(&raio)
		fmt.Print("Digite a altura do cilindro: ")
		fmt.Scan(&altura)

		volume := pi * raio * raio * altura
		area := 2 * pi * raio * altura

		fmt.Printf("Volume:          %.2f\n", volume)
		fmt.Printf("Área superfície: %.2f\n", area)

	} else if opcao == 3 {
		fmt.Print("Digite o raio da esfera: ")
		fmt.Scan(&raio)

		volume := (4.0 / 3.0) * pi * raio * raio * raio
		area := 4 * pi * raio * raio

		fmt.Printf("Volume: %.2f\n", volume)
		fmt.Printf("Área superfície: %.2f\n", area)
	}
}