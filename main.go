package main

import (
	"alura/banco/contas"
	"fmt"
)

/*
Create module for import:
go mod init "alura/banco/contas"
*/

func PagarBoleto(conta verificarConta, valorDoBoleto float64) {
	conta.Sacar(valorDoBoleto)
}

type verificarConta interface {
	Sacar(valor float64) string
}

func main() {

	/********************************************************************************/

	contaDenis := contas.ContaPoupanca{}
	contaDenis.Depositar(100)
	PagarBoleto(&contaDenis, 60)
	fmt.Println(contaDenis.ObterSaldo())

	// contaLuisa := contas.ContaCorrente{}
	// contaLuisa.Depositar(500)
	// PagarBoleto(&contaLuisa, 1000)
	// fmt.Println(contaLuisa.ObterSaldo())

	/********************************************************************************/
	// clienteBruno := clientes.Titular{"Bruno", "123.111.123.12", "Desenvolvedor GO"}
	// contaDoBruno := contas.ContaCorrente{
	// 	Titular:       clienteBruno,
	// 	NumeroAgencia: 123,
	// 	NumeroConta:   123456,
	// 	Saldo:         100,
	// }

	// fmt.Println(contaDoBruno)
	/********************************************************************************/
	// 	contaExemplo := contas.ContaCorrente{}
	// 	status, valor := contaExemplo.Depositar(100)
	// 	fmt.Println(status, valor)
	// 	fmt.Println(contaExemplo.ObterSaldo())

	/********************************************************************************/

	/********************************************************************************/
	// var Titular string = "Daniel"
	// var numeroAgencia int = 589
	// var numeroConta int = 123456
	// var Saldo float64 = 125.50

	// contaDoDaniel := ContaCorrente{}
	// contaDoDaniel.Titular = "Daniel"
	// contaDoDaniel.numeroAgencia = 589
	// contaDoDaniel.numeroConta = 123456
	// contaDoDaniel.Saldo = 125.50

	// contaDoDaniel := ContaCorrente{"Daniel", 589, 123456, 125.50}
	// contaDoDaniel2 := ContaCorrente{"Daniel2", 589, 123456, 125.50}

	// contaBruna := ContaCorrente{"Bruna", 222, 111222, 200}
	// contaBruna2 := ContaCorrente{"Bruna", 222, 111222, 200}

	// fmt.Println(contaDoDaniel == contaDoDaniel2)
	// fmt.Println(contaBruna == contaBruna2)

	// var contaDaCris *ContaCorrente
	// contaDaCris = new(ContaCorrente)
	// contaDaCris.Titular = "Cris"
	// contaDaCris.Saldo = 500

	// var contaDaCris2 *ContaCorrente
	// contaDaCris2 = new(ContaCorrente)
	// contaDaCris2.Titular = "Cris"
	// contaDaCris2.Saldo = 500

	// fmt.Println(&contaDaCris)
	// fmt.Pri.ntln(&contaDaCris2)
	// fmt.Println(contaDaCris == contaDaCris2)
	// fmt.Println(*contaDaCris == *contaDaCris2)

	// contaDaSilvia := ContaCorrente{}
	// contaDaSilvia.Titular = "Silvia"
	// contaDaSilvia.Saldo = 500
	// fmt.Println(contaDaSilvia.Saldo)
	// fmt.Println(contaDaSilvia.Sacar(-100))
	// status, valor := contaDaSilvia.Depositar(2000)
	// fmt.Println(status, valor)
	// fmt.Println(contaDaSilvia.Saldo)

	// fmt.Println("Hello!")
	///*****/////

	// contaDaSilvia := contas.ContaCorrente{Titular: "Silvia", Saldo: 300}
	// contaDoGustavo := contas.ContaCorrente{Titular: "Gustavo", Saldo: 100}

	// // status := contaDaSilvia.Transferir(200, &contaDoGustavo)
	// status := contaDoGustavo.Transferir(-100, &contaDaSilvia)

	// fmt.Println(status)
	// fmt.Println(contaDaSilvia)
	// fmt.Println(contaDoGustavo)

}
