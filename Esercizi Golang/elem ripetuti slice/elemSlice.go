package main

import "fmt"

func main() {
  fmt.Print("Inserire numero elementi 1^ slice: ")
  x := popolaSlice()
  fmt.Print("Inserire numero elementi 2^ slice: ")
  y := popolaSlice()
  stessi := contaUguali(x, y)
  fmt.Printf("%d valori di x compaiono in y", stessi)
}

func popolaSlice() []int {
  var n int
  fmt.Scan(&n)
  x := make([]int, n)
  for i := range x {
    fmt.Print("Inserisci valore: ")
    fmt.Scan(&x[i])
  }
  return x
}

func contaUguali(s1 []int, s2 []int) int {
  var val int
  for i := range s1 {
    for j := range s2 {
      if s1[i] == s2[j] {
        val++
      }
    }
  }
  return val
}
