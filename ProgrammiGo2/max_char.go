package main

import "fmt"

func main() {
  var c, max byte

  for i := 0; i < 5; i++ {
    fmt.Print("inserisci carattere: ")
    fmt.Scanf("%c", &c)
    if c >= max {
      max = c
    }
    fmt.Scanf("%c", &c)
  }
  
  fmt.Println(string(max))
}
