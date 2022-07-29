package main

import "fmt"

func main() {

	numero := 4

	// modo tradicional para a mesma variavel

	switch numero {
	case 1:
		fmt.Println("Segunda")
	case 2:
		fmt.Println("Terça")
	case 3:
		fmt.Println("Quarta")
	case 4:
		fmt.Println("Quinta")
	case 5:
		fmt.Println("Sexta")
		fallthrough // vai jogar a execucao para dentro do proximo corpo de condicao  (sem avaliar a condicao)
	default:
		fmt.Println("Fim de semana")
	}

	//modo para avaliar mais de uma variavel
	outroNumero := 6

	switch {
	case numero == 1:
		fmt.Println("Coisa")
	case outroNumero == 6:
		fmt.Println("Outra coisa")
	}
}
