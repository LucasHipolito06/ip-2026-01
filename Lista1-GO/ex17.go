package main

import "fmt"

func main() {

	var x, y int

	fmt.Print("Qual o valor de x? ")
	fmt.Scan(&x)

	fmt.Print("Qual o valor de y? ")
	fmt.Scan(&y)

	if x%2 == 0 {
		for i := 0; i < y; i++ {
			fmt.Printf("%d ", x+i*2)
		}
		fmt.Println()
	} else {
		fmt.Println("O primeiro número não é par")
	}
}
