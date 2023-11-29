package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	var waitGroup sync.WaitGroup
	waitGroup.Add(2)

	/*go func ()  {
		escrever("olá mundo")
		waitGroup.Done() // -1
	}{}

	go func ()  {
		escrever("programando em go")
		waitGroup.Done() // -1
	}{}*/

	waitGroup.Wait() // 2
}

func escrever(texto string) {

	for i := 0; i < 5; i++ {
		fmt.Print(texto)
		time.Sleep(time.Second)
	}
}
