package main

import "fmt"

func calculosMatematicos(n1, n2 int) (soma int, subtracao int) {
	soma = n1 + n2 //como temos retorno nomeado nao precisamos declarar nova variavel
	subtracao = n1 - n2
	return // como demos os valores dos retornos nomeados podemos apenas dar return
}

func main() {
	fmt.Println(calculosMatematicos(1, 2))
}
