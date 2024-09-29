package main

import (
	"fmt"
)

func main() {
	const pi float64 = 3.14
	var operacao float64
	var number1 float64
	var number2 float64
	var result float64

	fmt.Println("Qual expressão básica (até 2 números) você deseja executar?")
	fmt.Println("Digite 1 para somar")
	fmt.Println("Digite 2 para subtrair")
	fmt.Println("Digite 3 para multiplicar")
	fmt.Println("Digite 4 para dividir")
	fmt.Scanln(&operacao)

	switch operacao {
	case 1:
		number1 = receiveNumbers()
		number2 = receiveNumbers()
		result = sum(number1, number2)
		fmt.Println(result)
	case 2:
		number1 = receiveNumbers()
		number2 = receiveNumbers()
		result = minus(number1, number2)
		fmt.Println(result)
	case 3:
		number1 = receiveNumbers()
		number2 = receiveNumbers()
		result = multiply(number1, number2)
		fmt.Println(result)
	case 4:
		number1 = receiveNumbers()
		number2 = receiveNumbers()
		result = divide(number1, number2)
		fmt.Println(result)
	default:
		fmt.Println("Opção Inválida, encerrando programa.")
	}
}

func receiveNumbers() float64 {
	var number float64
	fmt.Println("Digite o valor?")
	fmt.Scanln(&number)
	return number
}

func sum(number1 float64, number2 float64) float64 {
	return number1 + number2
}

func minus(number1 float64, number2 float64) float64 {
	return number1 - number2
}

func multiply(number1 float64, number2 float64) float64 {
	return number1 * number2
}

func divide(number1 float64, number2 float64) float64 {
	return number1 / number2
}
