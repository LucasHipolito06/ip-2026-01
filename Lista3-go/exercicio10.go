package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	if n <= 0 {
		return
	}
	a, b := 0, 1
	for i := 1; i <= n; i++ {
		if i == 1 {
			fmt.Print(a)
		} else if i == 2 {
			fmt.Print(" ", b)
		} else {
			c := a + b
			fmt.Print(" ", c)
			a, b = b, c
		}
	}
	fmt.Println()
}
