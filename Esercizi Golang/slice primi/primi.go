package main

import "fmt"

func main() {
  var num int
  x := []int{}
  fmt.Print("Inserire numero: ")
  fmt.Scan(&num)

  x = append(x, primo(num, x)...)
  fmt.Println(x)
  fmt.Println("La lunghezza di questa slice e': ", len(x))
  fmt.Println("La capacita' di questa slice e': ", cap(x))

}

func primo(n int, slice []int) []int {
  for n > 0 {
    div := false
    for i := 2; i < n; i++ {
      if n % i == 0 {
        div = true
      }
    }

    if div != true {
      slice = append(slice, n)
      if (n + 1) % 10 == 0 {
          fmt.Print("*")
      }
    }
    n--
  }
  return slice
}
