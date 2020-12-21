package main
import "fmt"

func main() {
  var anno, età int
  var flag string
  data := 2020

  fmt.Print("Inserire anno di nascita: ")
  fmt.Scan(&anno)

  età = data - anno
  if età >= 25 {
    flag = "la Camera e il Senato"
  } else if età < 25 && età >= 18 {
    flag = "la Camera"
  }

  fmt.Println("La persona con l'anno di nascita inserito può votare per", flag)

}
