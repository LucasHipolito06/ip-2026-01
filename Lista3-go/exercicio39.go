package main

import "fmt"

func main() {
	var id, idMaisGordo, idMaisMagro int
	var peso, maiorPeso, menorPeso float64
	for i := 0; i < 90; i++ {
		fmt.Scan(&id, &peso)
		if i == 0 || peso > maiorPeso {
			maiorPeso = peso
			idMaisGordo = id
		}
		if i == 0 || peso < menorPeso {
			menorPeso = peso
			idMaisMagro = id
		}
	}
	fmt.Println(idMaisGordo, maiorPeso)
	fmt.Println(idMaisMagro, menorPeso)
}
