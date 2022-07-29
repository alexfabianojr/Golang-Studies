package main

import "fmt"

// defer adia a execução de um pedaço de código
// A funcao eh executada antes do fim da execucao do escopo onde ela está inserida
func main() {

	// defer = ADIAR
	defer funcao1()
	funcao2()
}

func funcao1() {
	fmt.Println("Executando a funcao 1")
}

func funcao2() {
	fmt.Println("Executando a funcao 2")
}
