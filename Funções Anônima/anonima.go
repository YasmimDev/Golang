package main

import "fmt"

func main(){

	retorno := func(texto string) string {
		return fmt.Sprintf("Recebendo -> %s", texto)
	}("Passando Parâmetro")

	fmt.Println(retorno)
}