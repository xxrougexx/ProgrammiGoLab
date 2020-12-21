package main

import "fmt"

func main() {
  var n int
  var s string
  a := "*"
  fmt.Scan(&n)
  for i := 0; i < n; i++ {
    for j := 0; j < n; j++ {
      s = ""
      for z := 0; z < n; z++ {
        s += " "
        fmt.Print(s)
      }
      fmt.Print(a)
    }
    fmt.Println()
  }




}
