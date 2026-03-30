package main

import "fmt"

func main() {

	var altura float64
	var aresta float64

	fmt.Print("Altura da pirâmide em metros: ")
	fmt.Scan(&altura)

	fmt.Print("Aresta da pirâmide em metros: ")
	fmt.Scan(&aresta)

	// raiz de 3 aproximada como 1.73
	areaBase := 3 * aresta * aresta * 1.73 / 2
	volume := areaBase * altura / 3

	fmt.Printf("Volume da pirâmide: %.2f m³\n", volume)
}
