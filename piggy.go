package main

import (
	"fmt"
	"math/rand"
)

func main() {
	piggyBank := 0.0
	days := 120

	for i := 0; i < days; i++ {
		piggyBank += rand.Float64()
		if piggyBank > 20.0 {
			break
		}
	}
	fmt.Printf("Ok now I do have %0.2f", piggyBank)
}
