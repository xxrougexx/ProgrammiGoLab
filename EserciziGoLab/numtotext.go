package main

import (
  "fmt"
  "strings"
)

func main() {
  var n, str string
  var m1 map[string]string
  m1 = make(map[string]string)

  fmt.Print("un numero: ")
  fmt.Scan(&n)
  split := strings.Split(n, "")
  for i := range split {
    _, check := m1[split[i]]
    if check == false {
      fmt.Printf("parola per %s?",split[i])
      fmt.Scan(&str)
      m1[split[i]] = str
    }
  }
  for x := range split {
    fmt.Print(m1[split[x]]," ")
  }
}
