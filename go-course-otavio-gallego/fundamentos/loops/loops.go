package main

import "fmt"

func main() {

	i := 0

	for i < 10 {
		fmt.Println(i)
		i++
	}

	for j := 10; j <= 10 && j > -10; j-- {
		fmt.Println(j)
	}

	// usamos o range  para iterar em slices e arrays

	nomes := []string{"Joao", "Paulo", " Gabriel"}

	for indice, valor := range nomes {
		fmt.Println(indice, valor)
	}

	for _, valor := range nomes {
		fmt.Println(valor)
	}

	for indice, letra := range "PALAVRA" {
		fmt.Println(indice, string(letra))
	}
}
