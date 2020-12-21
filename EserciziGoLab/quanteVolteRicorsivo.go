package main

import "fmt"

func main() {
  var s []rune
  var c rune
  var n int
  fmt.Print("Inserisci carattere: ")
  fmt.Scan(&c)
  fmt.Print("Inserisci lunghezza parola: ")
  fmt.Scan(&n)
  s = make([]rune, n)
  fmt.Print("\nInserisci parola: ")
  for i := range s {
    fmt.Scan(&s[i])
  }
  volte := quanto(s, c)
  fmt.Println(volte)
}

func quanto(s []rune, c rune) int {
    var n int
    if len(s) == 0 {
      return n
    }
    if s[0] == c {
      n++
    }
    return n + quanto(s[1:], c)
}
