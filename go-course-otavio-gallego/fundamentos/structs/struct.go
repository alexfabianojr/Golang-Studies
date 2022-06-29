package main

import "fmt"

// Nao temos classes em go mas structs
// type [nome] struct {}

type usuario struct {
	nome     string
	idade    int8
	endereco endereco
}

type endereco struct {
	rua    string
	numero int8
}

func main() {
	var u1 usuario
	u1.nome = "Alex"
	u1.idade = 20

	fmt.Println(u1)

	u2 := usuario{nome: "Joao", idade: 21}

	fmt.Println(u2)

	u3 := usuario{
		"Rafael",
		30,
		endereco{rua: "Avenida", numero: 2},
	}

	fmt.Println(u3)
}
