package main

import "fmt"

func main() {

	var numeros [10]int

	for i := 0; i < 10; i++ {
		fmt.Printf("Digite o %d° valor: ", i+1)
		fmt.Scan(&numeros[i])
	}

	fmt.Println("Valores na ordem inversa:")

	for i := 9; i >= 0; i-- {
		fmt.Println(numeros[i])
	}
}
