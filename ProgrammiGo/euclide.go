package main
import "fmt"

func main(){
  var a, b, mcd, resto int

  fmt.Print("Inserisci a e b: ")
  fmt.Scan(&a,&b)

  for {
    if b == 0 {
      mcd = a
      fmt.Println(mcd)
      break
    } else {
      resto = a % b
      a, b = b, resto
    }
  }
}
