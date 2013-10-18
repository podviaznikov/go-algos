package main

import (
    "fmt"
)


func sum(data []int) int {
  if len(data) == 0 {
    return 0
  }
  return data[0] + sum(data[1:]) 
}

func main() {
    data := []int{44, 4, 2, 47, 4, 8, 12, 1, 0, 2, 9, 23}
    fmt.Println("initial:", data)
    fmt.Println("sum:", sum(data))
}