package main

import (
  "fmt"
  "time"
  "math/rand"
)
//Ordine c-q-f-p
type Carta struct {
  valore, seme string
}

type Mazzo struct {
  carta []Carta
}

const (
  SEMI = 4
  VALORI = 13
  MAZZO = 52
)

func main() {
  var c string
  var n int
  rand.Seed(time.Now().UTC().UnixNano())
  fmt.Println("Inserire numero carta: ")
  fmt.Scan(&n)
  check, c := carta(n)
  if check {
    fmt.Println("Intervallo inserito fuori dal limite")
  }
  fmt.Println(c)
  check, random := carta(rand.Intn(MAZZO))
  if check {
    fmt.Println("Intervallo inserito fuori dal limite")
  }
  fmt.Println(random)
  mano := dai4Carte()
  for i := range mano {
    fmt.Print(mano[i]," ")
  }
}

func carta(num int) (bool, string) {
  valore := [...]string{"A", "2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K"}
  seme := [...]string{"C", "Q", "F", "P"}
  ok := false

  if num >= 52 {
    ok = true
    return ok, ""
  }
  return ok, valore[num % VALORI] + " di " + seme[num / VALORI]
}

func dai4Carte() []string {
  mano := make([]string, 4)

  for i := 0;i < 4;i++ {
    _, mano[i] = carta(rand.Intn(MAZZO))
  }
  return mano
}

func mazzo() {
  valori := [...]string{"A", "2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K"}
  semi := [...]string{"C", "Q", "F", "P"}
  var set Mazzo
  for i := 0; i < VALORI; i++ {
    for j := 0; j < SEMI; j++ {
      set = Mazzo {
        carta: []Carta {
          Carta{ valore: valori[i], seme: semi[j] },
        },
      }
      fmt.Println(set.carta)
    }
  }
}
