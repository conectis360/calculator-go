package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("You find yourself in a dimly lit cavern.")

	var command = "walk outside"
	var exit = strings.Contains(command, "outside")

	fmt.Println("You leave the cave:", exit)

	var wearShades = true
	fmt.Println("did you wear the shades? ", wearShades)

	var commandRead = "read the sign"
	var read = strings.Contains(commandRead, "read")
	fmt.Println(read)
}
