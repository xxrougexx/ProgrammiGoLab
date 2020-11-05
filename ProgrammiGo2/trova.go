package main

import "fmt"

func main() {
  var carat,invio rune
  var str string
  var pos int

  fmt.Print("inserisci carattere: ")
  fmt.Scanf("%c", &carat)
  fmt.Scanf("%c", &invio)
  fmt.Print("inserisci stringa: ")
  fmt.Scanf("%s", &str)

  for i, char := range str {
    if char == carat {
      pos = i
    }
  }

  if pos < 0 {
    fmt.Println(-1)
  } else {
    fmt.Println(pos)
  }
}
