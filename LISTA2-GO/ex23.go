package main

import "fmt"

func main () {
	var idade int

	fmt.Printf("Digite a idade do eleitor: ")
	fmt.Scan(&idade)

	if idade <= 16 {
		fmt.Println("A classe eleitoral é Não-eleitor")
	} else if idade >= 18 && idade <= 65 {
		fmt.Println("A classe eleitoral é Eleitor obrigatório")
	} else if idade > 65 {
		fmt.Println("A classe eleitoral é Eleitor facultativo")

	}
}