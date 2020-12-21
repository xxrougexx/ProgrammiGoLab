package main

import (
  "fmt"
  "time"
)

type Clock struct {
  hour, min, sec int
}

func main() {
  var ore, minuti, secondi int
  fmt.Println("Inserire orario di partenza: ")
  fmt.Scan(&ore, &minuti, &secondi)
  p := &Clock {
    hour: ore,
    min: minuti,
    sec: secondi,
  }
  i := 0
  for {
    time.Sleep(time.Duration(1) * time.Second)
    countdown(p)
    if p.hour == 0 && p.min == 0 && p.sec == 0 {
      fmt.Println("ZERO!")
      break
    }
    fmt.Println(p.hour, p.min, p.sec)
    i++
  }
}

func countdown(s *Clock) {
  (*s).sec--
  if (*s).sec == 0 {
    (*s).sec = 59
    changeMin(s)
  }

}

func changeMin(s *Clock) {
  (*s).min--
  if (*s).min == 0 {
    (*s).min = 59
    changeHour(s)
  }
}

func changeHour(s *Clock) {
  (*s).hour--
}
