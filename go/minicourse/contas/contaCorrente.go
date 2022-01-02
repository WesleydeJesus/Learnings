package contas

import "Learnings/go/minicourse/clientes"

type ContaCorrente struct{
	Titular clientes.Titular
	NAgencia, NConta int
	saldo float64
}


func (c *ContaCorrente) Saque (valor float64) string {
	saque := valor > 0 && valor <= c.saldo
	if saque {
		c.saldo -= valor
		return "Saque Efetuado \n " 
	} else {
		return "saldo insuficiente \n "
	} 
}

func (c *ContaCorrente) Deposito (valor float64) (string, float64) {
	if valor > 0 {
		c.saldo += valor
		return "Valor depositado com sucesso", c.saldo
	}else {
		return "Algo deu errado", c.saldo
	}
}

func (c *ContaCorrente) Transferencia (valor float64, conta *ContaCorrente) (bool){
	if valor < c.saldo && valor > 0 {
		c.saldo -= valor
		conta.Deposito(valor)
		return true;
	}else {
		return false
	}
}

func (c *ContaCorrente) ObterSaldo () float64 {
	return c.saldo
}