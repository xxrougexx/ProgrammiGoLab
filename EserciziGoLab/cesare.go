package main

import "fmt"

func main() {
  var k byte
  var str string
  const LETTERE_ALFABETO = 26

  fmt.Print("chiave: ")
  fmt.Scan(&k)
  fmt.Print("caratteri da cifrare(minuscoli): ")
  fmt.Scan(&str)

  for i := 0; i < len(str); i++ {
    if str[i] + k < 'z' {
      fmt.Printf("%c",str[i] + k)
    } else {
      fmt.Printf("%c", str[i] - LETTERE_ALFABETO + k)
    }
  }

  fmt.Print(" è il testo cifrato")
}
