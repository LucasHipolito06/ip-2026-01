package main

import "fmt"

func main() {
	var dia, mes, ano int
	var nomeMes string

	fmt.Print("Digite o dia: ")
	fmt.Scan(&dia)
	fmt.Print("Digite o mês (número): ")
	fmt.Scan(&mes)
	fmt.Print("Digite o ano: ")
	fmt.Scan(&ano)

	if mes == 1 {
		nomeMes = "Janeiro"
	} else if mes == 2 {
		nomeMes = "Fevereiro"
	} else if mes == 3 {
		nomeMes = "Março"
	} else if mes == 4 {
		nomeMes = "Abril"
	} else if mes == 5 {
		nomeMes = "Maio"
	} else if mes == 6 {
		nomeMes = "Junho"
	} else if mes == 7 {
		nomeMes = "Julho"
	} else if mes == 8 {
		nomeMes = "Agosto"
	} else if mes == 9 {
		nomeMes = "Setembro"
	} else if mes == 10 {
		nomeMes = "Outubro"
	} else if mes == 11 {
		nomeMes = "Novembro"
	} else if mes == 12 {
		nomeMes = "Dezembro"
	} else {
		nomeMes = "inválido"
	}

	if nomeMes == "inválido" {
		fmt.Println("Erro: mês inválido!")
	} else {
		fmt.Printf("\n%d de %s de %d\n", dia, nomeMes, ano)
	}
}