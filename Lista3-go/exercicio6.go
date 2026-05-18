package main

import "fmt"

func eTriangular(n int) bool {
	for k := 1; ; k++ {
		produto := k * (k + 1) * (k + 2)
		if produto == n {
			return true
		}
		if produto > n {
			return false
		}
	}
}

func main() {
	var numero int
	fmt.Scan(&numero)
	if eTriangular(numero) {
		fmt.Println("e triangular")
	} else {
		fmt.Println("nao e triangular")
	}
}
