package main

import "fmt"

func calculate(a int, b int) []int {
    results := []int{a + b, a - b, a * b, a % b}
    return results
}

func main() {
    fmt.Println(calculate(20, 10))
    fmt.Println(calculate(700, 70))
}
