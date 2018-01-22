package main

import "fmt"

func main() {
    fmt.Println("hello, world\n")

    var age1 int32 = 30

    fmt.Printf("age1 = %d\n", age1)
    fmt.Printf("age1 type %T\n", age1)

    age2 := 32

    fmt.Printf("age2 = %d\n", age2)
    fmt.Printf("age2 type %T\n", age2)

    var age3 = 6

    fmt.Printf("age3 = %d\n", age3)
    fmt.Printf("age3 type %T\n", age3)

    var age4 int = 2

    fmt.Printf("age4 = %d\n", age4)
    fmt.Printf("age4 type %T\n", age4)
}
