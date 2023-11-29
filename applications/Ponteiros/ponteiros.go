package main

import "fmt"

func main(){

var var1 int = 10
var var2 int = var1

fmt.Println(var1, var2)

// apenas o valor da var1 será alterado

var1++
fmt.Println(var1, var2)

//Salvar um endereço de memória;

var var3 int
var ponteiro *int

var3 = 100
ponteiro = &var3

fmt.Println(var3, *ponteiro) //desreferenciação

} 