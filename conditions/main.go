package main

import (
	"fmt"
)

// Basic "if" statement 
func firstStatement() {
	if 20 > 18 {
		fmt.Println(true)
	}
}

// "if" statement with variables
func secondStatement() {
	x := 20
	y := 18
	if x > y {
		fmt.Println("y is smaller than x")
	}
}

// "if else" statement 
func thirdStatement() {
	time := 20
	if time < 18 {
		fmt.Println("It's a beautiful day")
	} else {
		fmt.Println("Good evening.")
	}
}

// "if else if else" statement (multiple conditions)
func fourthStatement() {
	time := 18
	if time < 10 {
		fmt.Println("Good morning!")
	} else if time < 20 {
		fmt.Println("Good day.")
	} else {
		fmt.Println("Good evening.")
	}
}

func fithStatement()  {
    num := 9
    if num >= 10 {
        fmt.Println("Number is bigger than 10")
        if num > 15 {
            fmt.Println("Number is also bigger than 15")
        }
    } else {
        fmt.Println("Number is less than 10")
    }
}

func main() {
	firstStatement()
	secondStatement()
	thirdStatement()
	fourthStatement()
    fithStatement()
}


