package main

import "fmt"

func eQuadradoPerfeito(n int) bool {
	for i := 1; i*i <= n; i++ {
		if i*i == n {
			return true
		}
	}
	return false
}

func main() {
	for {
		var numero int
		fmt.Scan(&numero)
		if numero <= 0 {
			break
		}
		if eQuadradoPerfeito(numero) {
			fmt.Println("quadrado perfeito")
		} else {
			fmt.Println("nao e quadrado perfeito")
		}
	}
}
