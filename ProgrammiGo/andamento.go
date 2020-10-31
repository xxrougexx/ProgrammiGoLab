package main

import "fmt"

func main() {
  var x, max, somma int
  for {
    fmt.Print("Inserisci numero (>-1): ")
    fmt.Scan(&x)
    somma += x
    if x == -1 {
      fmt.Print("La somma di tutti i valori letti e': ", somma)
      break
    }
    if max <= x {
      max = x
      fmt.Println("+")
    } else {
      fmt.Println("-")
    }
  }
}
