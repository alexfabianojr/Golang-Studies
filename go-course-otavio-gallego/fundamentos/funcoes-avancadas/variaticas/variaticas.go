package main

import "fmt"

// vai receber de 1 a N ints
func soma(numeros ...int) int {
	total := 0
	for _, numero := range numeros {
		total += numero
	}

	return total
}

func main() {
	fmt.Println(soma(1, 2, 3, 4))
	fmt.Println(soma(1, 2, 3, 4, 5, 6, 7, 8, 9, 10))
	fmt.Println(soma())
}
