package main

import (
  "bufio"
  "fmt"
  "os"
  "strconv"
  "strings"
)

func main() {
  var correct bool
  var contaCorrect int
  var riga string
  b := bufio.NewScanner(os.Stdin)
  for b.Scan() {
    riga = b.Text()
    rigaSep := strings.Split(strings.TrimSpace(riga), ":")
    correct = passOk2(rigaSep[0], rigaSep[1])
    if correct {
      contaCorrect++
      fmt.Fprintf(os.Stdout, "Correct!\n")
    }
  }
  fmt.Println(contaCorrect)
}


func passOk(param string, pass string) bool {
  var conta int
  corretto := false
  pass = strings.TrimSpace(pass)
  pass1 := strings.Split(param, "-")
  pass2 := strings.Split(pass1[1], " ")
  min, _ := strconv.Atoi(pass1[0])
  max, _ := strconv.Atoi(pass2[0])
  chiave := pass2[1]
  if strings.ContainsAny(pass, chiave) {
    for i := range pass {
      if string(pass[i]) == chiave {
        conta++
      }
    }
    if conta <= max && conta >= min {
      corretto = true
    } else {
      corretto = false
    }
  } else {
    corretto = false
  }
  return corretto
}


func passOk2(param string, pass string) bool {
  corretto := false
  pass = strings.TrimSpace(pass)
  pass1 := strings.Split(param, "-")
  pass2 := strings.Split(pass1[1], " ")
  pos1, _ := strconv.Atoi(pass1[0])
  pos2, _ := strconv.Atoi(pass2[0])
  pos1--
  pos2--
  chiave := pass2[1]
  if strings.ContainsAny(pass, chiave) {
    if string(pass[pos1]) == chiave && string(pass[pos2]) != chiave || string(pass[pos1]) != chiave && string(pass[pos2]) == chiave {
      corretto = true
    } else {
      corretto = false
    }
  } else {
    corretto = false
  }
  return corretto
}
