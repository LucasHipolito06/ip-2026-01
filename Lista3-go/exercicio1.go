package main

import "fmt"

func main() {
	var base, expoente int
	fmt.Scan(&base, &expoente)
	resultado := 1
	for i := 0; i < expoente; i++ {
		resultado *= base
	}
	fmt.Println(resultado)
}
