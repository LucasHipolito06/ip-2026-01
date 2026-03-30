package main

import "fmt"

func main() {

	var inicio, razao, nTermos int

	fmt.Print("Valor inicial da PA: ")
	fmt.Scan(&inicio)

	fmt.Print("Razão da PA: ")
	fmt.Scan(&razao)

	fmt.Print("Quantidade de termos: ")
	fmt.Scan(&nTermos)

	soma := 0
	termo := inicio

	for i := 1; i <= nTermos; i++ {
		soma += termo
		termo += razao
	}

	fmt.Println("Soma da PA:", soma)
}
