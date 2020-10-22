package main
import "fmt"

func main(){
  var a int
  const (
    VOTO_MAX = 30
    VOTO_MIN = 0
  )
  fmt.Print("Inserisci numero compreso tra 0 e 30: ")
  fmt.Scan(&a)
  if !(a >= VOTO_MIN && a <= VOTO_MAX) {
    fmt.Println("Il voto inserito non è valido")
  }

}
