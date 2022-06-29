package main

import "fmt"

func main() {

	usuario := map[string]string{
		"nome":      "Alex",
		"sobrenome": "Silva",
	}

	fmt.Println(usuario["nome"])
	fmt.Println(usuario["sobrenome"])

	usuario2 := map[string]map[string]string{
		"nome": {
			"primeiro": "Alex",
			"segundo":  "Silva",
		},
	}

	fmt.Println(usuario2["nome"]["segundo"])

	delete(usuario, "sobrenome")
}
