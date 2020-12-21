package main
import "fmt"

func main(){
  var imp,prezzo float64

  fmt.Println("Inserire imponibile: ")
  fmt.Scan(&imp)

  prezzo = (0.2 * imp) + imp

  fmt.Println("Il prezzo e': ", prezzo)

}
