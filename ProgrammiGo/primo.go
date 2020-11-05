package main

import "fmt"
import "math"

func main() {
  var primo int
  div := false

  fmt.Print("Inserire numero: ")
  fmt.Scan(&primo)
  for i := 2; float64(i) < math.Sqrt(float64(primo)); i++ {
    if primo % i == 0 {
      div = true
    }
  }

  if div != true {
    fmt.Println("Il numero inserito e' primo")
  }
}
