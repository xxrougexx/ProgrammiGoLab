package main

import (
  "fmt"
  "time"
  "math/rand"
  "regexp"
)

func main() {
  var i int
  rand.Seed(time.Now().UTC().UnixNano())
  for i < 10 {
    rand := rand.Intn(4)
    fmt.Println(rand)
    i++
  }
  fmt.Println(isError())
  valph := regexp.MustCompile(`^\#?[a-z]*$`)
  fmt.Println(valph.MatchString("#fffas"))
}

func isError() error {
  err := fmt.Errorf("Prova")
  return err
}
