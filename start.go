package main

import (
    "fmt"
)

func insertionSort(data []int) []int {
  for i := 1; i < len(data); i++ {
      fmt.Println("main iteration", i)
      for k := i; k > 0; k-- {
          fmt.Println("nested iteration", k)
          if data[k] < data[k-1] {
            
            s := data[k]
            b := data[k - 1]
            data[k - 1] = s
            data[k] = b
          }

      }
    }
    return data;
}

func main() {
    insertionSortData := []int{44, 4, 2, 47, 4, 8, 12, 1, 0, 2, 9, 23}
    fmt.Println("initial:", insertionSortData)
    fmt.Println("sorted:", insertionSort(insertionSortData  ))
}
