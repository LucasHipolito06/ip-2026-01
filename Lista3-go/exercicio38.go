package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

func onlyDigits(s string) []int {
	digitos := make([]int, 0, len(s))
	for _, r := range s {
		if unicode.IsDigit(r) {
			digitos = append(digitos, int(r-'0'))
		}
	}
	return digitos
}

func calculaDigito(base []int, pesoInicial int) int {
	soma := 0
	peso := pesoInicial
	for _, d := range base {
		soma += d * peso
		peso--
	}
	resto := soma % 11
	if resto < 2 {
		return 0
	}
	return 11 - resto
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	var cpf string
	fmt.Fscan(reader, &cpf)
	digitos := onlyDigits(cpf)
	if len(digitos) != 11 {
		fmt.Println("CPF invalido")
		return
	}
	primeiro := calculaDigito(digitos[:9], 10)
	segundo := calculaDigito(append(digitos[:9], primeiro), 11)
	if primeiro == digitos[9] && segundo == digitos[10] {
		fmt.Println("CPF valido")
	} else {
		fmt.Println("CPF invalido")
	}
}
