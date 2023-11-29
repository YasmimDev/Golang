package main

import "fmt"

func main(){
	var variavel1 string = "variavel 1"
	variavel2 := "variavel 2" 

	fmt.Println(variavel1)
	fmt.Println(variavel2)

//declaração mais de uma váriavel
	var ( 
		variavel3 string = "lala"
		variavel4 string = "titi"
	)

	fmt.Println(variavel3, variavel4)

	variavel5, variavel6 := "variavel 5", "variavel 6"
	fmt.Println(variavel5, variavel6)

//declaração de uma constante
const constante1 string = "constante 1"
fmt.Println(constante1)


//inverter o valor da variável
variavel5, variavel6 = variavel6, variavel5
fmt.Println(variavel5, variavel6)

}