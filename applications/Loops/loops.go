package main

import (
	"fmt"
	"time"
	
)

func main(){
	
	
	i := 0
	for i < 10 {
		i++
	fmt.Println("Incrementando i")
	time.Sleep(time.Second)
	}
      fmt.Println(i)

	  for j := 0; j < 10; j++{
		fmt.Println("Incrementandoj", j)
		time.Sleep(time.Second)
	  }

	  nomes := [3]string{"Yasmim", "Alice", "Karina"}

	  for indice, nome := range nomes {  // indice = posição
		fmt.Println(indice, nome)
	  }
	   
	for indice, letra := range "PALAVRA" {
		fmt.Println(indice, string(letra))
	}
	
	usuario := map[string]string {
		"nome": "Alice",
		"Sobrenome":  "Paz",
	}

	for chave, valor := range usuario {
		fmt.Println(chave, valor)
		
	}

	type usuarioStruct struct {
		nome string
		Sobrenome string
	}

	
}