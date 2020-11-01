package main

import "fmt"

func main() {
  var capit, tasso, obiettivo, calcolo float64
  var anni int

  fmt.Print("Inserire capitale, tasso d'interesse e obiettivo finale: ")
  fmt.Scan(&capit,&tasso,&obiettivo)

  for {
    calcolo = capit * tasso + calcolo
    anni += 1
    if calcolo >= obiettivo {
      break
    }
  }
  fmt.Print("Serviranno ",anni + 1," anni per raggiungere l'obiettivo")
}
