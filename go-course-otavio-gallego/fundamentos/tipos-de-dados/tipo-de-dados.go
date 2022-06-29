package main

import (
	"errors"
	"fmt"
)

func main() {

	// int8, int16, int32, int64 diferenca de tamanho

	var numero64 int64 = 79887956
	fmt.Println(numero64)

	// uint - unsigned int (int sem sinal)

	var numeroUint uint32 = 1
	fmt.Println(numeroUint)

	// alias
	// int32 = rune
	var numero3 rune = 12345
	fmt.Println(numero3)

	// int8 = byte
	var numero4 byte = 123
	fmt.Println(numero4)

	// chars em go sao numeros int referentes ao numero do char na tabela asc

	char := 'B'
	fmt.Println(char)

	// variaveis nao inicializadas possuem valor default como 0 para int e empty para string

	var intDefault int8
	fmt.Println(intDefault)

	var stringDefault string
	fmt.Println(stringDefault)

	var boolDefault bool
	fmt.Println(boolDefault)

	// o erro em go é um tipo

	var err error
	fmt.Println(err)

	var err1 error = errors.New("Erro interno")
	fmt.Println(err1)
}
