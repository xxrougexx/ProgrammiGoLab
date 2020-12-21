package main

import "fmt"

func main() {
  var n int
  fmt.Print("Inserire numero di interi: ")
  fmt.Scan(&n)
  somma := make([]int, n)
  if n != 0 {
    fmt.Print("Inserisci valori: ")
    for i := range somma {
      fmt.Scan(&somma[i])
    }
  }
  sum := f(somma)
  fmt.Println("La somma totale è: ",sum)
}

func f(n []int) int {
  if len(n) == 0 {
    return 0
  }
  return n[0] + f(n[1:])
}
