package main

import "fmt"

func main() {
	howMuchMoney := 0
	cents := 20
	count := 100

	for i := 0; i < count; i++ {
		if cents >= 100 {
			howMuchMoney += 1
			cents = 0
		}
		fmt.Printf("%v.%v \n", howMuchMoney, cents)
		cents += 20
	}
}
