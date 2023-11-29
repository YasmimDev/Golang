package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
)

type cachorro struct {
	Nome  string `json:"nome"`
	Raça  string `json:"raça"`
	Idade uint   `json:"idade"`
}

func main() {

	c := cachorro{"Rex", "Dalmata", 1}

	cachorroEmJSON, erro := json.Marshal(c)
	if erro != nil {
		log.Fatal(erro)
	}
	fmt.Println(cachorroEmJSON)

	fmt.Println(bytes.NewBuffer(cachorroEmJSON))

	c2 := map[string]string{
		"nome": "Max",
		"raça": "beagle",
	}

	cachorroEmJSON, err := json.Marshal(c2)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(cachorroEmJSON)
	fmt.Println(bytes.NewBuffer(cachorroEmJSON))
}
