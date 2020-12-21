package main
import "fmt"

func main() {
  var num,frazionato float64
  var cNum int

  fmt.Println("Inserisci numero con virgola: ")
  fmt.Scan(&num)

  cNum = int(num)
  frazionato = num - float64(cNum)

  fmt.Println("La parte frazionaria del numero inserito e': ", frazionato)
}
