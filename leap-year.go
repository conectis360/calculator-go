package main

import (
	"fmt"
	"math/rand"
)

var era = "AD"

func main() {
	year := rand.Intn(2024) + 1
	month := rand.Intn(12) + 1
	daysInMonth := 31
	leapYear := year % 4

	switch month {
	case 2:
		if leapYear != 0 {
			daysInMonth = 29
		} else {
			daysInMonth = 29
		}
	case 4, 6, 9, 11:
		daysInMonth = 30
	}

	for x := 0; x < 10; x++ {
		day := rand.Intn(daysInMonth) + 1
		fmt.Printf("date %d \n", x)
		fmt.Println(era, year, month, day)
	}
}
