package main

import "fmt"

func main() {
	var preco float64
	var diaSem, categoria int

	fmt.Print("Digite o preço normal do DVD: R$ ")
	fmt.Scan(&preco)

	fmt.Println("1 - Segunda")
	fmt.Println("2 - Terça")
	fmt.Println("3 - Quarta")
	fmt.Println("4 - Quinta")
	fmt.Println("5 - Sexta")
	fmt.Println("6 - Sábado")
	fmt.Println("7 - Domingo")
	fmt.Print("Digite o dia da semana: ")
	fmt.Scan(&diaSem)

	fmt.Println("1 - Comum")
	fmt.Println("2 - Lançamento")
	fmt.Print("Digite a categoria: ")
	fmt.Scan(&categoria)

	if diaSem < 1 || diaSem > 7 {
		fmt.Println("Erro: dia da semana inválido!")
	} else if categoria < 1 || categoria > 2 {
		fmt.Println("Erro: categoria inválida!")
	} else {
		precoFinal := preco

		if diaSem == 1 || diaSem == 2 || diaSem == 4 {
			precoFinal = precoFinal * 0.60 
		}

		if categoria == 2 {
			precoFinal = precoFinal * 1.15
		}

		fmt.Printf("\nPreço normal:  R$ %.2f\n", preco)
		fmt.Printf("Preço final:   R$ %.2f\n", precoFinal)
	}
}