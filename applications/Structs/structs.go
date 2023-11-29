package main

import "fmt"

type usuario struct {
	nome string
	idade uint8
	endereco endereco
}

type endereco struct {
	logadouro string
	numero int8
}

func main() {
	fmt.Println("Arquivo structs")


	var u usuario
	u.nome = "Yas"
	u.idade = 21
     fmt.Println(u)

enderecoExemplo := endereco{"Rua dos bobos", 0}


	 usuario2 := usuario{"Yas", 21, enderecoExemplo}
	 fmt.Println(usuario2)

	 usuario3 := usuario{nome: "Yas"}
	 fmt.Println(usuario3)


}