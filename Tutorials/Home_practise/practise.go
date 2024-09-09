package main

import "fmt"

func lineOne() {
    var name string = "test"
    fmt.Println("This is a", name)
}

func lineTwo()  {
    var difficulty bool = false                
    fmt.Println("Is it easy?", difficulty) 
}

func lineThree()  {
    var verb string = "having"                   
    fmt.Println("Am I", verb, "a good time?")
}

func lineFour() {
    var percentage int = 100
    fmt.Println("Yes,", percentage, "%")
}

func main() {
    lineOne()
    lineTwo()
    lineThree()
    lineFour()
}


