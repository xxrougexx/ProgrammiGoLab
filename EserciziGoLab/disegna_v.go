package main

import "fmt"

func main() {
  var backslash int
  spazio := ""

  fmt.Print("dimensione v: ")
  fmt.Scan(&backslash)
  slash := 2 * (backslash - 1)

  for i := 0; i < backslash; i++ {
    fmt.Print(spazio, "*")
    spazio2 := ""
    spazio += " "

    for j := 0; j < slash - 1; j++ {
      spazio2 += " "
    }
    fmt.Print(spazio2)
    if i != backslash - 1{
      fmt.Print("*")
    }
    fmt.Println()
    slash -= 2
  }
}
