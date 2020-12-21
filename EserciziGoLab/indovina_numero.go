package main

import "fmt"
import "math/rand"

func main() {
  var indovina int
  const(
    MAX = 100
    SEED = 1
  )
  
  rand.Seed(SEED)
  x := rand.Intn(MAX)

  for i := 0; i < MAX / 2; i++ {
    fmt.Print("Indovinare numero generato: ")
    fmt.Scan(&indovina)
    if indovina > MAX {
      fmt.Println("fuori intervallo!")
    }
    if indovina == x {
      fmt.Print("tentativi usati per indovinare: ", i + 1)
      break
    }
    if i == MAX / 2 - 1 {
      fmt.Print("hai perso, il numero era ", x)
      break
    }
  }

}
