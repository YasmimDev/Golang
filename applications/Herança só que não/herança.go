package main

import "fmt"

type pessoa struct {
      nome string
	  sobrenome string
	  idade uint8
	  altura uint8

}

type estudante struct {
	pessoa
	curso string
	faculdade string

}

func main() {
    fmt.Println("Herança")
	

	p1 := pessoa{"Yasmim", "Paz", 18, 160}
	fmt.Println(p1)

	e1 := estudante{p1, "Análise", "Unit"}
	fmt.Println(e1)
}
