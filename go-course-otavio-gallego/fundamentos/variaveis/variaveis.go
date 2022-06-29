package main

import "fmt"

func main() {
	var variavel1 string = "Modo antigo de se declarar uma variavel"

	variavel2 := "Modo moderno e correto de se declarar uma variavel"

	fmt.Println(variavel1, variavel2)

	var (
		variavel3 string = "A"
		variavel4 string = "B"
	)

	fmt.Println(variavel3, variavel4)

	variavel5, variavel6 := "C", "D"

	fmt.Println(variavel5, variavel6)

	const constante1 string = "E"

	const constante2 = "F"

	fmt.Println(constante1, constante2)
}
