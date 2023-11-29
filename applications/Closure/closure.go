package main

import "fmt"


func closure() func (){
	texto := "Função Closure"

	funcao := func ()  {
		fmt.Println(texto)
	}

	return funcao
}

func main(){
	texto := "Dentro da Função"
	fmt.Println(texto)

	funcaoNova := closure()
	funcaoNova()
} 