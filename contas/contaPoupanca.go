package contas

import (
	"banco/cliente"
	"fmt"
	"strconv"
)

// STRUCT CONTA POUPANÇA
type ContaPoupanca struct {
	Titular                              cliente.Titular
	NumeroAgencia, NumeroConta, Operacao int
	saldo                                float64
}

// FUNÇÃO PARA SACAR
func (c *ContaPoupanca) Sacar(valorDoSaque float64) (string, float64) {
	if valorDoSaque > 0 {
		podeSacar := valorDoSaque <= c.saldo

		if podeSacar {
			c.saldo -= valorDoSaque
			return fmt.Sprintf("Saque de R$%.2f realizado com sucesso! Saldo atual: R$%.2f", valorDoSaque, c.saldo), c.saldo
		} else {
			return fmt.Sprintf("Saldo insuficiente para realizar o saque de R$%.2f. Saldo atual: R$%.2f", valorDoSaque, c.saldo), c.saldo
		}
	} else {
		return fmt.Sprintf("O valor do saque deve ser positivo. Saldo atual: R$%.2f", c.saldo), c.saldo
	}
}

// FUNÇÃO PARA DEPOSITAR
func (c *ContaPoupanca) Depositar(valorDoDeposito float64) (string, float64) {
	if valorDoDeposito > 0 {
		c.saldo += valorDoDeposito
		return "Depósito realizado com sucesso! Saldo atual: R$", c.saldo
	} else {
		return "O valor do depósito deve ser positivo. saldo atual: R$", c.saldo
	}
}

// FUNÇÃO PARA OBTER O SALDO
func (c *ContaPoupanca) ObterSaldo() string {
	saldoStr := strconv.FormatFloat(c.saldo, 'f', 2, 64)
	return c.Titular.Nome + ", Seu saldo atual é: R$" + saldoStr
}
