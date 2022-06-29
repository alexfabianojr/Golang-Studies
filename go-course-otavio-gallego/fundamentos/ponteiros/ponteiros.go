package main

import "fmt"

// Ponteiro é uma referência de memória
// Ao passarmos um ponteiro estamos passando o endereço de memória de uma variável e não o valor dela
func main() {

	// Quando fazemos assim estamos copiando o valor de um para outro
	variavel1 := 1
	variavel2 := variavel1

	variavel1++

	fmt.Println(variavel2)

	// Ponteiro

	var variavel3 int = 100
	var ponteiro *int // declarando um ponteiro

	ponteiro = &variavel3 // o & me retorna o endereço de memória de uma variável

	fmt.Println(ponteiro)

	fmt.Println(*ponteiro) // o * indica que o programa deve acessar o valor guardado no endereço de memória salvo no ponteiro

	variavel3 += 500

	fmt.Println(*ponteiro)
}
