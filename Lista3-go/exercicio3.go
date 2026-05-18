package main

import "fmt"

func main() {
	var salarioCarlos float64
	fmt.Scan(&salarioCarlos)
	salarioJoao := salarioCarlos / 3.0
	meses := 0
	for salarioJoao < salarioCarlos {
		salarioCarlos *= 1.02
		salarioJoao *= 1.05
		meses++
	}
	fmt.Println(meses)
}
