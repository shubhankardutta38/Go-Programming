package main

import "fmt"

func main() {
    var num1, num2 int

    // Taking input from the user for the first number
    fmt.Print("Enter the first number: ")
    _, err := fmt.Scan(&num1)
    if err != nil {
        fmt.Println("Error reading input:", err)
        return
    }

    // Taking input from the user for the second number
    fmt.Print("Enter the second number: ")
    _, err = fmt.Scan(&num2)
    if err != nil {
        fmt.Println("Error reading input:", err)
        return
    }

    // Adding the two numbers
    sum := num1 + num2

    // Displaying the result
    fmt.Printf("The sum of %d and %d is: %d\n", num1, num2, sum)
}
