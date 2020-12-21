package main

import "fmt"

func main() {
  var n, vicino int
  const TARGET = 20

  for i := 0; i < 5; i++ {
    fmt.Print("Inserire numero intero tra 0 e 20: ")
    fmt.Scan(&n)
    if n > vicino && n < TARGET {
      vicino = n
    }
  }
  
  fmt.Println("Il valore più vicino al target preimpostato è: ", vicino)
}
