package main

import "fmt"

func main() {
	var idade int
	var altura, peso float64
	var continuar int
	var qtdMais50 int
	var somaAlturas float64
	var qtdEntre10e20 int
	var qtdPesoMenor40 int
	var total int

	for {
		fmt.Scan(&idade, &altura, &peso)
		total++
		if idade > 50 {
			qtdMais50++
		}
		if idade >= 10 && idade <= 20 {
			somaAlturas += altura
			qtdEntre10e20++
		}
		if peso < 40 {
			qtdPesoMenor40++
		}
		fmt.Scan(&continuar)
		if continuar != 1 {
			break
		}
	}

	mediaAltura := 0.0
	if qtdEntre10e20 > 0 {
		mediaAltura = somaAlturas / float64(qtdEntre10e20)
	}
	percentual := 0.0
	if total > 0 {
		percentual = float64(qtdPesoMenor40) * 100 / float64(total)
	}

	fmt.Println(qtdMais50)
	fmt.Printf("%.2f\n", mediaAltura)
	fmt.Printf("%.2f\n", percentual)
}
