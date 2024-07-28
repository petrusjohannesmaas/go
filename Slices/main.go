package main

import "fmt"

func firstSlice() {
    fmt.Println("This is the first function ###")
	mySlice := [] int {1, 2, 3}
	fmt.Println(len(mySlice))
}

func secondSlice() {
    fmt.Println("This is the second function ###")
    quote := [] string {"Must", "have", "no", "preferences"}
    fmt.Println(cap(quote))
    fmt.Println(quote)
}

func thirdSlice () {
    fmt.Println("This is the third function ###")
    arr1 := [7] string {"a", "b", "c", "d", "e", "f", "g"}
    testSlice := arr1 [2:5]
    fmt.Printf("testSlice = %v\n", testSlice)
    fmt.Printf("length = %d\n", len(testSlice))
    fmt.Printf("capacity = %d\n", cap(testSlice))
}

func fourthSlice () {
    fmt.Println("This is the fourth function ###")
    fooSlice := make([] int, 5, 10 )
    fmt.Printf("fooSlice = %v\n", fooSlice)
    fmt.Printf("length = %d\n", len(fooSlice))
    fmt.Printf("capacity = %d\n", cap(fooSlice))
}

func main() {
	firstSlice()
    secondSlice()
    thirdSlice()
    fourthSlice()
}
