package main

import "fmt"

func runPrimitivesLogic() {
	var name string = "PJ"
	var fact bool = true

	fmt.Println("Hi my name is", name)
	fmt.Println("This is", fact)

	var bankBalance int8 = 127

	fmt.Println("I have R", bankBalance, "in my bank account")

	var debtOutstanding int16 = -20

	fmt.Println("I owe R", debtOutstanding, "to the security guard")

	var weight float32 = 80.1
	fmt.Println("Currently, my weight is:", weight)
}

func main() {
	runPrimitivesLogic()
}
