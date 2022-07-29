package main

import "fmt"

func main() {

	numero := 15

	if numero > 15 {
		fmt.Println(numero)
	} else if numero == 5 {
		fmt.Println(5)
	} else {
		fmt.Println(0)
	}

	if outroNumero := numero; outroNumero == 0 {
		// outroNumero vai existir apenas dentro desse scopo
		fmt.Println("Algo")
	}
}
