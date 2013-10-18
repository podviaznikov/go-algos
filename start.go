package main

import (
    "fmt"
)

func swap (data []int, start int, end int) {
  s := data[start]
  e := data[end]
  data[start] = e
  data[end] = s
}
/*
  Stable
  O(1) extra space
  O(n2) comparisons and swaps
  Adaptive: O(n) time when nearly sorted
  Very low overhead
*/
func insertionSort(data []int) []int {
  for i := 1; i < len(data); i++ {
    for k := i; k > 0; k-- {
      if data[k] < data[k-1] {
        swap(data, k, k - 1)
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
        swap(data, i, k)
      }
    }
  }
  return data
}

/*
  Stable
  O(1) extra space
  O(n2) comparisons and swaps
  Adaptive: O(n) when nearly sorted
*/
func bubbleSort(data []int) []int {
  for i := 0; i < len(data); i++ {
    for k := 0; k < len(data) - 1 - i; k++ {
      if data[k + 1] < data[k] {
        swap(data, k + 1, k)
      }
    }
  }
  return data
}


func merge(left []int, right []int) []int {
  result := make([]int, len(left) + len(right))
  lc := 0
  rc := 0
  for i :=0; ; i++ {
    if i >= len(left) + len(right) {
      break
    }
    //rc >= len(right) || lc >= len(left) || lif rc < len(right)
    if len(left) > lc && len(right) <= rc {
      result[i] = left[lc]
      lc++
    } else {
        if len(right) > rc && len(left) <= lc {
          result[i] = right[rc]
          rc++
        } else {
          if left[lc] < right[rc] {
            result[i] = left[lc]
            lc++
          } else {
            result[i] = right[rc]
            rc++
        }
      }
    }
  }
  return result
}


func mergeSort(data []int) []int {
  start := 0
  end := len(data)
  if len(data) < 2 {
    return data
  }
  middle := (end - start)/2
  leftData := mergeSort(data[:middle])
  rightData := mergeSort(data[middle:])
  return merge(leftData, rightData)
}


func main() {
    insertionSortData := []int{44, 4, 2, 47, 4, 8, 12, 1, 0, 2, 9, 23}
    selectionSortData := []int{44, 4, 2, 47, 4, 8, 12, 1, 0, 2, 9, 23}
    bubbleSortData := []int{44, 4, 2, 47, 4, 8, 12, 1, 0, 2, 9, 23}
    mergeSortData := []int{44, 4, 2, 47, 4, 8, 12, 1}
    fmt.Println("initial:", insertionSortData)
    fmt.Println("inserted sort:", insertionSort(insertionSortData))
    fmt.Println("selection sort:", selectionSort(selectionSortData))
    fmt.Println("bubble sort:", bubbleSort(bubbleSortData))
    fmt.Println("merge sort:", mergeSort(mergeSortData))
}
