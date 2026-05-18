package main

import "fmt"

func main() {
	var fib [50]int64

	fib[0] = 1
	fib[1] = 1

	for i := 2; i < 50; i++ {
		fib[i] = fib[i-1] + fib[i-2]
	}

	fmt.Println("Os primeiros 50 termos da série de Fibonacci:")
	for i := 0; i < 50; i++ {
		fmt.Printf("Termo %d: %d\n", i+1, fib[i])
	}
}
