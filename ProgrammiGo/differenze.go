package main

import "fmt"

func main() {
  var n, diff float64
  
  for {
    fmt.Print("Inserire numero: ")
    fmt.Scan(&n)
    if diff == 0 {
      diff = n
    } else {
      fmt.Print("La differenza tra ", diff, " e ", n)
      diff -= n
      fmt.Println(" e':", diff)
    }
    if n == 0 {
      break
    }
  }
}
