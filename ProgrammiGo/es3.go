package main
import "fmt"

func main(){
  var n,somma int
  const K = 5
  for i := 0; i < K; i++{
    fmt.Print("Inserisci n: ")
    fmt.Scan(&n)
    somma = somma + n
  }
  fmt.Println(somma)
}
