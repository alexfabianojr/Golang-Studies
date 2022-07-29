package main

import "fmt"

func alunoEstaAprovado(n1, n2 float64) bool {
	defer recuperarExecucao()

	media := (n1 + n2) / 2

	if media > 6 {
		return true
	} else if media < 6 {
		return false
	}

	panic("Media eh 6")
}

func recuperarExecucao() {
	if recover := recover(); recover != nil {
		fmt.Println("Execucao recuperada")
	}
}

func main() {
	fmt.Println(alunoEstaAprovado(6, 6))
}
