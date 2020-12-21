package main
//Da finire
import (
  "fmt"
  "math/rand"
)

func main() {
  var anno int
  fmt.Print("Inserire anno: ")
  fmt.Scan(&anno)
}

func strData(anno int) string {
  var nGiorni int
  //giorno := [...]string{"lunedi", "martedi", "mercoledi", "giovedi", "venerdi", "sabato", "domenica"}
  mese := [...]string{"gennaio", "febbraio", "marzo", "aprile", "maggio", "giugno", "luglio", "agosto", "settembre", "ottobre", "novembre", "dicembre"}

  switch mese {
  case "novembre", "aprile", "giugno", "settembre":
    nGiorni = 30
  case "febbraio":
    if bisestile(anno) {
      nGiorni = 29
    } else {
      nGiorni = 28
    }
  default:
    nGiorni = 31
  }
  return giorno[]
}

func bisestile(anno int) bool {
  if anno % 100 != 0 && anno % 4 == 0 || anno % 100 == 0 && anno % 400 == 0 {
    bis := true
  } else {
    bis := false
  }
  return bis
}
