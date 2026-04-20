package main

import "fmt"

func main() {
	var b, n int
	fmt.Scan(&b, &n)
	resultado := 1
	for i := 0; i < n; i++ {
		resultado *= b
	}
	fmt.Println(resultado)
}
