package main

import "fmt"

func main() {

	func() {
		fmt.Println("Olá mundo")
	}()

	minhaFunc := func(texto string) {
		fmt.Println(texto)
	}

	minhaFunc("Hello")

	func(texto string) {
		fmt.Println(texto)
	}("Olarrr")
}
