package main

import (
	"fmt"

	"./data/handlers"
	"./data/utils"
)

func main() {
	fmt.Println("Main function running...")
	handlers.HandleThis()
	utils.DoSomething()
}
