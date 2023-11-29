package main

import "fmt"

func main(){
 
	//ARITMETICOS

	soma := 1 + 2
	subtracao := 1 - 2
	divisao := 10 / 4
	multiplicacao := 10 * 5
	restoDaDivisao := 10 % 2

	fmt.Println(soma, subtracao, divisao, multiplicacao, restoDaDivisao)


var numero1 int16 = 10
var numero2  int16 = 25
soma2 := numero1 + numero2
fmt.Println(soma2)





//-------------------------------

// ATRIBUIÇÃO

var variavel string = "String"
variavel2 := "String2"
fmt.Println(variavel, variavel2)

//--------------------------------

//OPERADORES RELACIONAIS

fmt.Println(1 > 2)
fmt.Println(1 < 2)
fmt.Println(1 >= 2)
fmt.Println(1 == 2)
fmt.Println(1 <= 2)
fmt.Println(1 != 2)

//--------------------------------------

// OERADORES LÓGICOS

verdadeiro, falso := true, false
fmt.Println(verdadeiro && falso)
fmt.Println(verdadeiro || falso)
fmt.Println(!verdadeiro)
fmt.Println(!falso)

//---------------------------------------------

//  OPERADORES UNÁRIOS

numero := 10
numero++
numero += 13 // numero = numero + 13
fmt.Println(numero)

numero--
numero -= 20 //numero = numero - 20

numero *= 3  //  numero = numero * 3
numero /= 10
numero %= 3

//---------------------------------------------------

//OPERADORES TERNARIOS
var  texto string
if numero > 5{
	texto = "maior que 5"
} else {
	texto = "menor que 5"
}

fmt.Println(texto)

}