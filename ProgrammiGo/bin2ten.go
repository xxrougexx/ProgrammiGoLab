package main

import "fmt"

func main() {
  var bin, dec int
  i := 1
  valoreValido := true
  fmt.Print("Inserire numero binario: ")
  fmt.Scan(&bin)
  controlloBin := bin

  //Controllo input
  for {
    if controlloBin % 10 != 0 && controlloBin % 10 != 1 {
      fmt.Print("errore valore inserito non valido")
      valoreValido = false
      break
    }
    controlloBin /= 10
    if controlloBin / 1 == 0 {
      break
    }
  }

  //Conversione da binario a decimale
  for {
    if valoreValido == false {
      break
    }
    if bin % 10 == 1 {
      dec = dec + i
    }
    bin /= 10
    if bin / 1 == 0 {
      break
    }
    i *= 2
  }
  
  if valoreValido {
    fmt.Println("Il numero convertito in decimale e':",dec)
  }
}
