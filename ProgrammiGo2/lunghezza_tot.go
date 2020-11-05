package main

import "fmt"

func main() {
  var str string
  var totLen, lenStrings int

  fmt.Print("Inserire lunghezza totale frase: ")
  fmt.Scan(&totLen)

  for {
    fmt.Scanf("%s", &str)
    lenStrings += len(str)
    if lenStrings >= totLen {
      fmt.Println(lenStrings)
      break
    }
  }

}
