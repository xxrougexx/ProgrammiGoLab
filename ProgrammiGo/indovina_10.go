package main
import "fmt"

func main(){
  var n int
  const INDOVINA = 10

  fmt.Print("Inserisci numero tra 1 e 10: ")
  fmt.Scan(&n)

  if n < 0 || n > 10 {
    fmt.Println("Valore non valido")
  } else if n == INDOVINA {
    fmt.Println("Hai indovinato!")
  } else {
    fmt.Println("Non hai indovinato")
  }
}
