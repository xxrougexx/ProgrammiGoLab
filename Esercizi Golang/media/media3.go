package main
import "fmt"

func main() {
  var val1, val2, val3, media float64
  fmt.Println("Inserisci primo valore: ")
  fmt.Scan(&val1)
  fmt.Println("Inserisci secondo valore: ")
  fmt.Scan(&val2)
  fmt.Println("Inserisci terzo valore: ")
  fmt.Scan(&val3)

  media = (val1 + val2 + val3) / 3
  fmt.Println("La media dei tre valori e': ", media)

}
