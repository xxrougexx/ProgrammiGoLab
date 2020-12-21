package main

import "fmt"

const (
  GIORNI_IN_MESE = 30
  GIORNI_IN_ANNO = 365
)

func main() {
  var data1, data2 [3]int

  data1 = inputData(data1)
  data2 = inputData(data2)
  diff := differenza(data1, data2)
  fmt.Println(diff[2] * GIORNI_IN_ANNO + diff[1] * GIORNI_IN_MESE + diff[0])
}

func differenza(dat [3]int, dat1 [3]int) [3]int {
  var diff [3]int
  for i := 0; i < 3; i++ {
    diff[i] = dat[i] - dat1[i]
  }
  return diff
}

func inputData(dat [3]int) [3]int {
  for i := 0; i < 3; i++ {
    fmt.Scan(&dat[i])
  }
  return dat
}
