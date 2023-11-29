package main

import (
	"fmt"
	"time"
)

func main() {
	canal := make(chan string)
	go escrever("Olá Mundo!", <-canal)

	mensagem := <-canal
	fmt.Println(mensagem)

}

func escrever(texto string, canal string) {

	for i := 0; i < 5; i++ {
		//canal <- texto
		time.Sleep(time.Second)
	}
}
