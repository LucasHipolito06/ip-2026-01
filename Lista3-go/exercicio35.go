package main

import (
	"fmt"
	"strconv"
)

func main() {
	var n int
	fmt.Scan(&n)
	if n == 0 {
		fmt.Println("0")
		return
	}
	partes := make([]string, 0)
	for n > 0 {
		partes = append(partes, strconv.Itoa(n%2))
		n /= 2
	}
	for i := len(partes) - 1; i >= 0; i-- {
		fmt.Print(partes[i])
	}
	fmt.Println()
}
