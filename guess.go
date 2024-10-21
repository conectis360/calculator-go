package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var magicNumber = rand.Intn(100)
	var number int
	fmt.Println("Digite com quantos algarismos você gostaria de calcular?")
	_, error := fmt.Scanln(&number)
	if error != nil {
		println("error: " + error.Error())
	}
	if magicNumber == number {
		fmt.Println("congratulations!")
	} else {
		fmt.Printf("sadly it's another number %v", magicNumber)
	}
}
