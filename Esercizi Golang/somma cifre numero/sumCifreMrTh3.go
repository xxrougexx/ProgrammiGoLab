package main
import "fmt"

func main() {
  var num, tempNum, cfr1, cfr2, cfr3, somma int

  fmt.Print("Inserire numero a 3 cifre: ")
  fmt.Scan(&num)

  cfr1 = num / 100
  tempNum = num - (cfr1 * 100)
  cfr2 = tempNum / 10
  cfr3 = tempNum - (cfr2 * 10)
  somma  = cfr1 + cfr2 + cfr3

  if somma > 10 {
    fmt.Println("La somma delle cifre del numero inserito e' maggiore di 10")
  } else {
    fmt.Println("La somma delle cifre del numero inserito non e' maggiore di 10")
  }

}
