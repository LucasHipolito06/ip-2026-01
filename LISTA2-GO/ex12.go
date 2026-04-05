package main 

import "fmt"

func main () {
	var idade int

	fmt.Printf("Digite a idade da pessoa: ")
	fmt.Scan(&idade)

	if idade >= 0 && idade <= 2 {
		fmt.Println("A pessoa é um recém nascido")
	} else if idade > 2 && idade <= 11 {
		fmt.Println("A pessoa é uma criança")
	} else if idade > 11 && idade <= 19 {
		fmt.Println("A pessoa é um adolescente")
	} else if idade > 19 && idade <= 55 {
		fmt.Println("A pessoa é um adulto")
	} else if idade > 55 {
		fmt.Println("A pessoa é um idoso")
	}
}