package main

import (
	"fmt"
	"myapp/doctor"
)

func main() {
	whatToSay := doctor.Intro()

	fmt.Println(whatToSay)
}
