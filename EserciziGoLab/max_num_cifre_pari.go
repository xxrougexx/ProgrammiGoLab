package main

import "fmt"

func main() {
  var cfr string
  var pari int

  fmt.Print("Inserisci sequenza di cifre: ")
  fmt.Scanf("%s",&cfr)

  for _, char := range cfr {
    if int(char) % 2 == 0 {
      pari++
    }
  }
  fmt.Printf("Tra le cifre inserite %d sono pari", pari)
}
