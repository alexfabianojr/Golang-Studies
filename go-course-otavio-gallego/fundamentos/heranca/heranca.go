package main

import "fmt"

type pessoa struct {
	nome      string
	sobrenome string
	idade     uint8
	altura    uint8
}

// em go usamos composicao no lugar de heranca mas nao como declaracao de variavel em si como na composicao de java

type estudante struct {
	pessoa    // "heranca" - composicao direta sem precisar declarar variavel e tipo de variavel
	curso     string
	faculdade string
}

func main() {
	p1 := pessoa{nome: "Alex", sobrenome: "Fabiano", idade: 20, altura: 123}
	e1 := estudante{p1, "TI", "QI"}

	fmt.Println(e1.nome)
	fmt.Println(e1.sobrenome)
}
