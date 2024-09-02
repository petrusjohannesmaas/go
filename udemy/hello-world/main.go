package main

import (
	_ doctor
	"fmt"
)

func main() {
	var whatToSay string

	whatToSay = doctor.Intro()

	fmt.Println(whatToSay)
}
