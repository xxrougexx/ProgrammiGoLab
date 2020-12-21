package main

import "fmt"

func main() {
  var n int
  var fib string
  fmt.Print("inserisci numero intero: ")
  fmt.Scan(&n)

  f1 := ""
  f2 := "*"

  for i := 1; i <= n; i++ {
    fib = f1 + f2
    f2, f1 = f1, fib
    fmt.Println(fib)
  }
}
