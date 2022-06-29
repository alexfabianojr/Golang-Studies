package main

import "fmt"

func main() {
	//  Aritmeticos
	soma := 1 + 2
	subtracao := 1 - 2
	divisao := 10 / 4
	multiplic := 5 * 2
	resto := 8 % 9

	fmt.Println(soma, subtracao, divisao, multiplic, resto)

	fmt.Println(1 == 2)
	fmt.Println(1 > 2)
	fmt.Println(1 >= 2)
	fmt.Println(1 <= 2)
	fmt.Println(1 != 2)

	fmt.Println(true && false)
	fmt.Println(!false)

	numero := 10
	numero++
	numero += 10
	fmt.Println(numero)

	numero--
	numero -= 10
	fmt.Println(numero)
}
