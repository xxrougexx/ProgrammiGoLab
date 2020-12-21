package main

import "fmt"

func main() {
  var num, contaSeq, separatore int

  fmt.Print("inserisci sequenza di uni e zeri che finisca con 2: ")

  for {
    fmt.Scan(&num)

    if num == 2 {
      break
    }
    if num == 1 {
      separatore = num
    } else if separatore == 1 {
      contaSeq++
      separatore = num
    }
  }

  fmt.Print(contaSeq)
}
