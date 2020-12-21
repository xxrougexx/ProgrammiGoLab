package main

import (
  "fmt"
  "io"
)

const DIM = 100

func main() {
  var alt[DIM] int
  media := 0
  alt = popolaArray(alt)

  stampaArray(alt)
  media = calcolaMedia(alt)
  fmt.Print("\nLa media delle altezze inserite e': ",media)
}

func popolaArray(h[DIM] int) [DIM]int {
  i := 0
  for {
    _, err := fmt.Scan(&h[i])
    if err == io.EOF {
      break
    }
    i++
  }
  return h
}
func stampaArray(h[DIM] int) {
  for j := range h {
    if h[j] == 0 {
      break
    }
    fmt.Print(h[j])
  }
}
func calcolaMedia(h[DIM] int) int {
  somma, num := 0, 0
  for i := range h {
    if h[i] == 0 {
      break
    }
    somma += h[i]
    num++
  }
  return(float64(somma)/float64(num))
}
