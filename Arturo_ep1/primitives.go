package main

import "fmt"

func runPrimitivesLogic() {
    var name string = "PJ"
    var fact bool = true
    
    fmt.Println("Hi my name is ", &name)  
    fmt.Println("This is ", &fact)

    var bankBalance int8 = 127
    var mem uintptr = 0x1400012c002

    fmt.Println("I have ", &bankBalance, " rand in my bank account")
    fmt.Println("My wallet adress is ", &mem)

    var weight float32 = 175.8
    fmt.Println("Currently, my weight is: ", &weight)
}
