package main

import (
	"errors"
	"fmt"
)

func main() {
	var operacao float64
	arrayOfNumbers := []float64{}

	fmt.Println("Qual expressão básica (até 2 números) você deseja executar?")
	fmt.Println("Digite 1 para somar")
	fmt.Println("Digite 2 para subtrair")
	fmt.Println("Digite 3 para multiplicar")
	fmt.Println("Digite 4 para dividir")
	fmt.Scanln(&operacao)

	switch operacao {
	case 1:
		numbersToOperate, err := howManyNumbers()
		if err != nil {
			fmt.Print(err)
			break
		}
		fmt.Println(numbersToOperate)
		arrayOfNumbers = repeatReceiver(numbersToOperate)

		fmt.Println(sum(arrayOfNumbers))
	case 2:
		numbersToOperate, err := howManyNumbers()
		if err != nil {
			fmt.Print(err)
			break
		}
		fmt.Println(numbersToOperate)
		arrayOfNumbers = repeatReceiver(numbersToOperate)
		fmt.Println(minus(arrayOfNumbers))
	case 3:
		numbersToOperate, err := howManyNumbers()
		if err != nil {
			fmt.Print(err)
			break
		}
		fmt.Println(numbersToOperate)
		arrayOfNumbers = repeatReceiver(numbersToOperate)
		fmt.Println(multiply(arrayOfNumbers))
	case 4:
		numbersToOperate, err := howManyNumbers()
		if err != nil {
			fmt.Print(err)
			break
		}
		fmt.Println(numbersToOperate)
		arrayOfNumbers = repeatReceiver(numbersToOperate)
		fmt.Println(divide(arrayOfNumbers))
	default:
		fmt.Println("Erro: opção Inválida, encerrando programa.")
	}
}

func howManyNumbers() (int, error) {
	var number int
	fmt.Println("Digite com quantos algarismos você gostaria de calcular?")
	_, err := fmt.Scanln(&number)
	if err != nil {
		return 0, errors.New("digite apenas a quantidade de algoritmos por favor")
	}
	return number, err
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
		} else {
			fmt.Println(erro)
			break
		}
		number++
	}
	return arrayOfNumbers
}

func receiveNumbers() (float64, error) {
	var number float64
	fmt.Println("Digite o valor?")
	_, err := fmt.Scanln(&number) // scan has a double return, an integer and a error, using err to return a error
	if err != nil {
		return 0, errors.New("por favor digite um numero") // return 0 and the error for the caller (main function in our case)
	}
	return number, nil // no errors, return the scanned number
}

func sum(arrayOfNumbers []float64) float64 {
	var number float64
	for i := range len(arrayOfNumbers) {
		number += arrayOfNumbers[i]
	}
	return number
}

func minus(arrayOfNumbers []float64) float64 {
	var number float64
	for i := range len(arrayOfNumbers) {
		number += arrayOfNumbers[i]
	}
	return number
}

func multiply(arrayOfNumbers []float64) float64 {
	var number float64
	for i := range len(arrayOfNumbers) {
		number *= arrayOfNumbers[i]
	}
	return number
}

func divide(arrayOfNumbers []float64) float64 {
	var number float64
	for i := range len(arrayOfNumbers) {
		number /= arrayOfNumbers[i]
	}
	return number
}
