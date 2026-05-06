package main

import "fmt"

func main() {
	var codigos [10]int
	var saldos [10]float64

	fmt.Println("Cadastro de 10 contas bancárias:")
	for i := 0; i < 10; i++ {
		fmt.Printf("\nConta %d:\n", i+1)
		fmt.Print("  Código: ")
		fmt.Scan(&codigos[i])
		fmt.Print("  Saldo inicial: ")
		fmt.Scan(&saldos[i])
	}

	for {
		fmt.Println("\n===== MENU =====")
		fmt.Println("1. Efetuar depósito")
		fmt.Println("2. Efetuar saque")
		fmt.Println("3. Consultar o ativo bancário")
		fmt.Println("4. Finalizar o programa")
		fmt.Print("Opção: ")

		var opcao int
		fmt.Scan(&opcao)

		switch opcao {
		case 1:
			var cod int
			var valor float64
			fmt.Print("Código da conta: ")
			fmt.Scan(&cod)
			idx := buscarConta(codigos[:], cod)
			if idx == -1 {
				fmt.Println("Conta não encontrada.")
			} else {
				fmt.Print("Valor do depósito: ")
				fmt.Scan(&valor)
				saldos[idx] += valor
				fmt.Printf("Depósito realizado. Novo saldo: %.2f\n", saldos[idx])
			}

		case 2:
			var cod int
			var valor float64
			fmt.Print("Código da conta: ")
			fmt.Scan(&cod)
			idx := buscarConta(codigos[:], cod)
			if idx == -1 {
				fmt.Println("Conta não encontrada.")
			} else {
				fmt.Print("Valor do saque: ")
				fmt.Scan(&valor)
				if valor > saldos[idx] {
					fmt.Println("Saldo insuficiente.")
				} else {
					saldos[idx] -= valor
					fmt.Printf("Saque realizado. Novo saldo: %.2f\n", saldos[idx])
				}
			}

		case 3:
			ativo := 0.0
			for i := 0; i < 10; i++ {
				ativo += saldos[i]
			}
			fmt.Printf("Ativo bancário total: %.2f\n", ativo)

		case 4:
			fmt.Println("Programa finalizado.")
			return

		default:
			fmt.Println("Opção inválida!")
		}
	}
}

func buscarConta(codigos []int, cod int) int {
	for i, c := range codigos {
		if c == cod {
			return i
		}
	}
	return -1
}
