package main

import "fmt"

func main() {
	var salario, kw float64

	fmt.Print("Qual o salário mínimo? R$ ")
	fmt.Scan(&salario)
	fmt.Print("Quantos kw foram gastos neste mês? ")
	fmt.Scan(&kw)

	custoPorKw := salario * 0.007
	custoConsumo := custoPorKw * kw
	custoDesconto := custoConsumo * 0.9

	fmt.Printf("Custo por kw:       R$ %.2f\n", custoPorKw)
	fmt.Printf("Custo do consumo:   R$ %.2f\n", custoConsumo)
	fmt.Printf("Custo com desconto: R$ %.2f\n", custoDesconto)
}
