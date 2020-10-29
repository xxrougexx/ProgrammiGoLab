package main
import "fmt"

func main(){
  var n, quadN, quadK int

  fmt.Print("Inserisci n: ")
  fmt.Scan(&n)
  quadN = n * n
  for k := 1; quadK < quadN; k++{
    quadK = k * k
  }
  fmt.Println(quadK)
}
