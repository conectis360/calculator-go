package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var magicNumber = 7
	var tries = 0
	for {
		if magicNumber == rand.Intn(100) {
			fmt.Printf("congratulations! you did it in %v tries", tries)
			break
		} else {
			tries++
			fmt.Println("trying again!")
		}
	}
}
