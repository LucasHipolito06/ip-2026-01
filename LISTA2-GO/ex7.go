package main

import "fmt"

func main() {
	var a, b, c int
	var menor, maior, inter int

	fmt.Print("Digite o valor de A: ")
	fmt.Scan(&a)
	fmt.Print("Digite o valor de B: ")
	fmt.Scan(&b)
	fmt.Print("Digite o valor de C: ")
	fmt.Scan(&c)

	if a < b && a < c {
		menor = a
	} else if b < a && b < c {
		menor = b
	} else {
		menor = c
	}

	if a > b && a > c {
		maior = a
	} else if b > a && b > c {
		maior = b
	} else {
		maior = c
	}

	if a != menor && a != maior {
		inter = a
	} else if b != menor && b != maior {
		inter = b
	} else {
		inter = c
	}

	fmt.Println("\nOrdem crescente:")
	fmt.Println("Menor:", menor)
	fmt.Println("Inter:", inter)
	fmt.Println("Maior:", maior)
}