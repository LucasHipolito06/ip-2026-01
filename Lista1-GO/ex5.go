package main

import "fmt"

func main() {
	var conta int
	var consumo float64
	var tipo string

	fmt.Print("Número da conta: ")
	fmt.Scan(&conta)
	fmt.Print("Consumo: ")
	fmt.Scan(&consumo)
	fmt.Print("Tipo (R/C/I): ")
	fmt.Scan(&tipo)

	var valor float64
	switch tipo {
	case "R":
		valor = 5 + consumo*0.05
	case "C":
		valor = 500 + consumo*0.25
	case "I":
		valor = 800 + consumo*0.04
	default:
		fmt.Println("Tipo inválido")
		return
	}

	fmt.Printf("Conta: %d | Valor: R$ %.2f\n", conta, valor)
}
