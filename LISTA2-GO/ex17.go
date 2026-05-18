package main

import "fmt"

func main() {
	var conta, tipo int
	var consumo, total float64

	fmt.Print("Digite o número da conta do cliente: ")
	fmt.Scan(&conta)

	fmt.Println("1 - Residencial")
	fmt.Println("2 - Comercial")
	fmt.Println("3 - Industrial")
	fmt.Print("Digite o tipo: ")
	fmt.Scan(&tipo)

	fmt.Print("Digite o consumo em m³: ")
	fmt.Scan(&consumo)

		if tipo == 1 {
			total = 5.00 + (consumo * 0.05)
		} else if tipo == 2 {
			if consumo <= 80 {
				total = 500.00
			} else {
				total = 500.00 + ((consumo - 80) * 0.25)
			}
		} else if tipo == 3 {
			if consumo <= 100 {
				total = 800.00
			} else {
				total = 800.00 + ((consumo - 100) * 0.04)
			}
		}

		fmt.Printf("Conta: %d\n", conta)
		fmt.Printf("Consumo: %.2f m³\n", consumo)
		fmt.Printf("Total: R$ %.2f\n", total)
	}
