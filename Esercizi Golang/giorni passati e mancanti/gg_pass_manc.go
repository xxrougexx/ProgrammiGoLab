package main
import "fmt"

func main() {

  var m,
      ggAtt,
      ggNat,
      ggGen int

  fmt.Print("Inserire giorno del mese: ")
  fmt.Scan(&ggAtt)
  fmt.Print("Inserire mese: ")
  fmt.Scan(&m)

  ggAtt = ggAtt + ((m - 1) * 30)
  ggNat = 354 - ggAtt
  ggGen = ggAtt - 1

  fmt.Println("\n\t\tSono passati ", ggGen, " giorni dal primo di gennaio")
  fmt.Println("\t\tIl numero di giorni mancanti a Natale e': ", ggNat)

}
