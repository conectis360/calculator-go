package main

import (
	"errors"
	"fmt"
)

func main() {
	const pi float64 = 3.14
	var operacao float64
	var number1 float64
	var number2 float64
	var result float64
	var numbersToOperate int
	arrayOfNumbers := []float64{}

	fmt.Println("Qual expressão básica (até 2 números) você deseja executar?")
	fmt.Println("Digite 1 para somar")
	fmt.Println("Digite 2 para subtrair")
	fmt.Println("Digite 3 para multiplicar")
	fmt.Println("Digite 4 para dividir")
	fmt.Scanln(&operacao)

	switch operacao {
	case 1:
		numbersToOperate = howManyNumbers()
		fmt.Println(numbersToOperate)
		arrayOfNumbers = repeatReceiver(numbersToOperate)

		fmt.Println(sum(arrayOfNumbers))
	case 2:
		result = minus(number1, number2)
		fmt.Println(result)
	case 3:
		result = multiply(number1, number2)
		fmt.Println(result)
	case 4:
		result = divide(number1, number2)
		fmt.Println(result)
	default:
		fmt.Println("Opção Inválida, encerrando programa.")
	}
}

func howManyNumbers() int {
	var number int
	fmt.Println("Digite com quantos algarismos você gostaria de calcular?")
	fmt.Scanln(&number)
	return number
}

func repeatReceiver(quantity int) []float64 {
	var number int = 0
	var number1 float64
	var erro error
	arrayOfNumbers := []float64{}
	for number <= quantity {
		number1, erro = receiveNumbers()
		if erro == nil {
			arrayOfNumbers = append(arrayOfNumbers, number1)
		}
		number++
	}
	return arrayOfNumbers
}

func receiveNumbers() (float64, error) {
	var number float64
	fmt.Println("Digite o valor?")
	value, err := fmt.Scanln(&number) // Check if input parsing fails
	if err != nil {
		return 0, errors.New("math: digite um numero")
	}
	return number, nil
}

func sum(arrayOfNumbers []float64) float64 {
	var number float64
	for i := range len(arrayOfNumbers) {
		number += arrayOfNumbers[i]
	}
	return number
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
