package main

import (
	"errors"
	"fmt"
	"log"
)

func main() {
	var operacao float64

	fmt.Println("Qual expressão básica (até 2 números) você deseja executar?")
	fmt.Println("Digite 1 para somar")
	fmt.Println("Digite 2 para subtrair")
	fmt.Println("Digite 3 para multiplicar")
	fmt.Println("Digite 4 para dividir")
	fmt.Scanln(&operacao)

	switch operacao {
	case 1:
		handleOperation(sum)
	case 2:
		handleOperation(minus)
	case 3:
		handleOperation(multiply)
	case 4:
		handleOperationWithError(divide)
	default:
		fmt.Println("Erro: opção Inválida, encerrando programa.")
	}
}

func handleOperationWithError(operation func([]float64) (float64, error)) {
	numbersToOperate, err := howManyNumbers()
	if err != nil {
		fmt.Println(err)
		return
	}
	arrayOfNumbers := repeatReceiver(numbersToOperate)
	result, err := operation(arrayOfNumbers)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Result:", result)
}

func handleOperation(operation func([]float64) float64) {
	numbersToOperate, err := howManyNumbers()
	if err != nil {
		fmt.Print(err)
		return
	}
	arrayOfNumbers := repeatReceiver(numbersToOperate)
	result := operation(arrayOfNumbers)
	fmt.Println(result)
}

func howManyNumbers() (int, error) {
	var number int
	fmt.Println("Digite com quantos algarismos você gostaria de calcular?")
	_, err := fmt.Scanln(&number)
	if err != nil || number <= 0 {
		return 0, errors.New("digite numeros positivos por favor")
	}
	return number, nil
}

func repeatReceiver(quantity int) []float64 {
	arrayOfNumbers := make([]float64, 0, quantity) // Pre-allocate slice
	for i := 0; i < quantity; i++ {
		number, err := receiveNumbers()
		if err != nil {
			fmt.Println(err)
			return arrayOfNumbers
		}
		arrayOfNumbers = append(arrayOfNumbers, number)
	}
	return arrayOfNumbers
}

func receiveNumbers() (float64, error) {
	var number float64
	fmt.Println("Digite o valor?")
	_, err := fmt.Scanln(&number)
	if err != nil {
		log.Println("LOG: ", err)
		return 0, errors.New("por favor digite um numero")
	}
	return number, nil
}

func sum(arrayOfNumbers []float64) float64 {
	var result float64 = 0
	for _, num := range arrayOfNumbers {
		result += num
	}
	return result
}

func minus(arrayOfNumbers []float64) float64 {
	result := arrayOfNumbers[0]
	for _, num := range arrayOfNumbers[1:] {
		result -= num
	}
	return result
}

func multiply(arrayOfNumbers []float64) float64 {
	result := arrayOfNumbers[0]
	for _, num := range arrayOfNumbers[1:] {
		result *= num
	}
	return result
}

func divide(arrayOfNumbers []float64) (float64, error) {
	result := arrayOfNumbers[0]
	for _, num := range arrayOfNumbers[1:] {
		if num == 0 {
			return 0, errors.New("não é possível dividir por zero")
		}
		result /= num
	}
	return result, nil
}
