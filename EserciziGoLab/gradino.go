package main

import "fmt"
import "io"

func main() {
  var valGradino, val1, val2, contaTemp, contaMax int
  fineGradino := false

  for {
    _, err1 := fmt.Scan(&val1)
    _, err2 := fmt.Scan(&val2)

    if err1 == io.EOF || err2 == io.EOF {
      break
    }
    
  }
}
