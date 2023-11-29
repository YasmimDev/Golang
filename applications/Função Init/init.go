package main

import "fmt"

var n int
func init(){
	fmt.Println("Função Init")
	n = 10
}


func main(){
	fmt.Println("Função Main")
	fmt.Println(n)
}