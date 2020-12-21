package main

import (
  "fmt"
  "math/rand"
  "time"
)

const (
  NUMERO_CARTE_PER_SEME = 13
  NUMERO_CARTE = NUMERO_CARTE_PER_SEME * 4
  CARTE_MANO = 5
)

type set []int

func main() {
  rand.Seed(time.Now().UTC().UnixNano())
  m := mazzoNuovo(true)
  for _, c := range m {
    fmt.Printf("%s\n", strCarta(c))
  }
  coppia := coppiaInMano(m)
  if coppia {
    fmt.Println("La mano pescata ha almeno una coppia")
  }
}

func mazzoNuovo(mescolato bool) (s set) {
  s = make(set, NUMERO_CARTE)
  for i := 0; i < NUMERO_CARTE; i++ {
    s[i] = i
  }
  if mescolato {
    for n := NUMERO_CARTE; n > 0; n--{
      pos := rand.Intn(n)
      s[pos], s[n-1] = s[n-1], s[pos]
    }
  }
  return
}

func strCarta(carta int) string {
  seme := [...]string{"cuori", "quadri", "fiori", "picche"}
  valore := [...]string{"A", "2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K"}

  return valore[carta % NUMERO_CARTE_PER_SEME] + " di " + seme[carta / NUMERO_CARTE_PER_SEME]
}

func valoreCarta(carta int) string {
  valore := [...]string{"A", "2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K"}
  return valore[carta % NUMERO_CARTE_PER_SEME]
}
func coppiaInMano(mazzo set) bool {
  mano := make(set, CARTE_MANO)
  var cop bool
  x := 1
  for i := range mano {
    mano[i] = mazzo[i]
  }
  fmt.Println("\nMano pescata: ")
  for _, c := range mano {
    fmt.Println(strCarta(c))
  }
  for i := range mano {
    for j := x; j < 5; j++ {
      if valoreCarta(mano[i]) == valoreCarta(mano[j]) {
        cop = true
      }
    }
    x++
  }
  return cop
}
