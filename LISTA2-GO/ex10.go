package main

import "fmt"

func main() {
	var destino, retorno int
	var preco float64

	fmt.Println("1 - Região Norte")
	fmt.Println("2 - Região Nordeste")
	fmt.Println("3 - Região Centro-Oeste")
	fmt.Println("4 - Região Sul")
	fmt.Print("\nDigite o destino: ")
	fmt.Scan(&destino)

	if destino < 1 || destino > 4 {
		fmt.Println("Erro: destino inválido!")
	} else {
		fmt.Print("Inclui retorno? (1 - Sim / 2 - Não): ")
		fmt.Scan(&retorno)

		if retorno < 1 || retorno > 2 {
			fmt.Println("Erro: opção de retorno inválida!")
		} else {
			

			if destino == 1 && retorno == 1 {
				preco = 900.00
			} else if destino == 1 && retorno == 2 {
				preco = 500.00
			} else if destino == 2 && retorno == 1 {
				preco = 650.00
			} else if destino == 2 && retorno == 2 {
				preco = 350.00
			} else if destino == 3 && retorno == 1 {
				preco = 600.00
			} else if destino == 3 && retorno == 2 {
				preco = 350.00
			} else if destino == 4 && retorno == 1 {
				preco = 550.00
			} else if destino == 4 && retorno == 2 {
				preco = 300.00
			}

			fmt.Printf("\nO preço da passagem é: R$ %.2f\n", preco)
		}
	}
}