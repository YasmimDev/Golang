package main

import "fmt"

func main(){

	usuario := map[string]string{

		"nome":    "Yasmim",
		"sobrenome":  "Paz",

	}
    fmt.Println(usuario)

	usuario_2 := map[string]map[string]string{
		"nome": {
			"Primeiro": "Yas",
			"último":   "Victoria",
		},

		"curso": {
			"nome": "ADS",
			"campus": "campus 2",
		},
	}  

	fmt.Println(usuario_2)
	delete(usuario_2, "nome")
	fmt.Println(usuario_2)

	usuario_2["signo"] = map[string]string{
		"nome": "Libra",
	}
	fmt.Println(usuario_2)
}     

