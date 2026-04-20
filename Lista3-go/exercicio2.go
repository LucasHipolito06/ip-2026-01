package main

import "fmt"

func main() {
	soma := 0
	quantidade := 0
	for numero := 50; numero <= 70; numero += 2 {
		soma += numero
		quantidade++
	}
	fmt.Println(soma)
	fmt.Printf("%.2f\n", float64(soma)/float64(quantidade))
}
