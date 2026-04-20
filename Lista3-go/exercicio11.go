package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	if n < 0 {
		fmt.Println("valor invalido")
		return
	}
	resultado := 1
	for i := 2; i <= n; i++ {
		resultado *= i
	}
	fmt.Println(resultado)
}
