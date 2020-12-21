package main

import (
  "amazon"
  "list"
  "fmt"
  "person"
)

type hasLength interface { Length() int }
type hasVolume interface { Volume() float64 }
type hasAge interface { Age() int }

func averageLength(x []hasLength) float64 {
  var sum int
  for _, b := range x {
    sum += b.Length()
  }
  return float64(sum) / float64(len(x))
}

func averageVolume(x []hasVolume) float64 {
  var sum float64
  for _, b := range x {
    sum += b.Volume()
  }
  return sum / float64(len(x))
}

func cercaEta(x []hasAge, y int) int {
  var count int
  for _, b := range x {
    if b.Age() <= y {
      count++
    }
  }
  return count
}

func main() {
  var x list.List
  x = x.AddInOrder("prova")
  x = x.AddInOrder("asd")
  fmt.Println("Pacchetto list: ")
  x.PrintRecursiveWithArrows()
  fmt.Printf("\nLunghezza della lista: %d\n", x.Length())

  var y []hasLength
  y = append(y, amazon.Box{30, 20, 10})
  y = append(y, amazon.Sphere{20})
  var z []hasVolume
  z = append(z, amazon.Box{30, 20, 10})
  z = append(z, amazon.Sphere{20})
  var j []hasAge
  j = append(j, person.Data{"Giovanni", "Boccia", 1999})

  fmt.Println("Pacchetto amazon: ")
  fmt.Printf("Lunghezza media: %.2f\n", averageLength(y))
  fmt.Printf("Volume medio: %.2f\n", averageVolume(z))
  fmt.Println("Pacchetto person: ")
  fmt.Printf("Persone con eta' minore o uguale a 30: %d\n", cercaEta(j, 30))
}
