package main

import (
	"fmt"
	"reflect"
)

func main() {

	var arr1 [5]int
	arr1[2] = 99

	fmt.Println(arr1)

	arr2 := [3]string{"Oi", "Tchau", "Olar"}

	fmt.Println(arr2)

	// Array "flexivel"

	// esse modo n é mt usado
	arr3 := [...]int{1, 2, 3, 4, 5, 6} //O tamanho do array sera a quantidade de itens passados na inicialização

	fmt.Println(arr3)

	//Slice estrutura feita com base em array mas mais flexivel
	//No slice nao temos tamanho fixo
	// slice e um pedaco de um array
	slice1 := []int{1, 2, 3, 4, 5, 6, 7}

	fmt.Println(reflect.TypeOf(slice1))

	slice1 = append(slice1, 10) // retorna o slice antigo + o item novo

	// Arrays internos
	// A funcao make inicializa um objeto destes 3 tipos apenas: slice, map, or chan
	// Se nao passarmos a capacidade maxima dele o slice comeca com capacidade igual ao tamanho
	slice2 := make([]float32, 10, 15) //cria um array de 15 posicoes (capacidade maxima) e retorna um slice com as 10 primeiras posicoes (array interno)

	fmt.Println(slice2)
	fmt.Println(len(slice2)) // length tamanho
	fmt.Println(cap(slice2)) // capacidade

	slice2 = append(slice2, 5)

	fmt.Println(slice2)
	fmt.Println(len(slice2)) // length tamanho
	fmt.Println(cap(slice2)) // capacidade

	slice2 = append(slice2, 5)
	slice2 = append(slice2, 5)
	slice2 = append(slice2, 5)
	slice2 = append(slice2, 5)
	slice2 = append(slice2, 5)
	slice2 = append(slice2, 5)

	// O slice nunca tem um tamanho limite
	// Toda vez que chegarmos no tamanho máximo ele vai alocar mais

	fmt.Println(slice2)
	fmt.Println(len(slice2)) // length tamanho
	fmt.Println(cap(slice2)) // capacidade
}
