package main

import "fmt"

func main() {

	var numeros [5]int
	soma := 0

	for i := 0; i < 5; i++ {
		fmt.Printf("Digite o %v° valor: ", i+1)
		fmt.Scan(&numeros[i])
	}

	for _, n := range numeros {
		soma += n
	}

	fmt.Println("Soma total dos números é:", soma)
}
