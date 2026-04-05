package main

import "fmt"

func main () {
	var n1, n2 int

	fmt.Printf("Digite dois números: ")
	fmt.Scan(&n1, &n2)

	if n1 % n2 == 0 {
		fmt.Println("O número é divisível")
	} else {
		fmt.Println("O número não é divisível")
	}
}