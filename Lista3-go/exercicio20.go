package main

import "fmt"

func main() {
	for i := 1; i < 10; i++ {
		for j := 0; j < i; j++ {
			fmt.Printf("(%d,%d) ", i, j)
		}
		fmt.Println()
	}
}
