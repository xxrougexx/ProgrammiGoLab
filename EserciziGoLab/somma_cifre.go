package main

import "fmt"

func main() {
  var num, cifra, somma int

  fmt.Print("Inserisci intero positivo: ")
  fmt.Scan(&num)

  for {
    cifra = num % 10
    somma += cifra
    num /= 10
    if num / 1 == 0 {
      break
    }
  }
  fmt.Println("La somma delle cifre del numero e':",somma)
}
