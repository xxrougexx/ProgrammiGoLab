package main

import "fmt"

func main() {
  var val1, val2 float64
  var segno string
  fmt.Scanf("%f%s%f",&val1, &segno, &val2)
  corretto, risultato := calc(val1, segno, val2)
  if corretto {
    fmt.Println("Il risultato dell'operazione è: ", risultato)
  } else {
    fmt.Println("Segno inserito non valido")
  }
}

func calc(x float64, sign string, y float64) (bool, float64) {
  var calcolo float64
  switch sign {
  case "+":
    calcolo = x + y
  case "-":
    calcolo = x - y
  case "*":
    calcolo = x * y
  case "/":
    calcolo = x / y
  default:
    return false, 0
  }
  return true, calcolo
}
