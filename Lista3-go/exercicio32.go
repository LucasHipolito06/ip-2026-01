package main

import "fmt"

func main() {
	var n1, n2 int
	fmt.Scan(&n1, &n2)
	resultado := 0
	positivo := n2 >= 0
	if !positivo {
		n2 = -n2
	}
	for i := 0; i < n2; i++ {
		if positivo {
			resultado += n1
		} else {
			resultado -= n1
		}
	}
	fmt.Println(resultado)
}
