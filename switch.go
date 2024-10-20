package main

import "fmt"

func main() {
	var room = "lake"

	switch {
	case room == "cave":
		fmt.Println("You find yourself in a dimly lit cavern.")
	case room == "lake":
		fmt.Println("The ice seems solid enough.")
		fallthrough // this one happens by default in Java, C and Javascript, but in GO it take a safer approach, requiring a explicit keyword use.
	case room == "underwater":
		fmt.Println("The water is freezing cold")
	}
}
