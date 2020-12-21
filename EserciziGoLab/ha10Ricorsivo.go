package main

import "fmt"

func main() {
  var n int
  fmt.Print("Inserire numero di interi: ")
  fmt.Scan(&n)
  ten := make([]int, n)
  if n != 0 {
    fmt.Print("Inserisci valori: ")
    for i := range ten {
      fmt.Scan(&ten[i])
    }
  }
  hasTen := ha10(ten)
  fmt.Println(hasTen)
}

func ha10(n []int) bool {
  if len(n) == 0 {
    return false
  }
  if n[0] == 10 {
    return true
  }
  return ha10(n[1:])
}
