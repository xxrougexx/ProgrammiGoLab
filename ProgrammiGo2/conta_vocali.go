package main

import "fmt"

func main() {
  var str string
  var c int
  vocali := "aeiou"

  fmt.Print("Inserire parola: ")
  fmt.Scanf("%s", &str)

  for _, char := range str {
    for _, vow := range vocali {
      if char == vow {
        c++
      }
    }
  }
  
  fmt.Println(c)
}
