package main

import "fmt"

func main() {
    fmt.Print("Enter value in meters: ")
    var input float64
    var MtoF = 3.28084
    fmt.Scanf("%f", &input)   

    output := input * MtoF

    fmt.Println("Result in feet:", output)
}