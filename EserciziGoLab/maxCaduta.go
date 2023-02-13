package main

import (
  "fmt"
  "io"
)

func main() {
  var num, max, min, diff int
  fmt.Print("Inserire numeri: ")
  _, err := fmt.Scan(&num)
  max = num
  for err != io.EOF {
    _, err = fmt.Scan(&num)
    if num >= max {
      max = num
    } else {
      min = num
    }
  }
  diff = max - min
  fmt.Println("La caduta maggiore e': ", diff)
}
