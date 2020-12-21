package main
import "fmt"

func main() {
  var divid, divis, quoz, resto int

  fmt.Println("Inserisci dividendo e divisore: ")
  fmt.Scan(&divid,&divis)

  quoz = divid / divis
  resto = divid % divis
  
  fmt.Println("Il quoziente della divisione e'", quoz,
              "mentre il resto e' ", resto)

}
