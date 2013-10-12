package main

import (
    "fmt"
)

/*
  Stable
  O(1) extra space
  O(n2) comparisons and swaps
  Adaptive: O(n) time when nearly sorted
  Very low overhead
*/
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

/*
  Not stable
  O(1) extra space
  Θ(n2) comparisons
  Θ(n) swaps
  Not adaptive
*/
func selectionSort(data []int) []int {
  for i := 0; i < len(data); i++ {
    for k := i+1; k < len(data); k++ {
      if data[k] < data[i] {
        s := data[k]
        b := data[i]
        data[k] = b
        data[i] = s
      }
    }
  }
  return data
}

func main() {
    insertionSortData := []int{44, 4, 2, 47, 4, 8, 12, 1, 0, 2, 9, 23}
    selectionSortData := []int{44, 4, 2, 47, 4, 8, 12, 1, 0, 2, 9, 23}
    fmt.Println("initial:", insertionSortData)
    fmt.Println("inserted sort:", insertionSort(insertionSortData))
    fmt.Println("selection sort:", selectionSort(selectionSortData))
}
