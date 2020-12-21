package main
import "fmt"

func main(){
  var iva,imp,prezzo float64

  fmt.Println("Inserire prezzo finale: ")
  fmt.Scan(&prezzo)
  fmt.Println("Inserire imponibile: ")
  fmt.Scan(&imp)

  iva = (prezzo - imp) / imp

  fmt.Println("L'IVA e': ", iva)

}
