package main
import (
  "fmt"
  "strings"
  "bufio"
  "os"
  )

func main() {
  var s string
  b := bufio.NewScanner(os.Stdin)
  fmt.Print("Inserisci frase: ")


  b.Scan()
  s = b.Text()

  parole := strings.Fields(s)

  fmt.Println(len(parole))
}
