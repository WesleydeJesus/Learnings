package contas

import "Learnings/go/minicourse/clientes"


type ContaPoupanca struct {
	Titular clientes.Titular
	NAgencia, NConta, Operacao int
	saldo float64
}

func (c *ContaPoupanca) Saque (valor float64) string {
	saque := valor > 0 && valor <= c.saldo
	if saque {
		c.saldo -= valor
		return "Saque Efetuado \n " 
	} else {
		return "saldo insuficiente \n "
	} 
}

func (c *ContaPoupanca) Deposito (valor float64) (string, float64) {
	if valor > 0 {
		c.saldo += valor
		return "Valor depositado com sucesso", c.saldo
	}else {
		return "Algo deu errado", c.saldo
	}
}

func (c *ContaPoupanca) ObterSaldo () float64 {
	return c.saldo
}