package main 

import "fmt"

func main () {
	var v_compra, v_venda float64

	fmt.Printf("Digite o valor da compra: ")
	fmt.Scan(&v_compra)

	if v_compra < 10 {
		v_venda = v_compra * 1.7 
	} else if v_compra >= 10 && v_compra < 30  {
		v_venda = v_compra * 1.5
	} else if v_compra >= 30 && v_compra < 50 {
		v_venda = v_compra * 1.4
	} else if v_compra >= 50 {
		v_venda = v_compra * 1.3
	}

	fmt.Println("O valor da venda deve ser:", v_venda)
}