package main

import "fmt"

func main() {
	// Funcoes sao um tipo
	// func nome (parametro) retorno

	fmt.Println(somar(1, 2))

	somados := somar(5, 6)

	fmt.Println(somados)

	textoNaTela := func(texto string) {
		fmt.Println(texto)
	}

	textoNaTela("Hello")

	fmt.Println(multiplicar(2, 2))

	fmt.Println(calculos(4, 4))

	somados2, multiplicados := calculos(1, 5)

	fmt.Println(somados2, multiplicados)

	_, multiplicados2 := calculos(1, 5)

	fmt.Println(multiplicados2)
}

func somar(n1 int8, n2 int8) int8 {
	return n1 + n2
}

func multiplicar(n1, n2 int8) int8 {
	return n1 * n2
}

func calculos(n1, n2 int8) (int8, int8) {
	return somar(n1, n2), multiplicar(n1, n2)
}
