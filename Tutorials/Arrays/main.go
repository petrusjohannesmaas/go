package main

import "fmt"

var bandName string = "Shongololo"
var album string = "Twin Turbo"

func firstArray () {
    instruments := [4] string {"Guitar", "Bass", "Drums", "Keyboard"}
    fmt.Println("My band's name is", bandName, "and I play:", instruments[2])
}

func secondArray()  {
    members := [3] string {"Ruan", "Pj", "Josh"}
    fmt.Println("We are", len(members), "band members")
    fmt.Println("Our names are", members)
}

func thirdArray()  {
    songs := [5] string {
        "Vat So",
        "Swart Hond",
        "Mrs. Krouse",
        "Haat My",
        "Krag",        
    }

    fmt.Println("We have", len(songs), "tracks and the album is called", album)
    fmt.Println("Our next single is", songs[1])
}

func main () {
    firstArray()
    secondArray()
    thirdArray()
}
