package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	var aprovados, exames, reprovados int
	var somaClasse float64

	for i := 0; i < n; i++ {
		var nota1, nota2 float64
		fmt.Scan(&nota1, &nota2)
		media := (nota1 + nota2) / 2
		somaClasse += media
		fmt.Printf("%.2f\n", media)
		if media <= 3 {
			fmt.Println("Reprovado")
			reprovados++
		} else if media < 7 {
			fmt.Println("Exame")
			exames++
		} else {
			fmt.Println("Aprovado")
			aprovados++
		}
	}

	if n > 0 {
		fmt.Println(aprovados)
		fmt.Println(exames)
		fmt.Println(reprovados)
		fmt.Printf("%.2f\n", somaClasse/float64(n))
	}
}
