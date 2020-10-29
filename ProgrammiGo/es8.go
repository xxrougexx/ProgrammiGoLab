package main
import "fmt"


func main(){
  var n, prod int
  const K = 10

  fmt.Print("Inserisci n: ")
  fmt.Scan(&n)

  for i := 1; i <= K; i++{
    prod = n * i
    fmt.Println(prod)
  }
}
