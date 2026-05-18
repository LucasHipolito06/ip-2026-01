package main

import "fmt"

func main() {
	var total, popular, geral, arqui, cadeira int

	fmt.Print("Público total do jogo: ")
	fmt.Scan(&total)

	fmt.Print("Porcentagem na popular: ")
	fmt.Scan(&popular)

	fmt.Print("Porcentagem na geral: ")
	fmt.Scan(&geral)

	fmt.Print("Porcentagem na arquibancada: ")
	fmt.Scan(&arqui)

	fmt.Print("Porcentagem na cadeira: ")
	fmt.Scan(&cadeira)

	renda_total := (popular*total*1 + geral*total*5 + arqui*total*10 + cadeira*total*20) / 100

	fmt.Println("A renda total do jogo é:", renda_total)
}
