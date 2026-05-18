package main

import (
	"fmt"
	"math/big"
)

func main() {
	total := big.NewInt(0)
	grao := big.NewInt(1)
	for i := 0; i < 64; i++ {
		total.Add(total, grao)
		grao.Mul(grao, big.NewInt(2))
	}
	fmt.Println(total.String())
}
