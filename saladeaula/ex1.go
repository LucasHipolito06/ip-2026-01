package main

import "fmt"

func main() {

	var nota1, nota2, nota3 float64

	fmt.Println("Digite a 1° nota: ")
	fmt.Scan(&nota1)
	fmt.Println("Digite a 2° nota: ")
	fmt.Scan(&nota2)
	fmt.Println("Digite a 3° nota: ")
	fmt.Scan(&nota3)

	notas := []float64{nota1, nota2, nota3}
	soma := 0.0

	for _, n := range notas {
		soma += n
	}

	media := soma / 3
	fmt.Printf("A média das três notas é %.2f\n", media)

	for i, n := range notas {
		if n >= media {
			fmt.Printf("A %d° nota está acima da média\n", i+1)
		} else {
			fmt.Printf("A %d° nota não está acima da média\n", i+1)
		}
	}
}
