package main

// type usuario struct{
// 	nome string
// 	idade int
// }

// func (u usuario) salvar(){
// 	fmt.Println("salvar no metodo")
// }

// func main(){
// 	usuario := usuario{"usuario 1,", 20}
// 	fmt.Println(usuario)
// 	usuario.salvar()
// }

type Conta struct {
	Saldo int
	Nome  string
}

func (c *Conta) nome() string {
	return c.Nome
}

func (c *Conta) saldo() int {
	return c.Saldo
}

func (c *Conta) Setnome(valor string) {
	c.Nome = valor
}

func (c *Conta) SetSaldo(valor int) {
	c.Saldo = valor
}
