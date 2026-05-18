package main

import "fmt"

func main() {
	var numero int

	fmt.Print("Digite um número inteiro de 3 casas: ")
	fmt.Scan(&numero)

	if numero < 100 || numero > 999 {
		fmt.Println("Erro: o número deve ter exatamente 3 casas!")
	} else {
		dezena := (numero / 10) % 10
		fmt.Println("O algarismo das dezenas é:", dezena)
	}
}

