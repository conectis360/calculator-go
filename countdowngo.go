package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var count = 10

	for count > 0 {
		fmt.Println(count)
		time.Sleep(time.Second)
		if rand.Intn(100) == 0 {
			break
		}
		count--
	}

	if count == 0 {
		fmt.Println("LiftOff!")
	} else {
		fmt.Println("Launch Failed!")
	}
}
