package main

import "fmt"

func main() {
  var x, max, somma int
  var stringa string

  fmt.Println("Inserisci numeri (>-1): ")

  for {
    fmt.Scan(&x)
    if x == -1 {
      break
    }
    somma += x
    if max == 0 {
      max = x
    } else if max <= x {
      max = x
      stringa += "+"
    } else {
      stringa += "-"
    }
  }
  
  fmt.Println(stringa)
  fmt.Print("La somma di tutti i valori letti e': ", somma)
}
