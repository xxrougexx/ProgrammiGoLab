package main
import "fmt"

func main() {
  var val1, val2, media float64
  fmt.Println("Inserisci primo valore: ")
  fmt.Scan(&val1)
  fmt.Println("Inserisci secondo valore: ")
  fmt.Scan(&val2)

  media = (val1 + val2) / 2
  fmt.Println("La media dei due valori e': ", media)

}
