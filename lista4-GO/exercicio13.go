package main

import "fmt"

func main() {
	var numEmp [100]int
	var meses [100]int
	n := 0

	fmt.Println("Digite o número do empregado e os meses de trabalho (0 0 para parar):")
	for n < 100 {
		var emp, mes int
		fmt.Printf("Empregado %d (número meses): ", n+1)
		fmt.Scan(&emp, &mes)
		if emp == 0 && mes == 0 {
			break
		}
		numEmp[n] = emp
		meses[n] = mes
		n++
	}

	// Encontrar os 3 com menor número de meses (mais recentes)
	for k := 0; k < 3 && k < n; k++ {
		menorIdx := k
		for j := k + 1; j < n; j++ {
			if meses[j] < meses[menorIdx] {
				menorIdx = j
			}
		}
		numEmp[k], numEmp[menorIdx] = numEmp[menorIdx], numEmp[k]
		meses[k], meses[menorIdx] = meses[menorIdx], meses[k]
	}

	fmt.Println("\nOs 3 empregados mais recentes:")
	limite := 3
	if n < 3 {
		limite = n
	}
	for i := 0; i < limite; i++ {
		fmt.Printf("Empregado %d - %d meses de trabalho\n", numEmp[i], meses[i])
	}
}
