package main
import "fmt"

func main(){
  var n,doppio int
  const K = 5
  for i := 0; i < K; i++{
    fmt.Print("Inserisci n: ")
    fmt.Scan(&n)
    doppio = n * 2
    fmt.Println(doppio)
  }
}
