package main

import "fmt"

func main() {
  var str string
  var maiu, minu int

  fmt.Print("Inserisci parola: ")
  fmt.Scan(&str)

  for _, char := range str {
    if char < 'a' {
      maiu += 1
    } else if char <= 'z' {
      minu += 1
    }
  }
  fmt.Printf("la stringa inserita ha %d lettere maiuscole e %d lettere minuscole",maiu,minu)
}
