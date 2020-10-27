package main

import "fmt"

func main() {
  var reddito int
  var coniugato bool
  var tasse float64

  const (
    ALIQUOTA1 = 10
    ALIQUOTA2 = 25
    SCAGLIONE1 = 32000
    SCAGLIONE2 = 64000
  )

  fmt.Print("inserire reddito ")
  fmt.Scan(&reddito)
  fmt.Print("coniugato? (t/f) ")
  fmt.Scan(&coniugato)



  if coniugato == true {
    if reddito >= 0 && reddito <= SCAGLIONE1 {
      tasseInt := reddito * ALIQUOTA1
      tasse = (float64(tasseInt)) / 100
    } else {
      tasseInt := reddito * ALIQUOTA2
      tasse = (float64(tasseInt)) / 100
    }
  } else {
    if reddito >= 0 && reddito <= SCAGLIONE2 {
      tasseInt := reddito * ALIQUOTA1
      tasse = (float64(tasseInt)) / 100
    } else {
      tasseInt := reddito * ALIQUOTA2
      tasse = (float64(tasseInt)) / 100
    }
  }
  fmt.Println("L'importo da pagare è: ", tasse,"euro")
}
