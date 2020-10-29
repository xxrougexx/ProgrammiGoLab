package main
import "fmt"


func main(){
  var n int

  fmt.Print("Inserisci n: ")
  fmt.Scan(&n)

  for i := 0; i <= n; i++{
    if i % 2 == 0 {
      fmt.Println(i)
    }
  }
}
