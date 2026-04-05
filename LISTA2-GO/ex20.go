package main

import "fmt"

func main() {
	var preco, total float64
	var codigo, parcelas int

	fmt.Print("Digite o preço do produto: R$ ")
	fmt.Scan(&preco)

	fmt.Println("1 - À vista (dinheiro/cheque) → 10% de desconto")
	fmt.Println("2 - À vista (cartão de crédito) → 5% de desconto")
	fmt.Println("3 - Em 2x → sem juros")
	fmt.Println("4 - Em 3x → 10% de juros")
	fmt.Print("Digite o código da condição: ")
	fmt.Scan(&codigo)

	if codigo == 1 {
			total = preco * 0.90
			parcelas = 1
	} else if codigo == 2 {
			total = preco * 0.95
			parcelas = 1
	} else if codigo == 3 {
			total = preco
			parcelas = 2
	} else if codigo == 4 {
			total = preco * 1.10
			parcelas = 3
	}


		fmt.Printf("Preço do produto: R$ %.2f\n", preco)
		fmt.Printf("Total a pagar: R$ %.2f\n", total)

		if parcelas > 1 {
			fmt.Printf("Parcelas: %dx de R$ %.2f\n", parcelas, total/float64(parcelas))
		} else {
			fmt.Println("Pagamento: À vista")
		}
	}
