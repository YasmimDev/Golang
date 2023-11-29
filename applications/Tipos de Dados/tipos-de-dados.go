package main

import (
	"errors"
	"fmt"
)

func main(){

	

	numero :=  -100000000000000
	fmt.Println(numero)

	var numero2 uint32 = 10000 //int sem sinal
	fmt.Println(numero2)

	//alias

	var numero3 rune = 12345 //int32 = rune
	fmt.Println(numero3)

	var numero4 byte = 123 //byte = int8
	fmt.Println(numero4)


//----------------------------------------------------------------
	
	var numeroReal1  float32 = 123.65
	fmt.Println(numeroReal1)

	var numeroReal2  float64 = 123000000000000.65
	fmt.Println(numeroReal2)

	numeroReal3 := 12345.90
	fmt.Println(numeroReal3)

	//-------------------------------------------------------

	var str string = "texto"
	fmt.Println(str)

	str2 := "texto2"
	fmt.Println(str2)

	char := 'B'
	fmt.Println(char)

	//--------------------

	var texto int
	fmt.Println(texto)
 
	//---------------------------

	var booleano1 bool
	fmt.Println(booleano1)

	//---------------------------------

	var erro error = errors.New("Erro interno") 
	fmt.Println(erro)

	//----------------------------------



}