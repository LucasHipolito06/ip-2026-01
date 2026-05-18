package main

import "fmt"

func main() {

	var horas int

	fmt.Print("Quantas horas a charrete foi alugada? ")
	fmt.Scan(&horas)

	blocos := horas / 3
	resto := horas % 3
	valor := blocos*10 + resto*5

	fmt.Println("O valor total a pagar é:", valor)
}
