package main

import (
	"fmt"
	"time"
)

func main() {
	go escrever("Olá Mundo!") // goroutines
	escrever("Programando em Go")
}

func escrever(texto string) {

	for {
		fmt.Print(texto)
		time.Sleep(time.Second)
	}
}
