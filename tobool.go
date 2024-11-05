package main

import "fmt"

func main() {
	yesNo := "no"

	launch := (yesNo == "yes")
	fmt.Println("Ready for launch: ", launch)

	launch = true
	var oneZero int
	if launch {
		oneZero = 1
	} else {
		oneZero = 0
	}
	fmt.Println("Ready for launch:", oneZero)
}
