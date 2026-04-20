package main

import "fmt"

func ehPrimo(n int) bool {
	if n < 2 {
		return false
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	var n1, n2 int
	fmt.Scan(&n1, &n2)
	if n1 > n2 {
		n1, n2 = n2, n1
	}
	for numero := n1; numero <= n2; numero++ {
		if ehPrimo(numero) {
			fmt.Println(numero)
		}
	}
}
