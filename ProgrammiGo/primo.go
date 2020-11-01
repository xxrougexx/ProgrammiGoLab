package main

import "fmt"

func main() {
  var primo int
  div := false

  fmt.Print("Inserire numero: ")
  fmt.Scan(&primo)

  for i := 2; i < primo; i++ {
    if primo % i == 0 {
      div = true
    }
  }

  if div != true {
    fmt.Print("Il numero inserito e' primo")
  }
}
