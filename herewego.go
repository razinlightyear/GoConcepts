package main

import "fmt"

func main() {
  for j := 0; j < 5; j++ {
    fmt.Println(j)
  }

  if true { fmt.Println("true") }

  if false == false { fmt.Println("false") }
}
