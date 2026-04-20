package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	resultado := 0
	mult := 1
	for n > 0 {
		resultado += (n % 10) * mult
		mult *= 8
		n /= 10
	}
	fmt.Println(resultado)
}
