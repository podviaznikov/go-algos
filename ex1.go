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


func count(data []int) int {
  // don't know how to check if array is empty in golang
  if len(data) == 0 {
    return 0
  }
  return 1 + count(data[1:])
}

func contains(data []int, item int) bool {
  if len(data) == 0 {
    return false
  }
  return data[0] == item || contains(data[1:], item)
}

func lastIndexOfIter(data []int, item int, index int) int {
  if len(data) == 0 {
    return -1
  }
  if data[len(data) - 1] == item {
    return index
  }
  return lastIndexOfIter(data[:len(data) - 1], item, index - 1)
}

func lastIndexOf(data []int, item int) int {
  if len(data) == 0 {
    return -1
  }
  return lastIndexOfIter(data, item, len(data) - 1)
}


func main() {
    data := []int{44, 4, 2, 47, 4, 8, 12, 1, 0, 2, 9, 23}
    fmt.Println("initial:", data)
    fmt.Println("sum:", sum(data))
    fmt.Println("sum:", count(data))
    fmt.Println("contains:", contains(data, 1))
    fmt.Println("lastIndexOf:", lastIndexOf(data, 47))
}
