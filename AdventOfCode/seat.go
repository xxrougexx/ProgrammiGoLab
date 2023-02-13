//F,L = 0 B,R = 1
package main

import (
  "fmt"
  "bufio"
)

func main() {
  var seatCodes []string
  var codedInt []int
  b := bufio.NewScanner(os.Stdin)
  for b.Scan() {
    seatCodes.append(b.Text())
    for i < 5 {
      if seatCodes[i] == 'B' {
        codedInt.append(1)
      } else {
        codedInt.append(0)
      }
    }
  }
}

func binToDec(bin []int) {
  var somma int
  for i := 0; i < 6; i++ {
    if i < 2 {
      if bin[i] == 1 {
        somma += 2 * i
      }
    } else {
      if bin[i] == 1 {
        somma += (2 * 2) * i
      }
    }
  }
}
