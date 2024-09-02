package main

import (
	"fmt"
	"hello-world/doctor"
)

func main() {
	var whatToSay string

	whatToSay = doctor.Intro()

	fmt.Println(whatToSay)
}
