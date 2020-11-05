package main

import "fmt"

func main() {
  var n, i int

  spazio := ""

  fmt.Print("dimensione slash: ")
  fmt.Scan(&n)

  for i < n {
    fmt.Print(spazio, "*\n")
    spazio += " "
    i++
  }
}
