package main

import (
	"banco/cliente"
	"banco/contas"
	"fmt"
)

// FUNÇÃO PARA PAGAR BOLETO, UTILIZANDO INTERFACE
func PagarBoleto(conta verificarConta, valorDoBoleto float64) string {
	// Chamando a função Sacar que realiza o pagamento
	conta.Sacar(valorDoBoleto)
	// Retornando a mensagem de sucesso e o saldo atualizado
	return fmt.Sprintf("Pagamento do boleto de R$%.2f realizado com sucesso! %s", valorDoBoleto, conta.ObterSaldo())
}

// INTERFACE PARA VERIFICAR QUAL É O TIPO DE CONTA
type verificarConta interface {
	Sacar(valor float64) (string, float64)
	ObterSaldo() string
}

func main() {
	// Criando conta exemplo, atribuindo apenas o nome
	clienteExemplo := cliente.Titular{Nome: "Gabriel"}
	contaExemplo := contas.ContaCorrente{
		Titular: clienteExemplo}
	// Depositando na conta exemplo
	contaExemplo.Depositar(10000)

	// Exibindo o saldo da conta
	fmt.Println(contaExemplo.ObterSaldo())

	// -----------------------------------------------

	// Criando a conta poupança do Giorgian
	contaDoGiorgian := contas.ContaPoupanca{
		Titular: cliente.Titular{
			Nome:      "Giorgian",
			Cpf:       "111.111.111.11",
			Profissao: "Jogador de Futebol",
		},
		NumeroAgencia: 123,
		NumeroConta:   987654,
		Operacao:      2,
	}

	// Depositando um valor na conta do Giorgian
	fmt.Println(contaDoGiorgian.Depositar(20000))

	// Pagando o boleto utilizando a função com a interface
	fmt.Println(PagarBoleto(&contaDoGiorgian, 2313))
}
