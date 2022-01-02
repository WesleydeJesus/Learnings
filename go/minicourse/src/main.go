package main

import (
	"Learnings/go/minicourse/contas"
	"fmt"
)

type verificaConta  interface{
	Saque(valor float64) string
}

func PagarBoleto( conta verificaConta, valor float64){
	conta.Saque(valor)
}


func main(){
	cCorrente := contas.ContaCorrente{}
	cPoupanca := contas.ContaPoupanca{}
	fmt.Println(cCorrente, cPoupanca)
}

