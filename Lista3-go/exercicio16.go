package main

import "fmt"

func main() {
	var n, a1, a2 int
	fmt.Scan(&n, &a1, &a2)
	if n <= 0 {
		return
	}
	if n >= 1 {
		fmt.Print(a1)
	}
	if n >= 2 {
		fmt.Print(" ", a2)
	}
	for i := 3; i <= n; i++ {
		var atual int
		if i%2 == 1 {
			atual = a2 + a1
		} else {
			atual = a2 - a1
		}
		fmt.Print(" ", atual)
		a1, a2 = a2, atual
	}
	fmt.Println()
}
