package main

import "fmt"

func main() {
	var preco float64
	var opcao string
	total := 0.0

	fmt.Print("Digite o preço inicial do carro: R$ ")
	fmt.Scan(&preco)

	fmt.Println("\nEscolha os opcionais (digite 'sim' para sim ou 'não' para não):")

	fmt.Print("Ar condicionado (R$ 1750,00)? ")
	fmt.Scan(&opcao)
	if opcao == "sim" {
		total += 1750.00
	}

	fmt.Print("Pintura metálica (R$ 800,00)? ")
	fmt.Scan(&opcao)
	if opcao == "sim" {
		total += 800.00
	}

	fmt.Print("Vidro elétrico (R$ 1200,00)? ")
	fmt.Scan(&opcao)
	if opcao == "sim" {
		total += 1200.00
	}

	fmt.Print("Direção hidráulica (R$ 2000,00)? ")
	fmt.Scan(&opcao)
	if opcao == "sim" {
		total += 2000.00
	}

	precoFinal := preco + total

	fmt.Printf("Preço final: R$ %.2f\n", precoFinal)
}