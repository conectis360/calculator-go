package main

import (
	"fmt"
)

func main() {
	third := 1.0 / 3
	fmt.Println(third)
	fmt.Printf("%v\n", third)
	fmt.Printf("%f\n", third)
	fmt.Printf("%.3f\n", third)
	fmt.Printf("%4.2f\n", third) // 4. width 2f precision, which specifies how many digits should appear after the decimal point
}
