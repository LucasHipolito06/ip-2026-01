package main

import "fmt"

func main() {
	var janela [24]int
	var corredor [24]int

	for {
		fmt.Println("\n===== VENDA DE PASSAGENS =====")
		fmt.Println("1. Comprar poltrona na janela")
		fmt.Println("2. Comprar poltrona no corredor")
		fmt.Println("3. Sair")
		fmt.Print("Opção: ")

		var opcao int
		fmt.Scan(&opcao)

		switch opcao {
		case 1:
			disponiveis := listarDisponiveis(janela[:], "Janela")
			if disponiveis == 0 {
				fmt.Println("Não há poltronas livres na janela!")
				if todasOcupadas(corredor[:]) {
					fmt.Println("O ônibus está completamente cheio!")
				}
			} else {
				var pol int
				fmt.Print("Escolha o número da poltrona: ")
				fmt.Scan(&pol)
				if pol >= 0 && pol < 24 && janela[pol] == 0 {
					janela[pol] = 1
					fmt.Printf("Poltrona %d da janela reservada com sucesso!\n", pol)
				} else {
					fmt.Println("Poltrona inválida ou já ocupada!")
				}
			}

		case 2:
			disponiveis := listarDisponiveis(corredor[:], "Corredor")
			if disponiveis == 0 {
				fmt.Println("Não há poltronas livres no corredor!")
				if todasOcupadas(janela[:]) {
					fmt.Println("O ônibus está completamente cheio!")
				}
			} else {
				var pol int
				fmt.Print("Escolha o número da poltrona: ")
				fmt.Scan(&pol)
				if pol >= 0 && pol < 24 && corredor[pol] == 0 {
					corredor[pol] = 1
					fmt.Printf("Poltrona %d do corredor reservada com sucesso!\n", pol)
				} else {
					fmt.Println("Poltrona inválida ou já ocupada!")
				}
			}

		case 3:
			fmt.Println("Programa encerrado.")
			return

		default:
			fmt.Println("Opção inválida!")
		}
	}
}

func listarDisponiveis(vetor []int, tipo string) int {
	count := 0
	fmt.Printf("Poltronas disponíveis (%s): ", tipo)
	for i, v := range vetor {
		if v == 0 {
			fmt.Printf("%d ", i)
			count++
		}
	}
	fmt.Println()
	return count
}

func todasOcupadas(vetor []int) bool {
	for _, v := range vetor {
		if v == 0 {
			return false
		}
	}
	return true
}
