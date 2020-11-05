package main

import "fmt"

func main() {
  var c rune
  var str string
  al := "altro"
  intervallo := 'A'
  //Stampa carattere inserito
  fmt.Print("Inserisci carattere: ")
  fmt.Scanf("%c", &c)

  fmt.Printf("%c\n", c)
  fmt.Printf("%c%c%c\n", c-1,c,c+1)

  for intervallo <= 'L' {
    if c == intervallo {
      al = "A-L"
    }
    intervallo++
  }

  fmt.Println(al)

  //Toglie a capo in più
  fmt.Scanf("%c", &c)

  //Stampa stringa inserita in colonna
  fmt.Print("Inserisci stringa: ")
  fmt.Scanf("%s", &str)

  for _, char := range str {
    fmt.Printf("%c\n", char)

  }
}
