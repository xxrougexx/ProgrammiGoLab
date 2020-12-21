package main

import "fmt"

func main() {
  var s1, s2 string
  fmt.Print("Inserisci prima parola: ")
  fmt.Scan(&s1)
  fmt.Print("Inserisci seconda parola: ")
  fmt.Scan(&s2)
  esito := isAnagram(s1,s2)
  if esito {
    fmt.Printf("%s e %s sono anagrammi", s1, s2)
  } else {
    fmt.Printf("%s e %s non sono anagrammi", s1, s2)
  }
}

func isAnagram(s1, s2 string) bool {
  var lettere1, lettere2 map[rune]int
  lettere1 = make(map[rune]int)
  lettere2 = make(map[rune]int)
  esito := false
  if len(s1) == len(s2) {
    for _, i := range s1 {
      lettere1[i]++
    }
    for _, j := range s2 {
      lettere2[j]++
    }
    for z := range lettere1 {
      if lettere1[z] == lettere2[z] {
        esito = true
      }
    }
  }
  return esito
}
