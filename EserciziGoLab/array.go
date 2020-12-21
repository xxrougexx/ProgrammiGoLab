package main

import (
  "fmt"
  "io"
)

const DIM = 10

func main() {
  var numeri [DIM]int
  fmt.Print("array: ")
  i := 0
  for {
    _, err := fmt.Scan(&numeri[i])
    if err == io.EOF {
      break
    }
    i++
  }
  contrario := reverse(numeri)
  swapped := switchFirstLast(numeri)
  for j := range contrario {
    fmt.Print(contrario[j]," ")
  }
  fmt.Println()
  for x := range swapped {
    fmt.Print(swapped[x]," ")
  }
}

func reverse(array [DIM]int) [DIM]int {
	for i := 0; DIM-i-1 > i; i++ {
		array[i], array[DIM-i-1] = array[DIM-i-1], array[i]
	}
	return array
}

func switchFirstLast(array [DIM]int) [DIM]int {
  array[DIM-1], array[0] = array[0], array[DIM-1]
  return array
}
