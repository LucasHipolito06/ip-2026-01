package main

import "fmt"

func main() {
	var n1, n2 int
	fmt.Scan(&n1, &n2)
	maior := n1
	if n2 > maior {
		maior = n2
	}
	multiplo := maior
	for {
		if multiplo%n1 == 0 && multiplo%n2 == 0 {
			fmt.Println(multiplo)
			return
		}
		multiplo += maior
	}
}
