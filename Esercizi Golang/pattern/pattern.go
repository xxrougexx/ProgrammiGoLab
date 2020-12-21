package main

import "fmt"

func main() {
  var n int
  var spazio string
  fmt.Print("Inserire numero: ")
  fmt.Scan(&n)
  for riga := 0; riga < n; riga++ {
    for colonna := 0; colonna < n - riga; colonna++ {
        fmt.Print("*")
    }
    spazio += " "
    fmt.Print("\n", spazio)
  }

}
