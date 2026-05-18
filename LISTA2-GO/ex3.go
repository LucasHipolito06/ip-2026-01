package main

import "fmt"

func main () {
	var n1, n2 int
	soma := 0

	fmt.Printf("Digite dois números: ")
	fmt.Scan(&n1, &n2)

	soma = (n1 + n2)
	if soma <= 20 {
		fmt.Println("A soma é:", soma - 5)
	} else {
		fmt.Println("A soma é: ", soma + 8)
	}
}