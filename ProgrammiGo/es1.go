package main
import "fmt"

func main(){
  var n int
  fmt.Print("Inserisci n: ")
  fmt.Scan(&n)

  for i := n; i > 0; i--{
    fmt.Print("*")
  }
  fmt.Println()
}
