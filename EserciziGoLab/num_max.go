package main

import "fmt"

func main() {
  var n, cMax int
  m := 0

  for i := 0; i < 9; i++ {
    fmt.Print("Inserisci numero intero: ")
    fmt.Scan(&n)
    if n > m {
      m = n
      cMax = 0
    }
    if n == m {
      cMax++
    }
  }
  
  fmt.Printf("%d appare %d volte nella sequenza",m,cMax)
}
