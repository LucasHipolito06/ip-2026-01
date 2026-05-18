package main

import "fmt"

func main() {

	var hora, minuto, segundo int

	fmt.Print("Em que hora estou? ")
	fmt.Scan(&hora)

	fmt.Print("Em que minuto estou? ")
	fmt.Scan(&minuto)

	fmt.Print("Em que segundo estou? ")
	fmt.Scan(&segundo)

	total := hora*3600 + minuto*60 + segundo

	fmt.Println("O tempo em segundos é:", total)
}
