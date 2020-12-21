package main

import (
  "fmt"
  "bufio"
  "os"
)


func main() {
  var s string
  var vocali map[rune]int
  vocali = make(map[rune]int)

  b := bufio.NewScanner(os.Stdin)
  fmt.Print("Inserisci frase: ")
  b.Scan()
  s = b.Text()

  contaVocali(s, vocali)
  fmt.Print("Inserisci frase: ")
  b.Scan()
  s = b.Text()
  contaVocali(s, vocali)
}

func contaVocali(s string, vocali map[rune]int) {
  for _, j := range s {
    if j != ' ' {
      switch j {
      case 'a','e','i','o','u':
        vocali[j]++
      }
    }
  }
  for key, value := range vocali {
    fmt.Printf("%q",key)
    fmt.Println(value)
  }
}
