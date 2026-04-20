package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	if n == 0 {
		fmt.Println("0")
		return
	}
	digitos := "0123456789ABCDEF"
	resultado := make([]byte, 0)
	for n > 0 {
		resultado = append(resultado, digitos[n%16])
		n /= 16
	}
	for i := len(resultado) - 1; i >= 0; i-- {
		fmt.Printf("%c", resultado[i])
	}
	fmt.Println()
}
