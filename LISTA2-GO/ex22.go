package main

import "fmt"

func main() {
	var matricula int
	var horasExtras float64
	var descontoINSS, descontoIR float64

	const salarioMinimo = 788.00
	const valorHoraExtra = 10.00

	fmt.Print("Digite a matrícula do funcionário: ")
	fmt.Scan(&matricula)
	fmt.Print("Digite a quantidade de horas-extras: ")
	fmt.Scan(&horasExtras)

	salarioHoraExtra := horasExtras * valorHoraExtra
	salarioBruto := 3*salarioMinimo + salarioHoraExtra

	if salarioBruto > 1500.00 {
		descontoINSS = salarioBruto * 0.12
	}

	if salarioBruto > 2000.00 {
		descontoIR = salarioBruto * 0.20
	}

	salarioLiquido := salarioBruto - descontoINSS - descontoIR

	fmt.Printf("Matrícula do funcionário: %d\n", matricula)
	fmt.Printf("Salário bruto: R$ %.2f\n", salarioBruto)
	fmt.Printf("Salário líquido: R$ %.2f\n", salarioLiquido)
}