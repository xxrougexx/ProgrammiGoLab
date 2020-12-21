package main

import "fmt"
import "io"

func main() {
  /*contaZeri serve solo se vengono inseriti
  dei valori successivi a uno o più zeri*/
  var mmPioggia, indice, contaZeri int
  fmt.Print("inserisci mm pioggia caduti: ")

  for {
    _, err := fmt.Scan(&mmPioggia)
    if err == io.EOF {
      break
    }
    if mmPioggia != 0 {
      indice++
      if zero != 0 {
        indice += contaZeri
        contaZeri = 0
      }
    } else {
      contaZeri++
    }
  }

  fmt.Printf("l'ultima pioggia è avvenuta al %d^ giorno", indice)
}
