package contas

import (
	"banco/cliente"
	"strconv"
)

// STRUCT CONTA CORRENTE
type ContaCorrente struct {
	Titular                    cliente.Titular
	NumeroAgencia, NumeroConta int
	saldo                      float64
}

// FUNÇÃO PARA SACAR
func (c *ContaCorrente) Sacar(valorDoSaque float64) (string, float64) {
	if valorDoSaque > 0 {
		podeSacar := valorDoSaque <= c.saldo

		if podeSacar {
			c.saldo -= valorDoSaque
			return "Saque realizado com sucesso! saldo atual: R$", c.saldo
		} else {
			return "saldo insuficiente para realizar o saque. saldo atual: R$", c.saldo
		}
	} else {
		return "O valor do saque deve ser positivo. saldo atual: R$", c.saldo
	}
}

// FUNÇÃO PARA DEPOSITAR

func (c *ContaCorrente) Depositar(valorDoDeposito float64) (string, float64) {
	if valorDoDeposito > 0 {
		c.saldo += valorDoDeposito
		return "Depósito realizado com sucesso! saldo atual: R$", c.saldo
	} else {
		return "O valor do depósito deve ser positivo. saldo atual: R$", c.saldo
	}
}

// FUNÇÃO PARA TRANSFERIR ENTRE CONTAS

func (c *ContaCorrente) Transferir(valorDaTransferencia float64, contaDestino *ContaCorrente) (string, float64) {
	if valorDaTransferencia <= c.saldo {
		c.saldo -= valorDaTransferencia
		contaDestino.Depositar(valorDaTransferencia)
		// Convertendo o saldo para string manualmente
		saldoStr := strconv.FormatFloat(c.saldo, 'f', 2, 64)
		// Acessando o nome do titular para realizar a concatenação
		return "Transferência realizada com sucesso para " + contaDestino.Titular.Nome +
			". saldo atual: R$" + saldoStr, c.saldo
	} else {
		return "O valor da transferência é maior que o saldo disponível.", c.saldo
	}
}

// FUNÇÃO PARA OBTER SALDO

func (c *ContaCorrente) ObterSaldo() string {
	saldoStr := strconv.FormatFloat(c.saldo, 'f', 2, 64)
	return c.Titular.Nome + ", seu saldo atual é: R$" + saldoStr
}
