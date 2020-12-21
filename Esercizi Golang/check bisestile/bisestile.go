package main
import "fmt"

func main() {
  var anno int

  fmt.Print("Inserire anno corrente: ")
  fmt.Scan(&anno)

  if anno % 100 != 0 && anno % 4 == 0 || anno % 100 == 0 && anno % 400 == 0 {
    fmt.Println("L'anno inserito e' bisestile")
  } else {
    fmt.Println("L'anno inserito non e' bisestile")
  }
}
