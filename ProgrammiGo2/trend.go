package main

import "fmt"

func main() {
  var intero, prec int

  fmt.Print("inserisci numero: ")
  fmt.Scan(&prec)

  for i := 0; i < 7; i++ {
    fmt.Print("inserisci numero: ")
    fmt.Scan(&intero)

    if intero > prec {
      fmt.Println("+")
    } else if intero < prec {
      fmt.Println("-")
    } else {
      fmt.Println("=")
    }
    prec = intero
  }
}
