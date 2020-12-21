package main

import "fmt"

func main() {
  fmt.Print("Inserire numero elementi 1^ slice: ")
  slice1 := popolaSlice()
  fmt.Print("Inserire numero elementi 2^ slice: ")
  slice2 := popolaSlice()
  stesso := confronto(slice1, slice2)
  if stesso {
    fmt.Println("Le due slice sono uguali")
  } else {
    fmt.Println("Le due slice sono diverse")
  }
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

func confronto(s1 []int, s2 []int) bool {
  var val bool
  if len(s1) == len(s2) {
    for i := range s1 {
      if s1[i] == s2[i] {
        val = true
      } else {
        val = false
        break
      }
    }
  }
  return val
}
