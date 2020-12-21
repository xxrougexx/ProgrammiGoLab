package main
import "fmt"

func main(){
  var imp,prezzo float64
  fmt.Println("Inserire prezzo finale: ")
  fmt.Scan(&prezzo)

  imp = prezzo / 1.2

  fmt.Println("L'imponibile è: ", imp)

}
