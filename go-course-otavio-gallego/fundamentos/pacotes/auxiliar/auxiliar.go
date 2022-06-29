package auxiliar

import "fmt"

//Escrever mensagem auxiliar  (boa pratica do go: comentário em cima de função public além do pacote)
//Letra minuscula e maiuscula define visibilidade fora do pacote ou n
func Escrever() {
	fmt.Println("Escrevendo no pacote auxiliar")
	escrevendo2()
}
